package transactions

import (
	"banking-app/backend/helpers"
	"banking-app/backend/interfaces"
)

func CreateTransaction(From uint, To uint, Amount int) {
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	helpers.DBconnection.Create(&transaction)
}