package service

import (
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/handler/mapper"
	"ubiquitous-payment/bank/model"
)

func (service *Service) GetMyAccount(clientId uint) (*dto.AccountResponseDTO, error) {
	client, err := service.Repository.GetClientById(clientId)
	if err != nil {
		return nil, err
	}
	var account *model.ClientAccount
	account = nil
	if len(client.Accounts) > 0 {
		account = &client.Accounts[0]
	}
	return mapper.AccountToAccountResponseDTO(account), nil
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

	acquirerTransactions, err := service.Repository.GetAcquirerTransactions(client.Accounts[0].AccountNumber)
	issuerTransactions, err := service.Repository.GetIssuerTransactions(panNumbers)
	response := make([]dto.TransactionResponseDTO, 0)

	for _, transaction := range acquirerTransactions {
		response = append(response, mapper.TransactionToTransactionResponseDTO(transaction, "+"))
	}
	for _, transaction := range issuerTransactions {
		response = append(response, mapper.TransactionToTransactionResponseDTO(transaction, "-"))
	}
	return response, nil
}

func (service *Service) GetAllTransactions() ([]dto.TransactionResponseDTO, error) {
	transactions, err := service.Repository.GetAllTransactions()
	if err != nil {
		return nil, err
	}

	response := make([]dto.TransactionResponseDTO, 0)
	for _, transaction := range transactions {
		response = append(response, mapper.TransactionToTransactionResponseDTO(transaction, ""))
	}
	return response, nil
}
