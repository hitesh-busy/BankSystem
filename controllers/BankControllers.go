package controllers

import (
	"fmt"
	"net/http"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

/*
A receiver base type cannot be a pointer or interface type and it must be defined in the same package as the method.
*/
func GetAllBanks(c *gin.Context) {

	bankResult, err := models.FetchBanks(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not fethch the banks "})
		return
	}
	c.JSON(200, bankResult)
}

func GetBank(c *gin.Context) {
	var bank models.Bank
	bankId := c.Param("bank_id")
	//bid, _ := strconv.ParseUint(bankId, 10, 0)

	err := bank.FetchBank(database.DB, bankId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "No Bank exist with this id"})
		return
	}
	c.JSON(200, bank)
}

func CreateBank(c *gin.Context) {
	var bank models.Bank

	//getting request body
	err := c.BindJSON(&bank)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not bind the the JSON "})
		return
	}

	//valdating JSON
	err = validate.Struct(&bank)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Not valid JSON "})
		return
	}

	//calling the query via bank method
	err = bank.InsertBank(database.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not insert bank into DB "})
		return
	}

	//extracting the branches out of the request body and creating those as well
	for _, v := range bank.Branches {
		fmt.Println("Branch is ", v)
		err := v.InsertBranch(database.DB)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Could not insert nested branches into DB "})
			return
		}
	}

	c.JSON(200, bank)
}

func UpdateBank(c *gin.Context) {
	//for updating we have to keep in mind that all the branches associated with this bank also need to be updates
	//so a transaction will be ruu to complete operation
	bankId := c.Param("bank_id")

	//bank should exist
	var bank models.Bank
	err := bank.FetchBank(database.DB, bankId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "No Bank exist with this id"})
		return
	}

	// BindJSON will only update the fields present in the request body
	if err := c.BindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bank.Update(database.DB, bankId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not update bank "})
		return
	}
	c.JSON(200, bank)

}
func GetBankWithBranches(c *gin.Context) {
	var branches []*models.Branch
	bankId := c.Param("bank_id")

	//var details []BankDetails
	branches, err := models.BankBranches(database.DB, bankId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "No branches assocaited with this bank"})
		return
	}
	fmt.Print("The result is ", branches)

	c.JSON(200, branches)
}

func DeleteBank(c *gin.Context) {
	//check if the bank exists
	bankId := c.Param("bank_id")

	err := database.DB.Model(&models.Bank{}).Where("bank_id=?", bankId).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Bank does not exist for this id"})
		return
	}

	//else delete
	result, err := database.DB.Model(&models.Bank{}).Where("bank_id= ?", bankId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete bank"})
		return
	}
	c.JSON(200, result.RowsAffected())
}
