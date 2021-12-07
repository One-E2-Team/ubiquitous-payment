package main

import (
	"errors"
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

func (p plugin) ExecuteTransaction(data dto.TransactionDTO) (dto.TransactionCreatedDTO, error) {
	var ret = dto.TransactionCreatedDTO{}
	var paypalOrder = dto.Order{}
	response, err := CallPayPalAPI(http.MethodPost, OrdersApiUrl, paypalOrder.DefaultInit(data.PspTransactionId,
		data.PayeeId, data.PayeeSecret, data.Currency, data.Amount, data.ClientBusinessName,
		data.SuccessUrl, data.FailUrl))
	if err != nil {
		return ret, err
	}
	id, ok := response["id"].(string)
	if !ok {
		return ret, errors.New("could not convert final transaction id")
	}
	ret.TransactionId = id
	linkObjects, ok := response["links"].([]interface{})
	if !ok {
		return ret, errors.New("could not convert final transaction links")
	}
	for _, l := range linkObjects {
		link, ok := l.(map[string]string)
		if !ok {
			return ret, errors.New("could not convert final transaction link object")
		}
		if link["method"] == "GET" && link["rel"] == "approve" {
			ret.RedirectUrl = link["href"]
			break
		}
	}
	if ret.RedirectUrl == "" {
		return ret, errors.New("could not find transaction redirect url")
	}
	return ret, nil
}

var Plugin plugin

func main() {

}
