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
