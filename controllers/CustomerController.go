package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/gin-gonic/gin"
)

func GetAllCustomers(c *gin.Context) {

	result, err := models.FetchCustomers(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not fetch the customers"})
		return
	}
	c.JSON(200, result)
}

func GetCustomer(c *gin.Context) {
	var customer models.Customer
	customerId := c.Param("customer_id")
	err := customer.FetchCustomer(database.DB, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not fethch the brances "})
		return
	}
	c.JSON(200, customer)

}

func GetCustomerAccount(c *gin.Context) {
	customerId := c.Param("customer_id")
	var customer models.Customer
	err := customer.FetchCustomerAccounts(database.DB, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not fethch the brances "})
		return
	}
	c.JSON(200, customer)

}
func CreateCustomer(c *gin.Context) {

	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not bind the the JSON "})
		return
	}

	//begin a transactoin to take care of multiple isnertions in Customer, account and CustomerAccount tables
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Err ": err.Error()})
		return
	}

	err = validate.Struct(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Not valid JSON "})
		return
	}
	branchId := strconv.FormatUint(uint64(customer.BranchId), 10)

	//check if the corressponding branch exists or not
	err = (&models.Branch{}).FetchBranch(database.DB, branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Branch not exists for this customer"})
		return
	}

	customer.JoiningDate = time.Now()

	err = customer.InsertCustomer(database.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could insert into table "})
		return
	}
	c.JSON(200, customer)

	for _, v := range customer.Account {
		err = v.Insert(tx)
		if err != nil {
			rollbackErr := tx.Rollback()
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"err Could insert into Customer Account ": rollbackErr.Error()})
			return
		}
	}
	tx.Commit()
	c.JSON(200, customer)

}

func UpdateCustomer(c *gin.Context) {
	customerId := c.Param("customer_id")

	var customer models.Customer

	//check if customer exists or not
	err := customer.FetchCustomer(database.DB, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "No customer exist with this id"})
		return
	}

	//bind the request body into customer var
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not bind JSON "})
		return
	}

	err = customer.Update(database.DB, customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not update customer "})
		return
	}
	c.JSON(200, customer)
}

func DeleteCustomer(c *gin.Context) {

	customerId := c.Param("customer_id")

	err := (&models.Customer{}).Delete(database.DB, customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "No customer with this id "})
		return
	}
	c.JSON(200, gin.H{"msg": "customer Successfully Deleted"})
}
