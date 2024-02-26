package models

import (
	"github.com/go-pg/pg/v10"
)

type CustomerAccount struct {
	CustAccId  uint      `pg:",pk"`
	AccountId  uint      `pg:",on_delete:CASCADE"`
	Account    *Account  `pg:"rel:has-one"`
	CustomerId uint      `pg:"join_fk:customer_id"`
	Customer   *Customer `pg:"rel:has-one"`
}

//customer jab bhi create ho raha h tab hum yahan explictly entry daal rahe h uski jise hum later on .Relation mein use kar sake
//joins bhi tabhi lagega jab beech mein kuch common table hoga
//Relation func backend mein apne aap kaam karleta h relate karne ka
//ALternate would be to left join the desired column with the CustomerAccount table
//for instace to get the Accounts of the Customer we could perform left join the account table and the CustomerAccount table

func InsertCustomerAccount(db *pg.DB, customer_accounts CustomerAccount) error {

	_, err := db.Model(&customer_accounts).Insert()
	if err != nil {
		return err
	}
	return nil

}
