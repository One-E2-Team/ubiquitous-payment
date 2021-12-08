package pspdto

type TransactionDTO struct {
	PspTransactionId     string
	OrderId              string
	PayeeId              string
	PayeeSecret          string
	Currency             string
	Amount               string
	ClientBusinessName   string
	SuccessUrl           string
	FailUrl              string
	ErrorUrl             string
	PricingPlan          bool
	PaymentInterval      int
	NumberOfInstallments int
	InstallmentUnit      InstallmentUnit
}

type InstallmentUnit string

const (
	Day   InstallmentUnit = "DAY"
	Week  InstallmentUnit = "WEEK"
	Month InstallmentUnit = "MONTH"
	Year  InstallmentUnit = "YEAR"
)

type TransactionCreatedDTO struct {
	TransactionId string
	RedirectUrl   string
}
