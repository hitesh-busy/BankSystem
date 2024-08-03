package account

import (
	"github.com/BankSystem/account/controller"
	"github.com/gin-gonic/gin"
)

func MountRoutes(router *gin.RouterGroup) {
	account := router.Group("/account")

	{
		account.GET("/", controller.GetAllAccounts)              //get all accountes
		account.GET("/:account_id", controller.GetAccount)       //get specific accountes details
		account.POST("/", controller.CreateAccount)              //create a account
		account.PATCH("/:account_id", controller.UpdateAccount)  //update specific account
		account.DELETE("/:account_id", controller.DeleteAccount) //delete specific account
		account.POST("/deposit", controller.DepositMoney)        //deposit money in specific amount
		account.POST("/withdraw", controller.WithdrawMoney)      //withdrawe money in specific amount
		account.GET("/:account_id/transactions", controller.GetTransactions)
	}
}
