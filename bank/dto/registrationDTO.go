package dto

type RegistrationDTO struct {
	Username string `json:"username" validate:"required,bad_username"`
	Password string `json:"password" validate:"required,common_pass,weak_pass"`
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
}
