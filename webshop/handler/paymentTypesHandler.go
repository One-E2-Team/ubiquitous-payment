package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetValidPaymentTypes(w http.ResponseWriter, _ *http.Request) {
	result, err := handler.WSService.GetValidPaymentTypes()
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetValidPaymentTypes", loggingService)
		return
	}
	util.MarshalResult(w, result)
	return
}
