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

func (handler *Handler) GetAccessTokenForWebShop(w http.ResponseWriter, r *http.Request) {
	accessToken, err := handler.PSPService.GetAccessTokenForWebShop(psputil.GetLoggedUserIDFromToken(r))
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, accessToken)
}

func (handler *Handler) LoginWebShop(w http.ResponseWriter, r *http.Request) {
	var webShopLoginDTO dto.WebShopLoginDTO
	err := util.UnmarshalRequest(r, &webShopLoginDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	resp, err := handler.PSPService.LoginWebShop(webShopLoginDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, resp)
}
