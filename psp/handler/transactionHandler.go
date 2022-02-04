package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetDataForQrCode(w http.ResponseWriter, r *http.Request) {
	transactionId := util.GetPathVariable(r, "id")
	w.Header().Set(util.ContentType, util.ApplicationJson)
	result, err := handler.PSPService.GetDataForQrCode(transactionId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetDataForQrCode", loggingService)
		return
	}
	util.MarshalResult(w, result)
}
