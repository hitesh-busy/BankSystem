package controllers

import (
	"fmt"
	"net/http"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/gin-gonic/gin"
)

func CreateMapping(c *gin.Context) {
	type CustomerAccountsPayload struct {
		CustomerAccounts []models.CustomerAccount `json:"customerAccounts"`
	}

	var customer_accounts CustomerAccountsPayload

	err := c.ShouldBindJSON(&customer_accounts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Customer accont is ", customer_accounts.CustomerAccounts)
	for _, v := range customer_accounts.CustomerAccounts {
		err = models.InsertCustomerAccount(database.DB, v)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err ": err.Error()})
			return
		}
	}
	c.JSON(200, customer_accounts)
}
