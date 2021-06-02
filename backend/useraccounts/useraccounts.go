package useraccounts

import (
	"banking-app/backend/helpers"
	"banking-app/backend/interfaces"
	"banking-app/backend/transactions"
	"fmt"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	//account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}
	//helpers.DBconnection.Where("id = ? ", id).First(&account)
	//account.Balance = uint(amount)
	//helpers.DBconnection.Save(&account)
	//responseAcc.ID = account.ID
	//responseAcc.Name = account.Name
	//responseAcc.Balance = int(account.Balance)
	return responseAcc
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	//database.DBconnection.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)
	//if helpers.DBconnection.Where("id = ? ", id).First(&account).RecordNotFound() {
	//	return nil
	//}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
	fromAccount := getAccount(from)
	toAccount := getAccount(to)
	if fromAccount == nil || toAccount == nil {
		return map[string]interface{}{"message": "Account not found"}
	} else if fromAccount.UserID != userId {
		return map[string]interface{}{"message": "You are not owner of the account"}
	} else if int(fromAccount.Balance) < amount {
		return map[string]interface{}{"message": "Account balance is not enough"}
	}
	updatedAccount := updateAccount(from, int(fromAccount.Balance) - amount)
	updateAccount(to, int(toAccount.Balance) + amount)
	transactions.CreateTransaction(from, to, amount)
	var response = map[string]interface{}{"message": "all is fine"}
	response["data"] = updatedAccount
	return response
}
