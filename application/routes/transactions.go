package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"transactions/application/middlewares"
	"transactions/domain/entities"
	"transactions/domain/usecases"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction entities.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		log.Println("[Request Body Error]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction.AccountID = middlewares.CurrentAccountID
	transaction, err := usecases.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}
