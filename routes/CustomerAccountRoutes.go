package routes

import (
	"github.com/BankSystem/controllers"
	"github.com/gin-gonic/gin"
)

func CustomerAccountRoutes(router *gin.Engine) {
	customerAccount := router.Group("/customerAccount")
	//customerAccount.GET("/", controllers.GetAllCustomers)               //get all customerAccountes
	//customerAccount.GET("/:customerAccountId", controllers.GetCustomer)       //get specific customerAccountes details
	customerAccount.POST("/", controllers.CreateMapping)               //create a customerAccount
	// customerAccount.PATCH("/:customerAccount_id", controllers.UpdateCustomer)  //update specific customerAccount
	// customerAccount.DELETE("/:customerAccount_id", controllers.DeleteCustomer) //delete specific customerAccount

	// customerAccount.GET("/:customerAccount_id/accounts", controllers.GetCustomerAccount) 
}
