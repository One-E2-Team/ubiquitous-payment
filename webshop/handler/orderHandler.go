package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	pathVars := mux.Vars(r)
	err := handler.WSService.CreateOrder(util.String2Uint(pathVars["id"]), loggedUserId)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
