package main

import (
	"fmt"
	"github.com/smartwalle/alipay"
)

func main() {
	client := alipay.New(appid, aliPublicKey, privateKey, false)

	p := alipay.AliPayTradePagePay{}
	p.OutTradeNo = "112"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.TotalAmount = "1211.12"
	p.Subject = "iPhone XR MIN"
	p.ReturnURL = "http://www.itsyuekao.com"

	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	pageUrl := url.String()
	fmt.Println(pageUrl)
}
