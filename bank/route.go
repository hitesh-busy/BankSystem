package bank

import (
	"github.com/BankSystem/bank/controller"
	"github.com/gin-gonic/gin"
)

func MountRoutes(router *gin.RouterGroup) {

	bankRoutes := router.Group("bank")
	bankRoutes.GET("/", controller.GetAllBanks)                        //get all bankes
	bankRoutes.GET("/:bank_id", controller.GetBank)                    //get specific bankes
	bankRoutes.POST("/", controller.CreateBank)                        //create a bank
	bankRoutes.PATCH("/:bank_id", controller.UpdateBank)               //update specific bank
	bankRoutes.DELETE("/:bank_id", controller.DeleteBank)              //delete specific bank
	bankRoutes.GET("/:bank_id/branch", controller.GetBankWithBranches) //get branches of specific bank
}
