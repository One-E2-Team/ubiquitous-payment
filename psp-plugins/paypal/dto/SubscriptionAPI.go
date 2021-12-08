package dto

import "ubiquitous-payment/psp-plugins/pspdto"

type Subscription struct {
	PlanId             string             `json:"plan_id"`
	InvoiceId          string             `json:"custom_id"`
	ApplicationContext ApplicationContext `json:"application_context"`
}

func (s *Subscription) Init(planId string, t pspdto.TransactionDTO) Subscription {
	s.PlanId = planId
	s.InvoiceId = t.PspTransactionId
	s.ApplicationContext = ApplicationContext{
		BrandName:   t.ClientBusinessName,
		Locale:      "en-RS",
		LandingPage: Login,
		UserAction:  PayNow,
		ReturnUrl:   t.SuccessUrl,
		CancelUrl:   t.FailUrl,
	}
	return *s
}
