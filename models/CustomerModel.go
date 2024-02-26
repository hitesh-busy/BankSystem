package models

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
)

type Customer struct {
	CustomerId  uint       `pg:",pk"`
	Account     []*Account `pg:"many2many:customer_accounts"`
	BranchId    uint       `pg:"fk:customer_id,on_delete:CASCADE,on_update:CASCADE"`
	Branch      *Branch    `pg:"rel:has-one"`
	Name        string
	PAN         string
	DOB         time.Time
	Phone       int
	Address     string
	JoiningDate time.Time
	LeavingDate time.Time
}

func FetchCustomers(db *pg.DB) ([]*Customer, error) {
	var customers []*Customer
	err := db.Model(&customers).Select()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (customer *Customer) FetchCustomer(db *pg.DB, customerId string) error {
	err := db.Model(customer).Relation("Branch").Where("customer_id=?", customerId).Select()
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	return nil
}

func (customer *Customer) FetchCustomerAccounts(db *pg.DB, customerId string) error {
	err := db.Model(customer).Relation("Account").Where("customer_id=?", customerId).Select()
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	return nil
}

func (customer *Customer) InsertCustomer(db *pg.DB) error {
	_, err := db.Model(customer).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (customer *Customer) Update(db *pg.DB, customerId string) error {
	_, err := db.Model(customer).Where("customer_id=?", customerId).Update(&customer)
	if err != nil {
		return err
	}
	return nil
}
func (customer *Customer) Delete(db *pg.DB, customerId string) error {
	_, err := db.Model(customer).Where("customer_id=?", customerId).Delete()
	if err != nil {
		return err
	}
	return nil
}
