package service

import (
	"fmt"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
)

func CalculateTransactionsForTheDay() string{
	var totalCredit float64 = 0.0
	transactions, err := models.FethchTransactions(database.DB)
	if(err != nil){
		fmt.Println("Some problem fetching trnsactions ", err)
	}

	if(transactions == nil){
		fmt.Println("There are no transactions ")
	}

	
	for _, val := range(transactions){
		fmt.Println("the transactions is ", val)
		totalCredit += val.Amount
	}

	//can send even complext html data
	message := fmt.Sprintf("Total transaction is %v credits", totalCredit)
	return message
}	
