package routes

import (
	"github.com/BankSystem/controllers"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(router *gin.Engine) {
	customer := router.Group("/customer")
	customer.GET("/", controllers.GetAllCustomers)               //get all customeres
	customer.GET("/:customer_id", controllers.GetCustomer)       //get specific customeres details
	customer.POST("/", controllers.CreateCustomer)               //create a customer
	customer.PATCH("/:customer_id", controllers.UpdateCustomer)  //update specific customer
	customer.DELETE("/:customer_id", controllers.DeleteCustomer) //delete specific customer
	//customer.GET("/:customer_id/branch", controllers.GetCustomerBranches) //get branches of specific customer

	customer.GET("/:customer_id/accounts", controllers.GetCustomerAccount) //get accounts of the customer
	//to be added
}
