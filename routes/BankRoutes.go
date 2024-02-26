package routes

import (
	"github.com/BankSystem/controllers"
	"github.com/gin-gonic/gin"
)

func BankRoutes(router *gin.Engine) {

	bankRoutes := router.Group("bank")
	bankRoutes.GET("/", controllers.GetAllBanks)                        //get all bankes
	bankRoutes.GET("/:bank_id", controllers.GetBank)                    //get specific bankes
	bankRoutes.POST("/", controllers.CreateBank)                        //create a bank
	bankRoutes.PATCH("/:bank_id", controllers.UpdateBank)               //update specific bank
	bankRoutes.DELETE("/:bank_id", controllers.DeleteBank)              //delete specific bank
	bankRoutes.GET("/:bank_id/branch", controllers.GetBankWithBranches) //get branches of specific bank
}
