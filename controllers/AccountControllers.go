package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/gin-gonic/gin"
)

/*
A receiver base type cannot be a pointer or interface type and it must be defined in the same package as the method.
*/
func GetAllAccounts(c *gin.Context) {

	accountResult, err := models.Fetch(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not fethch the accounts "})
		return
	}
	c.JSON(200, accountResult)
}

func GetAccount(c *gin.Context) {
	var account models.Account
	accountId := c.Param("account_id")

	err := account.FetchById(database.DB, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err - No Account exist with this id Reason": err})
		return
	}
	c.JSON(200, account)
}

func GetTransactions(c *gin.Context){
	accountId := c.Param("account_id")
	var account models.Account
	err :=  (&account).FetchById(database.DB, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err No Account exist with this id Reason ": err})
		return
	}

	err = account.FetchTransactions(database.DB, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err No Transactions for this account": err})
		return

	}
	c.JSON(200, account)
}

func CreateAccount(c *gin.Context) {
	var account models.Account

	//getting request body
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not bind the the JSON "})
		return
	}

	//valdating JSON
	err = validate.Struct(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Not valid JSON "})
		return
	}

	//calling the query via account method
	err = account.Insert(database.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err Could not insert account into DB Reason ": err.Error()})
		return
	}

	//extracting the branches out of the request body and creating those as well
	// for _, v := range account.CustomerAcc {
	// 	fmt.Println("Customer is ", v)
	// 	err := v.InsertCustomer(database.DB)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"err Could not insert nested customers into DB Reason ": err})
	// 		return
	// 	}
	// }

	c.JSON(200, account)
}

func UpdateAccount(c *gin.Context) {
	//for updating we have to keep in mind that all the branches associated with this account also need to be updates
	//so a transaction will be ruu to complete operation
	accountId := c.Param("account_id")

	//account should exist
	var account models.Account
	err := account.FetchById(database.DB, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err No Account exist with this id Reason ": err})
		return
	}

	// BindJSON will only update the fields present in the request body
	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = account.Update(database.DB, accountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not update account "})
		return
	}
	c.JSON(200, account)

}
func GetAccountWithBranches(c *gin.Context) {
	var branches []*models.Branch
	//accountId := c.Param("account_id")

	//var details []AccountDetails
	// branches, err := models.AccountBranches(database.DB, accountId)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"err": "No branches assocaited with this account"})
	// 	return
	// }
	fmt.Print("The result is ", branches)

	c.JSON(200, branches)
}

func DeleteAccount(c *gin.Context) {
	//check if the account exists
	accountId := c.Param("account_id")

	err := database.DB.Model(&models.Account{}).Where("account_id=?", accountId).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Account does not exist for this id"})
		return
	}

	//else delete
	result, err := database.DB.Model(&models.Account{}).Where("account_id= ?", accountId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete account"})
		return
	}
	c.JSON(200, result.RowsAffected())
}

func DepositMoney(c *gin.Context) {
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle the error (e.g., log it or return an error response)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Parse the raw JSON data into a json.RawMessage
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(body, &rawJSON); err != nil {
		// Handle the error (e.g., log it or return an error response)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}
	fmt.Println(rawJSON)

	//deposit will be a part of transaction
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatalln("Error Begginning transaction ")
		c.JSON(http.StatusInternalServerError, gin.H{"err Occuered": err})
	}

	newAmount, err := models.Deposit(tx, rawJSON)
	if err != nil {
		tx.Rollback()
		log.Fatalln("Error Performing function ")
		c.JSON(http.StatusInternalServerError, gin.H{"err Occuered": err})
	}

	//also have to update the transaction details
	//err = models.InsertTransaction(tx, rawJSON)


	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatalln("Error Committing ")
		c.JSON(http.StatusInternalServerError, gin.H{"err Occuered": err})
	}
	tx.Close()

	c.JSON(200, gin.H{"New Balance is ": newAmount})
}

func WithdrawMoney(c *gin.Context) {
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle the error (e.g., log it or return an error response)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Parse the raw JSON data into a json.RawMessage
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(body, &rawJSON); err != nil {
		// Handle the error (e.g., log it or return an error response)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}
	fmt.Println(rawJSON)

	//withdraw will be a part of transaction
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatalln("Error Begginning transaction ")
		c.JSON(http.StatusInternalServerError, gin.H{"err Occuered": err})
	}

	newAmount, err := models.Withdraw(tx, rawJSON)
	if err != nil {
		// Error occurred during withdrawal
		c.JSON(http.StatusInternalServerError, gin.H{"err Occurred": err.Error(), "Balance": newAmount})
		
		// Rollback the transaction
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalln("Error Rolling back transaction ", rollbackErr)
		}
		
		return
	}
	err = tx.Commit()
	if err != nil {
		// Rollback the transaction
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalln("Error Rolling back transaction ", rollbackErr)
		}
		log.Fatalln("Error Committing ")
		c.JSON(http.StatusInternalServerError, gin.H{"err Occuered": err.Error()})
	}
	tx.Close()

	c.JSON(200, gin.H{"New Balance is ": newAmount})
}
