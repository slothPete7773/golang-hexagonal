package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
	// a int
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (a accountRepositoryDB) Create(account Account) (*Account, error) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (:customer_id, :opening_date, :account_type, :amount, :status) RETURNING account_id"

	result, err := a.db.NamedExec(query, account)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %v", err)
	}

	accountId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	account.AccountID = int(accountId)
	return &account, nil
}
func (a accountRepositoryDB) GetAll(customerId int) ([]Account, error) {
	var accounts []Account
	query := "SELECT * FROM accounts WHERE customer_id = ?"
	err := a.db.Select(&accounts, query, customerId)
	if err != nil {
		return nil, fmt.Errorf("error fetching accounts: %v", err)
	}
	return accounts, nil
}
