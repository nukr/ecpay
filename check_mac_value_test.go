package ecpay

import "testing"

func TestCheckMacValue(t *testing.T) {
	cases := []struct {
		src      string
		expected string
	}{
		{
			src:      "HashKey=5294y06JbISpM5x9&ChoosePayment=Credit&ClientBackURL=https://developers.allpay.com.tw/AioMock/MerchantClientBackUrl&CreditInstallment=&EncryptType=1&InstallmentAmount=&ItemName=MacBook 30元X2#iPhone6s 40元X1&MerchantID=2000132&MerchantTradeDate=2017/10/26 17:06:56&MerchantTradeNo=DX20171026170656oe09&PaymentType=aio&Redeem=&ReturnURL=https://developers.allpay.com.tw/AioMock/MerchantReturnUrl&StoreID=&TotalAmount=5&TradeDesc=建立信用卡測試訂單&HashIV=v77hoKGq4kWxNNIS",
			expected: "2E5E9B603748EA4D51F49E172B495020F3585BECFB78BC934682D40A0703C2B6",
		},
	}

	for _, c := range cases {
		actual := CheckMacValue(c.src)
		if actual != c.expected {
			t.Errorf("expected %s, actual %s", c.expected, actual)
		}
	}
}

func TestDotNetEncode(t *testing.T) {
	cases := []struct {
		src      string
		expected string
	}{
		{
			src:      "1234567890abcdefghijklmnopqrstuvwxyz",
			expected: "1234567890abcdefghijklmnopqrstuvwxyz",
		},
		{
			src:      "-_.%21%2a%28%29",
			expected: "-_.!*()",
		},
	}

	for _, c := range cases {
		actual := dotNetURLEncode(c.src)
		if actual != c.expected {
			t.Errorf("expected %s, actual %s", c.expected, actual)
		}
	}
}
