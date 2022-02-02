package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetAccountsByPaymentType(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	paymentTypeName := util.GetPathVariable(r, "name")
	ret, err := handler.WSService.GetAccountsByPaymentName(paymentTypeName, loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetAccountsByPaymentType", loggingService)
		return
	}
	util.MarshalResult(w, ret)
}
