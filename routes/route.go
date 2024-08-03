package routes

import (
	"github.com/BankSystem/account"
	"github.com/BankSystem/bank"
	"github.com/BankSystem/branch"
	"github.com/BankSystem/customer"
	"github.com/BankSystem/transaction"
	"github.com/gin-gonic/gin"
)

// MountRoutes : Initialise All Routes
func MountRoutes(r *gin.Engine) {
	//mountAuthRoutes(r)
	mountHttpRoutes(r)
}

func mountAuthRoutes(r *gin.Engine) {
	//route := r.Group("auth")

	{
		// route.Use(authPkg.HttpTokenValidator(func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(conf.String("auth.jwt_secret", "")), nil
		// }, "auth"))
		//authPkg.Validator(r)
		//route.Use(middleware.ClientValidator())
		// route.Use(middleware.TokenVlidator())

		//auth.MountRoutes(route)
	}
}

func mountHttpRoutes(r *gin.Engine) {

	route := r.Group("bank")
	{
		// route.Use(authPkg.HttpTokenValidator(func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(conf.String("auth.jwt_secret", "")), nil
		// }, "crm"))
		//route.Use(middleware.TokenVlidator())

		account.MountRoutes(route)
		bank.MountRoutes(route)
		customer.MountRoutes(route)
		transaction.MountRoutes(route)
		branch.MountRoutes(route)

	}
}
