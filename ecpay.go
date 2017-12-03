package ecpay

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

//go:generate stringer -type=Environment

//Environment ...
type Environment int

// env
const (
	Production Environment = iota
	Staging
)

// CreateTradeURLMap ...
var CreateTradeURLMap = map[Environment]string{
	Production: "https://payment.ecpay.com.tw/SP/CreateTrade",
	Staging:    "https://payment-stage.ecpay.com.tw/SP/CreateTrade",
}

// ScriptURLMap ...
var ScriptURLMap = map[Environment]string{
	Production: "https://payment.ecpay.com.tw/Scripts/SP/ECPayPayment_1.0.0.js",
	Staging:    "https://payment-stage.ecpay.com.tw/Scripts/SP/ECPayPayment_1.0.0.js",
}

// ECPay ...
type ECPay struct {
	Environment   Environment
	MerchantID    string
	HashKey       string
	HashIV        string
	CheckMacValue func(string) string
}

// CreateTradeConfig ...
type CreateTradeConfig struct {
	Amount    int64
	Desc      string
	ItemName  string
	TradeNo   string
	TradeDate time.Time
	ReturnURL string
}

// CreateTradeResponse ...
type CreateTradeResponse struct {
	ScriptURL  string
	MerchantID string
	SPToken    string
}

// CreateTrade ...
func (ecpay *ECPay) CreateTrade(config *CreateTradeConfig) (*CreateTradeResponse, error) {
	form := url.Values{
		"MerchantID":        []string{ecpay.MerchantID},
		"MerchantTradeNo":   []string{config.TradeNo},
		"MerchantTradeDate": []string{config.TradeDate.Format("2006/01/02 15:04:05")},
		"PaymentType":       []string{"aio"},
		"TotalAmount":       []string{strconv.FormatInt(config.Amount, 10)},
		"TradeDesc":         []string{config.Desc},
		"ItemName":          []string{config.ItemName},
		"ReturnURL":         []string{config.ReturnURL},
		"ChoosePayment":     []string{"ALL"},
	}
	var sortable [][]string
	for k, v := range form {
		sortable = append(sortable, []string{k, v[0]})
	}
	sort.Slice(sortable, func(i, j int) bool {
		return sortable[i][0] < sortable[j][0]
	})
	var temp []string
	for _, s := range sortable {
		temp = append(temp, strings.Join(s, "="))
	}
	result := fmt.Sprintf("HashKey=%s&%s&HashIV=%s", ecpay.HashKey, strings.Join(temp, "&"), ecpay.HashIV)
	form.Set("CheckMacValue", CheckMacValue(result))
	resp, err := http.Post(CreateTradeURLMap[ecpay.Environment], "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rtncode := gjson.GetBytes(respBody, "RtnCode").String()
	if rtncode != "1" {
		return nil, errors.New(string(respBody))
	}
	sptoken := gjson.GetBytes(respBody, "SPToken").String()
	return &CreateTradeResponse{
		ScriptURL:  ScriptURLMap[ecpay.Environment],
		MerchantID: ecpay.MerchantID,
		SPToken:    sptoken,
	}, nil
}
