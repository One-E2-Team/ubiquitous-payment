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
	ret.RedirectUrl, err = extractApproveLinkFromLinksInterface(response["links"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func ExecuteSubscription(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	productId, err := createRandomProduct()
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	planId, err := createPlan(productId, data)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	return createSubscription(planId, data)
}

func createRandomProduct() (string, error) {
	var product = dto.Product{}
	response, err := CallPayPalAPI(http.MethodPost, CatalogProductApiUrl, product.Init())
	if err != nil {
		return "", err
	}
	id, ok := response["id"].(string)
	if !ok {
		return "", errors.New("could not convert plan id")
	}
	return id, nil
}

func createPlan(productId string, data pspdto.TransactionDTO) (string, error) {
	var plan = dto.Plan{}
	response, err := CallPayPalAPI(http.MethodPost, PlansApiUrl, plan.Init(productId, data))
	if err != nil {
		return "", err
	}
	id, ok := response["id"].(string)
	if !ok {
		return "", errors.New("could not convert plan id")
	}
	status, ok := response["status"].(string)
	if !ok {
		return "", errors.New("could not convert plan status")
	}
	if status != "ACTIVE" {
		return "", errors.New("plan status is not in ACTIVE state")
	}
	return id, nil
}

func createSubscription(planId string, data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	var ret, subscription = pspdto.TransactionCreatedDTO{}, dto.Subscription{}
	response, err := CallPayPalAPI(http.MethodPost, SubscriptionsApiUrl, subscription.Init(planId, data))
	if err != nil {
		return ret, err
	}
	id, ok := response["id"].(string)
	if !ok {
		return ret, errors.New("could not convert subscription id")
	}
	ret.TransactionId = id
	status, ok := response["status"].(string)
	if !ok {
		return ret, errors.New("could not convert subscription status")
	}
	if status != "APPROVAL_PENDING" {
		return ret, errors.New("subscription status is not in APPROVAL_PENDING state")
	}
	ret.RedirectUrl, err = extractApproveLinkFromLinksInterface(response["links"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func extractApproveLinkFromLinksInterface(links interface{}) (string, error) {
	var ret string = ""
	linkObjects, ok := links.([]interface{})
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
			ret = href
			break
		}
	}
	if ret == "" {
		return ret, errors.New("could not find final client redirect approve url")
	}
	return ret, nil
}

func CaptureOrderPayment(id string) (bool, error) {
	response, err := CallPayPalAPI(http.MethodPost, "https://api.sandbox.paypal.com/v2/checkout/orders/"+id+"/capture", nil)
	if err != nil {
		return false, err
	}
	status, ok := response["status"].(string)
	if !ok {
		return false, errors.New("could not convert status of captured payment")
	}
	if status == "COMPLETED" {
		return true, nil
	} else {
		return false, nil
	}
}

func CaptureSubscriptionApproval(id string) (bool, error) {
	response, err := CallPayPalAPI(http.MethodGet, "https://api-m.sandbox.paypal.com/v1/billing/subscriptions/"+id, nil)
	if err != nil {
		return false, err
	}
	status, ok := response["status"].(string)
	if !ok {
		return false, errors.New("could not convert status of subscription")
	}
	if status == "ACTIVE" || status == "EXPIRED" {
		return true, nil
	} else {
		return false, nil
	}
}
