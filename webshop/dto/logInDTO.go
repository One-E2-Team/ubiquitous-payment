package dto

type LogInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Passcode string `json:"passcode"`
}
