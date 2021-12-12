package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) AcceptWebShop(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r) // TODO: add method in util for path params
	err := handler.PSPService.ChangeWebShopAcceptance(pathVars["webShopID"], true)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"AcceptWebShop", loggingService)
		return
	}
	w.WriteHeader(http.StatusNoContent) // TODO: add appropriate status codes on every response
}

func (handler *Handler) DeclineWebShop(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	err := handler.PSPService.ChangeWebShopAcceptance(pathVars["webShopID"], false)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"DeclineWebShop", loggingService)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
