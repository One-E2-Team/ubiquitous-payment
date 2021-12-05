package dto

type RegistrationDto struct {
	Username     string   `json:"username" validate:"required,bad_username"`
	Password     string   `json:"password" validate:"required,common_pass,weak_pass"`
	Name         string   `json:"name" validate:"required"`
	Email        string   `json:"email" validate:"required,email"`
	Role         string   `json:"role" validate:"required"`
}
