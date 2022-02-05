package handler

import (
	"net/http"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetMyPaymentTypes(w http.ResponseWriter, r *http.Request) {
	loggedUserId := psputil.GetLoggedUserIDFromToken(r)
	w.Header().Set(util.ContentType, util.ApplicationJson)
	result, err := handler.PSPService.GetMyPaymentTypes(loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyPaymentTypes", loggingService)
		return
	}
	util.MarshalResult(w, result)
}

func (handler *Handler) UpdateMyPaymentTypes(w http.ResponseWriter, r *http.Request) {
	loggedUserId := psputil.GetLoggedUserIDFromToken(r)
	var paymentTypes []string
	err := util.UnmarshalRequest(r, &paymentTypes)
	w.Header().Set(util.ContentType, util.ApplicationJson)
	err = handler.PSPService.UpdateMyPaymentTypes(loggedUserId, paymentTypes)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateMyPaymentTypes", loggingService)
		return
	}

	w.WriteHeader(http.StatusOK)
}