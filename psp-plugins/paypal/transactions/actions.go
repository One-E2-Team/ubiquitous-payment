package transactions

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
		link, ok := l.(map[string]interface{})
		if !ok {
			return ret, errors.New("could not convert final transaction link object")
		}
		method, ok := link["method"].(string)
		if !ok {
			return ret, errors.New("could not convert final transaction link internal object data - method")
		}
		rel, ok := link["rel"].(string)
		if !ok {
			return ret, errors.New("could not convert final transaction link internal object data - rel")
		}
		href, ok := link["href"].(string)
		if !ok {
			return ret, errors.New("could not convert final transaction link internal object data - href")
		}
		if method == "GET" && rel == "approve" {
			ret.RedirectUrl = href
			break
		}
	}
	if ret.RedirectUrl == "" {
		return ret, errors.New("could not find transaction redirect url")
	}
	return ret, nil
}
