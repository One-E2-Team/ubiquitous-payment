package handler

import (
	"net/http"
	"ubiquitous-payment/psp/dto"
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
	var request dto.WebShopOrderDTO
	err := util.UnmarshalRequest(r, &request)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	redirectLink, err := handler.PSPService.FillTransaction(request, util.GetWebShopNameFromToken(r))
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, redirectLink)
}
