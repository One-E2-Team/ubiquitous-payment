package handler

import (
	"encoding/json"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/model"
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

func (handler *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	id := util.String2Uint(util.GetPathVariable(r, "id"))
	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil{
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateAccount", loggingService)
	}
	err = handler.WSService.UpdateAccount(&account, id)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateAccount", loggingService)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil{
		util.HandleErrorInHandler(err, w, loggingClass+"CreateAccount", loggingService)
	}
	account.ProfileId = loggedUserId
	err = handler.WSService.CreateAccount(&account)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CreateAccount", loggingService)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := util.String2Uint(util.GetPathVariable(r, "id"))
	err := handler.WSService.DeleteAccount(id)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"DeleteAccount", loggingService)
		return
	}
	w.WriteHeader(http.StatusOK)
}
