package dto

import "ubiquitous-payment/psp-plugins/pspdto"

type Order struct {
	Intent             Intent             `json:"intent"`
	PurchaseUnits      []PurchaseUnit     `json:"purchase_units"`
	ApplicationContext ApplicationContext `json:"application_context"`
}

type Intent string

const (
	Capture   Intent = "CAPTURE"
	Authorize Intent = "AUTHORIZE"
)

type PurchaseUnit struct {
	ReferenceId string `json:"reference_id"`
	Amount      Amount `json:"amount"`
	Payee       Payee  `json:"payee"`
	InvoiceId   string `json:"invoice_id"`
	// payment_instruction for psp platform fees
}

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Payee struct {
	Email            string `json:"email_address"`
	MerchantIdSecret string `json:"merchant_id"`
}

type ApplicationContext struct {
	BrandName   string      `json:"brand_name"`
	Locale      string      `json:"locale"`
	LandingPage LandingPage `json:"landing_page"`
	UserAction  UserAction  `json:"user_action"`
	ReturnUrl   string      `json:"return_url"`
	CancelUrl   string      `json:"cancel_url"`
}

type LandingPage string

const (
	Login        LandingPage = "LOGIN"
	Billing      LandingPage = "BILLING"
	NoPreference LandingPage = "NO_PREFERENCE"
)

type UserAction string

const (
	Continue UserAction = "CONTINUE"
	PayNow   UserAction = "PAY_NOW"
)

func (o *Order) Init(t pspdto.TransactionDTO) Order {
	o.Intent = Capture
	o.PurchaseUnits = append(o.PurchaseUnits, PurchaseUnit{
		ReferenceId: t.OrderId,
		Amount: Amount{
			CurrencyCode: t.Currency,
			Value:        t.Amount,
		},
		Payee: Payee{
			Email:            t.PayeeId,
			MerchantIdSecret: t.PayeeSecret,
		},
		InvoiceId: t.PspTransactionId,
	})
	o.ApplicationContext = ApplicationContext{
		BrandName:   t.ClientBusinessName,
		Locale:      "en-RS",
		LandingPage: Login,
		UserAction:  PayNow,
		ReturnUrl:   t.SuccessUrl,
		CancelUrl:   t.FailUrl,
	}
	return *o
}
