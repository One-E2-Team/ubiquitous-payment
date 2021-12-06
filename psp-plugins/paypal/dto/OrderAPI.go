package dto

type ApplicationContext struct {
}

type Order struct {
	Intent             string             `json:"intent"`
	PurchaseUnits      []PurchaseUnits    `json:"purchase_units"`
	ApplicationContext ApplicationContext `json:"application_context""`
}

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}
type PurchaseUnits struct {
	Amount Amount `json:"amount"`
}
