package customer

import (
	"github.com/BankSystem/customer/controller"
	"github.com/gin-gonic/gin"
)

func MountRoutes(router *gin.RouterGroup) {
	customer := router.Group("/customer")
	customer.GET("/", controller.GetAllCustomers)               //get all customeres
	customer.GET("/:customer_id", controller.GetCustomer)       //get specific customeres details
	customer.POST("/", controller.CreateCustomer)               //create a customer
	customer.PATCH("/:customer_id", controller.UpdateCustomer)  //update specific customer
	customer.DELETE("/:customer_id", controller.DeleteCustomer) //delete specific customer
	//customer.GET("/:customer_id/branch", controller.GetCustomerBranches) //get branches of specific customer

	customer.GET("/:customer_id/accounts", controller.GetCustomerAccount) //get accounts of the customer
	//to be added
}
