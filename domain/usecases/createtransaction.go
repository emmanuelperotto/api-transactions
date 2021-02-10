package usecases

import (
	"errors"
	"log"
	"transactions/domain/entities"
	"transactions/infra/repositories"
)

func CreateTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	if transaction.AccountID == 0 {
		log.Println("[CreateTransaction Error] AccountID can't be null/empty")
		return transaction, errors.New("AccountID can't be null/empty")
	}

	transaction, err := repositories.TransactionRepo.Create(transaction)

	if err != nil {
		log.Println("[CreateTransaction Error]", err)
		return transaction, err
	}

	log.Println("Transaction Created")
	return transaction, nil
}
