package dto

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

func (o *Order) DefaultInit(pspOdrderID string, orderID string, payeeID string, payeeSecret string, currency string, amount string, webshop string, returnUrl string, cancelUrl string) Order {
	o.Intent = Capture
	o.PurchaseUnits = append(o.PurchaseUnits, PurchaseUnit{
		ReferenceId: orderID,
		Amount: Amount{
			CurrencyCode: currency,
			Value:        amount,
		},
		Payee: Payee{
			Email:            payeeID,
			MerchantIdSecret: payeeSecret,
		},
		InvoiceId: pspOdrderID,
	})
	o.ApplicationContext = ApplicationContext{
		BrandName:   webshop,
		Locale:      "en-RS",
		LandingPage: Login,
		UserAction:  PayNow,
		ReturnUrl:   returnUrl,
		CancelUrl:   cancelUrl,
	}
	return *o
}
