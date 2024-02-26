package routes

import (
	"github.com/BankSystem/controllers"
	"github.com/gin-gonic/gin"
)

func AccountRoutes(router *gin.Engine) {
	account := router.Group("/account")
	account.GET("/", controllers.GetAllAccounts)                        //get all accountes
	account.GET("/:account_id", controllers.GetAccount)                //get specific accountes details
	account.POST("/", controllers.CreateAccount)                        //create a account
	account.PATCH("/:account_id", controllers.UpdateAccount)           //update specific account
	account.DELETE("/:account_id", controllers.DeleteAccount)          //delete specific account
	account.POST("/deposit", controllers.DepositMoney)	//deposit money in specific amount
	account.POST("/withdraw", controllers.WithdrawMoney)	//withdrawe money in specific amount
	account.GET("/:account_id/transactions", controllers.GetTransactions)
}
