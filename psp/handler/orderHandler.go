package handler

import (
	"github.com/gorilla/mux"
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

func (handler *Handler) GetAvailablePaymentTypeNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	pathVars := mux.Vars(r)
	payments, err := handler.PSPService.GetAvailablePaymentTypeNames(pathVars["transactionID"])
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, payments)
}

func (handler *Handler) SelectPaymentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	var request dto.SelectedPaymentTypeDTO
	err := util.UnmarshalRequest(r, &request)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	redirectUrl, err := handler.PSPService.SelectPaymentType(request)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, redirectUrl)
}

func (handler *Handler) UpdateTransactionSuccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	externalId := r.FormValue("token")
	subscriptionId := r.FormValue("subscription_id")
	if subscriptionId != "" {
		externalId = subscriptionId
	}
	retUrl, err := handler.PSPService.UpdateTransactionSuccess(externalId)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	http.Redirect(w, r, retUrl, http.StatusSeeOther)
}

func (handler *Handler) UpdateTransactionFail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	externalId := r.FormValue("token")
	retUrl, err := handler.PSPService.UpdateTransactionFail(externalId)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	http.Redirect(w, r, retUrl, http.StatusSeeOther)
}

