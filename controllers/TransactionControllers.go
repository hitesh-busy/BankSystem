package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transaction []models.Transaction

	transaction, err := models.FethchTransactions(database.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg:NO Transactions  ": err.Error()})
		return
	}
	c.JSON(200, transaction)
}


func GetAllTransaction(c *gin.Context) {
	var transaction models.Transaction
	transactionId := c.Param("transaction_id")

	err := (transaction).FethchTransaction(database.DB,transactionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg:NO Transactions  ": err.Error()})
		return
	}
	c.JSON(200, transaction)
}
func CreateTransaction(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)

	var rawJson map[string]interface{}

	err := json.Unmarshal(body, &rawJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg:": err.Error()})
		return
	}
	//check if account exists for this transaction or not
	err = (&models.Account{}).FetchById(database.DB, strconv.FormatFloat(rawJson["AccountId"].(float64), 'f', -1, 64))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg: Account does not exist for this ": err.Error()})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg: Transaction could not be started": err.Error()})
		return
	}
	var transaction = models.Transaction{}
	err = transaction.InsertTransaction(tx, rawJson)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"msg: error during transactoin": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(200, transaction)

}

func DeleteTransaction(c *gin.Context) {
	//check if the transaction exists
	transactionId := c.Param("transaction_id")

	err := (&models.Transaction{}).FethchTransaction(database.DB, transactionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Tranacion does not exist for this id"})
		return
	}

	//else delete
	err = (&models.Transaction{}).Deletetranaction(database.DB, transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete bank"})
		return
	}
	c.JSON(200, gin.H{"msg": "Delted Successfully"})
}
