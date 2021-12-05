package dto

type ProductDTO struct {
	Name              string  `json:"name"`
	Price             float32 `json:"price"`
	Currency          string  `json:"currency"`
	Description       string  `json:"description"`
	NumOfInstallments uint    `json:"numOfInstallments"`
}
