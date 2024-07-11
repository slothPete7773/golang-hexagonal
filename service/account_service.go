package service

import (
	"bank/repository"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{
		accRepo: accRepo,
	}
}
func (a accountService) NewAccount(customerId int, newAccount NewAccountRequest) (*AccountResponse, error) {
	account := repository.Account{
		CustomerID:  customerId,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      1,
	}

	newAcc, err := a.accRepo.Create(account)
	if err != nil {
		return nil, err
	}

	resAcc := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &resAcc, nil
}
func (a accountService) GetAccounts(customerId int) ([]AccountResponse, error) {
	accounts, err := a.accRepo.GetAll(customerId)
	if err != nil {
		return nil, err
	}

	responses := []AccountResponse{}

	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}
