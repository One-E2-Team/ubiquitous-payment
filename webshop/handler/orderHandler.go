package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	retUrl, err := handler.WSService.CreateOrder(util.String2Uint(util.GetPathVariable(r, "id")), loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CreateOrder", loggingService)
		return
	}
	util.MarshalResult(w, retUrl)
}

func (handler *Handler) GetMyOrders(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	result, err := handler.WSService.GetMyOrders(loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyOrders", loggingService)
		return
	}
	util.MarshalResult(w, result)
}

func (handler *Handler) UpdatePspOrder(w http.ResponseWriter, r *http.Request) {
	id := util.GetPathVariable(r, "id")
	status := util.GetPathVariable(r, "status")
	err := handler.WSService.UpdatePspOrder(id, status)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyOrders", loggingService)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) GetSellersOrders(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	result, err := handler.WSService.GetSellersOrders(loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetSellersOrders", loggingService)
		return
	}
	util.MarshalResult(w, result)
}