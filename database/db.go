package database

import (
	"log"

	"github.com/BankSystem/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func ConnectToDB() {
	// Connect to PostgreSQL database
	DB = pg.Connect(&pg.Options{
		User:     "hiteshsharma",
		Password: "1234",
		Database: "Bank1",
		Addr:     "localhost:5432",
	})
	if DB == nil {
		log.Fatalln("Could not connect to the database ")
		return
	}
	log.Println("Connection to DB successful")

	// Create tables
	err := createSchema()
	if err != nil {
		log.Fatal(err)
	}

}
func createSchema() error {
	models := []interface{}{
		(*models.Bank)(nil),
		(*models.Branch)(nil),
		(*models.Account)(nil),
		(*models.Transaction)(nil),
		(*models.Customer)(nil),
		(*models.CustomerAccount)(nil),
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
		log.Println("Tables succesffuly created ")
	}
	return nil
}
