package transaction

import (
	"github.com/BankSystem/transaction/controller"
	"github.com/gin-gonic/gin"
)

func MountRoutes(router *gin.RouterGroup) {
	transactions := router.Group("transaction")
	transactions.GET("/", controller.GetAllTransactions)
	transactions.GET("/:transaction_id", controller.GetAllTransaction)
	transactions.POST("/", controller.CreateTransaction)
	transactions.DELETE("/:transaction_id", controller.DeleteTransaction)

}
