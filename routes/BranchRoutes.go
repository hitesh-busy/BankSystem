package routes

import (
	"github.com/BankSystem/controllers"
	"github.com/gin-gonic/gin"
)

func BranchRoutes(router *gin.Engine) {

	branch := router.Group("/branch")
	branch.GET("/", controllers.GetAllBranches)                           //get all branches
	branch.POST("/", controllers.CreateBranch)                            //create a branch
	branch.PATCH("/:branch_id", controllers.UpdateBranch)                 //update specific branch
	branch.GET("/:branch_id/bank", controllers.GetBankOfBranch)           //get bank of specific branch
	branch.DELETE("/:branch_id/", controllers.DeleteBranch)               //DElete  specific branch
	branch.GET("/:branch_id/customers", controllers.GetCustomersOfBranch) //get all customers of specific branch
}
