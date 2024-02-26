package routes

import (
	"github.com/BankSystem/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	transactions := router.Group("transaction")
	transactions.GET("/", controllers.GetAllTransactions)
	transactions.GET("/:transaction_id", controllers.GetAllTransaction)
	transactions.POST("/", controllers.CreateTransaction)
	transactions.DELETE("/:transaction_id", controllers.DeleteTransaction)

}
