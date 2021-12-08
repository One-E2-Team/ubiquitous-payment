package dto

type Subscription struct {
	PlanId             string             `json:"plan_id"`
	InvoiceId          string             `json:"custom_id"`
	ApplicationContext ApplicationContext `json:"application_context"`
}
