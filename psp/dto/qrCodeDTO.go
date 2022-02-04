package dto

type QrCodeDTO struct {
	WebShopName 	string 		`json:"webShopName"`
	Currency 		string 		`json:"currency"`
	Amount 			string		`json:"amount"`
	AccountID   	string      `json:"accountId"`
}
