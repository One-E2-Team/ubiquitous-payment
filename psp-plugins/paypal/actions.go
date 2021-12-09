package main

import (
	"errors"
	"net/http"
	"ubiquitous-payment/psp-plugins/paypal/dto"
	"ubiquitous-payment/psp-plugins/pspdto"
)

func ExecuteOrder(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	var ret = pspdto.TransactionCreatedDTO{}
	var paypalOrder = dto.Order{}
	response, err := CallPayPalAPI(http.MethodPost, OrdersApiUrl, paypalOrder.Init(data))
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

func ExecuteSubscription(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	planId, err := createPlan(data)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	return createSubscription(planId, data)
}

func createPlan(data pspdto.TransactionDTO) (string, error) {
	return "", nil
}

func createSubscription(planId string, data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	return pspdto.TransactionCreatedDTO{}, nil
}
