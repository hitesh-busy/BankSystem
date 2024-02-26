package main

import (
	"fmt"

	"github.com/BankSystem/database"
	"github.com/BankSystem/models"
	"github.com/BankSystem/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10/orm"
)

func init() {
	// Register many to many model so ORM can better recognize m2m relation.
	// This should be done before dependant models are used.
	//jo intermediatary table h woh yahan aata h 
	orm.RegisterTable((*models.CustomerAccount)(nil))
	database.ConnectToDB()
}

func main() {
	fmt.Println("Bank Assignemtn with gin and go pg")

	router := gin.Default()

	routes.BankRoutes(router)
	routes.BranchRoutes(router)
	routes.CustomerRoutes(router)
	routes.AccountRoutes(router)
	routes.CustomerAccountRoutes(router)
	routes.TransactionRoutes(router)

	router.Run(":3000")

}
