package handler

import (
	"encoding/json"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/model"
)

func (handler *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var req dto.LogInDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"LogIn", loggingService)
		return
	}
	var user *model.User
	user, err = handler.WSService.LogIn(req)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"LogIn", loggingService)
		return
	}

	token, err := util.CreateToken(user.ProfileId, "auth_service")
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"LogIn", loggingService)
		return
	}
	resp := dto.TokenResponseDTO{Token: token, Email: user.Email, ProfileId: user.ProfileId,
		Roles: user.Roles, Username: user.Username}
	util.MarshalResult(w, resp)
}

func (handler *Handler) SetPSPAccessToken(w http.ResponseWriter, r *http.Request) {
	var accessUuid string
	err := util.UnmarshalRequest(r, &accessUuid)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"SetPSPAccessToken", loggingService)
		return
	}
	err = handler.WSService.SetPSPAccessToken(accessUuid)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"SetPSPAccessToken", loggingService)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (handler *Handler) ConfirmPassword(w http.ResponseWriter, r *http.Request) {
	var password string
	err := util.UnmarshalRequest(r, &password)
	if err != nil{
		util.HandleErrorInHandler(err, w, loggingClass+"ConfirmPassword", loggingService)
		return
	}
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	res, err := handler.WSService.ConfirmPassword(loggedUserId, password)
	if err != nil{
		util.HandleErrorInHandler(err, w, loggingClass+"ConfirmPassword", loggingService)
		return
	}
	util.MarshalResult(w, res)
}