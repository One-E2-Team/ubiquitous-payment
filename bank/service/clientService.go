package service

import (
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/handler/mapper"
)

func (service *Service) GetMyAccount(clientId uint) (*dto.AccountResponseDTO, error) {
	client, err := service.Repository.GetClientById(clientId)
	if err != nil {
		return nil, err
	}

	return mapper.AccountToAccountResponseDTO(client.Accounts[0]), nil
}
