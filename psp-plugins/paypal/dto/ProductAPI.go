package dto

type Product struct {
	Name string      `json:"name"`
	Type ProductType `json:"type"`
}

type ProductType string

const (
	Physical ProductType = "PHYSICAL"
	Digital  ProductType = "DIGITAL"
	Service  ProductType = "SERVICE"
)

func (p *Product) Init() Product {
	p.Name = "NFT"
	p.Type = Digital
	return *p
}
