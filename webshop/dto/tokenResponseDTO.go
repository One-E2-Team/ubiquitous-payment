package dto

import "ubiquitous-payment/webshop/model"

type TokenResponseDTO struct {
	Token     string       `json:"token"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	ProfileId uint         `json:"profileId"`
	Roles     []model.Role `json:"roles"`
}
