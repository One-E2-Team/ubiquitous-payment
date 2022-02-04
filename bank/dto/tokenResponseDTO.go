package dto

import "ubiquitous-payment/bank/model"

type TokenResponseDTO struct {
	Token    string       `json:"token"`
	Username string       `json:"username"`
	ClientId uint         `json:"clientId"`
	Roles    []model.Role `json:"roles"`
}
