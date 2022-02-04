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

func (service *Service) GetMyTransactions(clientId uint) ([]dto.TransactionResponseDTO, error) {
	client, err := service.Repository.GetClientById(clientId)
	if err != nil {
		return nil, err
	}

	panNumbers, err := service.Repository.GetPanNumbersByClientId(clientId)
	if err != nil {
		return nil, err
	}

	transactions, err := service.Repository.GetClientTransactions(client.Accounts[0].AccountNumber, panNumbers)
	response := make([]dto.TransactionResponseDTO, 0)
	for _, transaction := range transactions {
		response = append(response, mapper.TransactionToTransactionResponseDTO(transaction))
	}
	return response, nil
}
