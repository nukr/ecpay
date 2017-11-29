package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nukr/ecpay"
)

func main() {
	// staging
	merchantID := "2000132"
	hashkey := "5294y06JbISpM5x9"
	hashiv := "v77hoKGq4kWxNNIS"

	sdk := ecpay.ECPay{
		Environment:   ecpay.Staging,
		HashKey:       hashkey,
		HashIV:        hashiv,
		MerchantID:    merchantID,
		CheckMacValue: ecpay.CheckMacValue,
	}

	result, err := sdk.CreateTrade(ecpay.CreateTradeConfig{
		Amount:    5000,
		Desc:      "hihi",
		ItemName:  "gg",
		ReturnURL: "https://google.com",
		TradeDate: time.Now(),
		TradeNo:   randomString(20),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func randomString(strlen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}
