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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *model.User
	user, err = handler.WSService.LogIn(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := util.CreateToken(user.ProfileId, "auth_service")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := dto.TokenResponseDTO{
		Token:     token,
		Email:     user.Email,
		ProfileId: user.ProfileId,
		Roles:     user.Roles,
		Username:  user.Username,
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respJson)
	w.Header().Set(util.ContentType, util.ApplicationJson)
}
