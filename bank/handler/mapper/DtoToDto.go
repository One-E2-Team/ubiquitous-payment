package mapper

import "ubiquitous-payment/bank/dto"

func PccOrderDtoToIssuerCardDto(pccOrder dto.PccOrderDTO) dto.IssuerCardDTO {
	return dto.IssuerCardDTO{
		Pan:        pccOrder.IssuerPAN,
		Cvc:        pccOrder.IssuerCVC,
		HolderName: pccOrder.IssuerHolderName,
		ValidUntil: pccOrder.IssuerValidUntil,
	}
}
