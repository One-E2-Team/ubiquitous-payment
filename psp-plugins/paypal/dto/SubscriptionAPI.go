package dto

import "ubiquitous-payment/psp-plugins/pspdto"

type Subscription struct {
	PlanId             string                         `json:"plan_id"`
	InvoiceId          string                         `json:"custom_id"`
	ApplicationContext SubscriptionApplicationContext `json:"application_context"`
}

type SubscriptionApplicationContext struct {
	BrandName  string                 `json:"brand_name"`
	Locale     string                 `json:"locale"`
	UserAction SubscriptionUserAction `json:"user_action"`
	ReturnUrl  string                 `json:"return_url"`
	CancelUrl  string                 `json:"cancel_url"`
}

type SubscriptionUserAction string

const (
	SubscriptionContinue SubscriptionUserAction = "CONTINUE"
	SubscribeNow         SubscriptionUserAction = "SUBSCRIBE_NOW"
)

func (s *Subscription) Init(planId string, t pspdto.TransactionDTO) Subscription {
	s.PlanId = planId
	s.InvoiceId = t.PspTransactionId
	s.ApplicationContext = SubscriptionApplicationContext{
		BrandName:  t.ClientBusinessName,
		Locale:     "en-RS",
		UserAction: SubscribeNow,
		ReturnUrl:  t.SuccessUrl,
		CancelUrl:  t.FailUrl,
	}
	return *s
}
