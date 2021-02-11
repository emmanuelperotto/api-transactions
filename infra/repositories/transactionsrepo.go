package repositories

import (
	"context"
	"transactions/domain/entities"
	"transactions/infra"
)

type transaction struct{}

type TransactionCreator interface {
	Create(entities.Transaction) (entities.Transaction, error)
}

func (t transaction) Create(transaction entities.Transaction) (entities.Transaction, error) {
	result, err := infra.DB.ExecContext(context.Background(),
		"INSERT INTO api_transactions.Transactions (AmountCents, AccountID) VALUES (?, ?)",
		transaction.AmountCents, transaction.AccountID)

	if err != nil {
		return transaction, err
	}

	id, _ := result.LastInsertId()
	transaction.ID = id

	return transaction, nil
}

var (
	TransactionRepo = transaction{}
)