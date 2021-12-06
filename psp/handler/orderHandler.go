package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetNewOrderId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	orderId, err := handler.PSPService.CreateEmptyTransaction()
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	_, err = w.Write([]byte(orderId))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
