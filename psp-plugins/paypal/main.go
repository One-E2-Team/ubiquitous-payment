package main

import (
	"fmt"
	"net/http"
	"ubiquitous-payment/psp-plugins/paypal/dto"
)

type plugin struct {
}

func (p plugin) Test() string {
	fmt.Println("Plug-in plug-out wasaaaaaaaaaa")
	return "ups"
}

var Plugin plugin

func main() {
	var data = dto.Order{}
	CallPayPalAPI(http.MethodPost, OrdersApiUrl, data.DefaultInit("transactionUUID1", "sb-064747x8893734@business.example.com", "35AYF8PFJWGPS", "USD", "50", "websgop Inc.", "http://localhost:200", "https://localhost:404"))
}
