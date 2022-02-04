package handler

import (
	"net/http"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/util"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var requestDTO dto.RegistrationDTO
	err := util.UnmarshalRequest(r, &requestDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Register", loggingService)
		return
	}

	err = handler.BankService.Register(requestDTO, w)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Register", loggingService)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(util.ContentType, util.ApplicationJson)
	_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
}

func (handler *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials dto.LoginDTO
	err := util.UnmarshalRequest(r, &credentials)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"LogIn", loggingService)
		return
	}

	client, err := handler.BankService.LogIn(credentials)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"LogIn", loggingService)
		return
	}

	token, err := util.CreateToken(client.ID, "auth_service")
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"LogIn", loggingService)
		return
	}
	resp := dto.TokenResponseDTO{Token: token, Username: client.Username, ClientId: client.ID, Roles: client.Roles}
	util.MarshalResult(w, resp)
}
func (handler *Handler) ConfirmPassword(w http.ResponseWriter, r *http.Request) {
	var password string
	err := util.UnmarshalRequest(r, &password)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"ConfirmPassword", loggingService)
		return
	}

	confirmed, err := handler.BankService.ConfirmPassword(util.GetLoggedUserIDFromToken(r), password)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"ConfirmPassword", loggingService)
		return
	}
	util.MarshalResult(w, confirmed)
}
