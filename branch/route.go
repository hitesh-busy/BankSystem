package branch

import (
	"github.com/gin-gonic/gin"
	"github.com/BankSystem/branch/controller"
)

func MountRoutes(router *gin.RouterGroup) {

	branch := router.Group("/branch")
	branch.GET("/", controller.GetAllBranches)                           //get all branches
	branch.POST("/", controller.CreateBranch)                            //create a branch
	branch.PATCH("/:branch_id", controller.UpdateBranch)                 //update specific branch
	branch.GET("/:branch_id/bank", controller.GetBankOfBranch)           //get bank of specific branch
	branch.DELETE("/:branch_id/", controller.DeleteBranch)               //DElete  specific branch
	branch.GET("/:branch_id/customers", controller.GetCustomersOfBranch) //get all customers of specific branch
}
