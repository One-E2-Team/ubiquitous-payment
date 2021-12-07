package handler

import (
	"encoding/json"
	"net/http"
	dto2 "ubiquitous-payment/psp/dto"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetNewOrderId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	orderId, err := handler.PSPService.CreateEmptyTransaction()
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, orderId)
}

func (handler *Handler) FillTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	var dto dto2.WebShopOrderDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	redirectLink, err := handler.PSPService.FillTransaction(dto)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, redirectLink)
}
