package handler

import (
	"net/http"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/util"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var requestDTO dto.RegisterDTO
	err := util.UnmarshalRequest(r, &requestDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}

	err = handler.PSPService.Register(w, requestDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var loginCredentials dto.LoginDTO
	err := util.UnmarshalRequest(r, &loginCredentials)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	user, err := handler.PSPService.Login(loginCredentials)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}

	token, err := psputil.CreateToken(util.MongoID2String(user.ID), "psp")
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	resp := dto.LoginResponseDTO{Token: token, Username: user.Username,
		ProfileID: util.MongoID2String(user.ID), Roles: user.Roles}

	util.MarshalResult(w, resp)
}
