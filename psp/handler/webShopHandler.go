package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) AcceptWebShop(w http.ResponseWriter, r *http.Request) {
	err := handler.PSPService.ChangeWebShopAcceptance(util.GetPathVariable(r, "webShopID"), true)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"AcceptWebShop", loggingService)
		return
	}
	w.WriteHeader(http.StatusNoContent) // TODO: add appropriate status codes on every response
}

func (handler *Handler) DeclineWebShop(w http.ResponseWriter, r *http.Request) {
	err := handler.PSPService.ChangeWebShopAcceptance(util.GetPathVariable(r, "webShopID"), false)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"DeclineWebShop", loggingService)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
