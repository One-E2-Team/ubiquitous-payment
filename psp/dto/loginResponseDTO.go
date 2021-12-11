package dto

import "ubiquitous-payment/psp/model"

type LoginResponseDTO struct {
	Token     string       `json:"token"`
	Username  string       `json:"username"`
	ProfileID string       `json:"profileId"`
	Roles     []model.Role `json:"roles"`
}
