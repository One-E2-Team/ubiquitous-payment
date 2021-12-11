package dto

type RegisterDTO struct {
	Username     string   `json:"username" validate:"required,bad_username"`
	Password     string   `json:"password" validate:"required,common_pass,weak_pass"`
	WebShopName  string   `json:"webShopName" validate:"required"`
	PaymentTypes []string `json:"paymentTypes"`
}
