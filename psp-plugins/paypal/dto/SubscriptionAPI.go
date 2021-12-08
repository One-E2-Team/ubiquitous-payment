package dto

type Subscription struct {
	PlanId             string             `json:"plan_id"`
	InvoiceId          string             `json:"custom_id"`
	ApplicationContext ApplicationContext `json:"application_context"`
}

func (s *Subscription) DefaultInit(planId string, pspOrderId string, webshop string, returnUrl string, cancelUrl string) Subscription {
	s.PlanId = planId
	s.InvoiceId = pspOrderId
	s.ApplicationContext = ApplicationContext{
		BrandName:   webshop,
		Locale:      "en-RS",
		LandingPage: Login,
		UserAction:  PayNow,
		ReturnUrl:   returnUrl,
		CancelUrl:   cancelUrl,
	}
	return *s
}
