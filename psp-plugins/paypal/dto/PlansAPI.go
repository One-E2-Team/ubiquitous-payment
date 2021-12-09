package dto

import "ubiquitous-payment/psp-plugins/pspdto"

type Plan struct {
	ProductId          string             `json:"product_id"`
	PlanName           string             `json:"name"`
	PlanStatus         PlanStatus         `json:"status"`
	BillingCycles      []BillingCycle     `json:"billing_cycles"`
	PaymentPreferences PaymentPreferences `json:"payment_preferences"`
}

type PlanStatus string

const (
	Active   PlanStatus = "ACTIVE"
	Created  PlanStatus = "CREATED"
	Inactive PlanStatus = "INACTIVE"
)

type BillingCycle struct {
	PricingScheme PricingScheme `json:"pricing_scheme"`
	Frequency     Frequency     `json:"frequency"`
	TenureType    TenureType    `json:"tenure_type"`
	Sequence      int           `json:"sequence"`
	TotalCycles   int           `json:"total_cycles"`
}

type PricingScheme struct {
	Version    int        `json:"version"`
	FixedPrice FixedPrice `json:"fixed_price"`
}

type FixedPrice struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Frequency struct {
	IntervalUnit  IntervalUnit `json:"interval_unit"`
	IntervalCount int          `json:"interval_count"`
}

type IntervalUnit string

const (
	Day   IntervalUnit = "DAY"
	Week  IntervalUnit = "WEEK"
	Month IntervalUnit = "MONTH"
	Year  IntervalUnit = "YEAR"
)

type TenureType string

const (
	Regular TenureType = "REGULAR"
	Trial   TenureType = "TRIAL"
)

type PaymentPreferences struct {
	AutoBillOutstanding bool `json:"auto_bill_outstanding"`
	//SetupFee                SetupFee              `json:"setup_fee"`
	//SetupFeeFailureAction   SetupFeeFailureAction `json:"setup_fee_failure_action"`
	//FailurePaymentThreshold int                   `json:"payment_failure_threshold"`
}

type SetupFee FixedPrice

type SetupFeeFailureAction string

const (
	ContinuePlan SetupFeeFailureAction = "CONTINUE"
	Cancel       SetupFeeFailureAction = "CANCEL"
)

func (p *Plan) Init(productId string, t pspdto.TransactionDTO) Plan {
	p.ProductId = productId
	p.PlanName = t.PspTransactionId
	p.PlanStatus = Active
	sequence := 1
	if t.InstallmentDelayedTimeUnits != 0 {
		sequence = 2
		p.BillingCycles = append(p.BillingCycles, BillingCycle{
			PricingScheme: PricingScheme{
				Version: 0,
				FixedPrice: FixedPrice{
					CurrencyCode: "USD",
					Value:        "0",
				},
				//PricingModel: Volume,
			},
			Frequency: Frequency{
				IntervalUnit:  IntervalUnit(t.InstallmentUnit),
				IntervalCount: t.PaymentInterval,
			},
			TenureType:  Trial,
			Sequence:    1,
			TotalCycles: t.InstallmentDelayedTimeUnits,
		})
	}
	p.BillingCycles = append(p.BillingCycles, BillingCycle{
		PricingScheme: PricingScheme{
			Version: 0,
			FixedPrice: FixedPrice{
				CurrencyCode: t.Currency,
				Value:        t.Amount,
			},
			//PricingModel: Volume,
		},
		Frequency: Frequency{
			IntervalUnit:  IntervalUnit(t.InstallmentUnit),
			IntervalCount: t.PaymentInterval,
		},
		TenureType:  Regular,
		Sequence:    sequence,
		TotalCycles: t.NumberOfInstallments,
	})
	p.PaymentPreferences = PaymentPreferences{
		AutoBillOutstanding: true, /*
			SetupFee: SetupFee{
				CurrencyCode: "",
				Value:        "",
			},
			SetupFeeFailureAction:   "",
			FailurePaymentThreshold: 0,*/
	}
	return *p
}
