package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetNewOrderId(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	orderId, err := handler.PSPService.CreateEmptyTransaction()
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetNewOrderId", loggingService)
		return
	}
	util.MarshalResult(w, orderId)
}

func (handler *Handler) FillTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	var request dto.WebShopOrderDTO
	err := util.UnmarshalRequest(r, &request)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"FillTransaction", loggingService)
		return
	}
	webShopOwnerID := psputil.GetLoggedUserIDFromToken(r)
	redirectLink, err := handler.PSPService.FillTransaction(request, webShopOwnerID)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"FillTransaction", loggingService)
		return
	}
	util.MarshalResult(w, redirectLink)
}

func (handler *Handler) GetAvailablePaymentTypeNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	pathVars := mux.Vars(r)
	payments, err := handler.PSPService.GetAvailablePaymentTypeNames(pathVars["transactionID"])
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetAvailablePaymentTypeNames", loggingService)
		return
	}
	util.MarshalResult(w, payments)
}

func (handler *Handler) SelectPaymentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	var request dto.SelectedPaymentTypeDTO
	err := util.UnmarshalRequest(r, &request)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"SelectPaymentType", loggingService)
		return
	}
	result, err := handler.PSPService.SelectPaymentType(request)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"SelectPaymentType", loggingService)
		return
	}
	util.MarshalResult(w, result)
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
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateTransactionSuccess", loggingService)
		return
	}
	http.Redirect(w, r, retUrl, http.StatusSeeOther)
}

func (handler *Handler) UpdateTransactionFail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	externalId := r.FormValue("token")
	retUrl, err := handler.PSPService.UpdateTransactionFail(externalId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateTransactionFail", loggingService)
		return
	}
	http.Redirect(w, r, retUrl, http.StatusSeeOther)
}

func (handler *Handler) CheckForPaymentBitcoin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(util.ContentType, util.ApplicationJson)
	pathVars := mux.Vars(r)
	result, err := handler.PSPService.CheckForPaymentBitcoin(pathVars["transactionID"])
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CheckForPaymentBitcoin", loggingService)
		return
	}
	util.MarshalResult(w, result)
}
