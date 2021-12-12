package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	pathVars := mux.Vars(r)
	retUrl, err := handler.WSService.CreateOrder(util.String2Uint(pathVars["id"]), loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CreateOrder", loggingService)
		return
	}
	util.MarshalResult(w, retUrl)
}
