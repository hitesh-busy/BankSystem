package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/gin-gonic/gin"
)

func GetAllBranches(c *gin.Context) {
	branches, err := models.FetchBranches(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not fetch the branches"})
		return
	}
	c.JSON(200, branches)
}
func CreateBranch(c *gin.Context) {
	//Can also extract account from the branch JSON body and create account as well
	//like did for Bank and Branch
	var branch models.Branch
	//var customer []*models.Customer

	err := c.BindJSON(&branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not bind the the JSON "})
		return
	}

	err = validate.Struct(branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Not valid JSON "})
		return
	}
	bankId := strconv.FormatUint(uint64(branch.BankID), 10)

	//check if the corressponding bank exists or not
	err = (&models.Bank{}).FetchBank(database.DB, bankId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Bank not exists for this branch"})
		return
	}

	err = branch.InsertBranch(database.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could insert into table "})
		return
	}

	//NOTE - Could've used transaction here but the idea is, ki jitni branches create ho rahi h hojaye jismein error aaraha h woh nahi hogi bas. And bank toh hoga hi create

	//extracting the customers out of the request body and creating those as well
	for _, v := range branch.Customer {
		fmt.Println("Branch is ", v)
		err := v.InsertCustomer(database.DB)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Could not insert nested branches into DB "})
			return
		}
	}

	c.JSON(200, branch)
}

func UpdateBranch(c *gin.Context) {
	branchId := c.Param("branch_id")

	var branch models.Branch

	//check if branch exists or not
	err := branch.FetchBranch(database.DB, branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "No branch exist with this id"})
		return
	}

	//bind the request body into branch var
	if err := c.BindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not update Branch "})
		return
	}

	err = branch.Update(database.DB, branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Could not update Branch "})
		return
	}
	c.JSON(200, branch)
}

func GetBankOfBranch(c *gin.Context) {
	branchId := c.Param("branch_id")
	var branch models.Branch

	err := database.DB.Model(&branch).Relation("Bank").Where("branch_id=?", branchId).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err ; Could not fethch the bank of Branch. Reason ": err})
		return
	}
	fmt.Println(branch)
	c.JSON(200, gin.H{"branch": branch})

}
func DeleteBranch(c *gin.Context) {

	branchId := c.Param("branch_id")

	err := (&models.Branch{}).Delete(database.DB, branchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "No branch with this id "})
		return
	}
	c.JSON(200, gin.H{"msg": "Branch Successfully Deleted"})
}

func GetCustomersOfBranch(c *gin.Context){
	branchId := c.Param("branch_id")

	
	customers, err := models.GetCustomers(database.DB, branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err":"No customers for this branch "})
		return
	}
	c.JSON(200, customers)
}
