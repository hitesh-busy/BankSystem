package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
)
type Account struct {
	Transactions []*Transaction `pg:"rel:has-many" json:"transactions"`
	AccountId    uint           `pg:",pk" json:"account_id"`
	AccountNo    uint           `json:"account_no"`
	Balance      float64        `json:"balance"`
	Type         string         `json:"type"`
	BranchId     uint           `pg:"fk:branch_id,on_delete:SET NULL" json:"branch_id"`
	Branch       *Branch        `pg:"fk:branch_id rel:has-one" json:"branch"`
	OpeningDate  time.Time      `json:"opening_date"`
	CustomerAcc  []*Customer    `pg:"many2many:customer_accounts" json:"customer_accounts"`
}

func Fetch(db *pg.DB) ([]*Account, error) {
	var accounts []*Account
	err := db.Model(&accounts).Select()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (account *Account) FetchById(db *pg.DB, accountId string) error {
	//Relation method only works for has-one, and has-many relations

	//basically ismein humne kya kara ki bhai particualr id pe jo account h woh toh aayega hi uska owner customer bhi aayega
	//this could be achieveed via CustomerAccount Model

	
	err := db.Model(account).Relation("CustomerAcc").Where("account_id=?", accountId).Select()
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	fmt.Println("Insude fetch by Id")
	return nil
}

// func (account *Account) Insert(db *pg.DB) error {
// 	_, err := db.Model(account).Insert()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (account *Account) Insert(param interface{}) error {

	if tx, ok := param.(*pg.Tx); ok {
		// If a *pg.Tx is provided, use it as the executor
		fmt.Println("this is transaction  Tx")
		_, err := tx.Model(account).Insert()
		if err != nil {
			return err
		}
	}
	if db, ok := param.(*pg.DB); ok {
		fmt.Println("this is normal DB")
		_, err := db.Model(account).Insert()
		if err != nil {
			return err
		}
	}
	return nil

}
func (account *Account) FetchTransactions(db *pg.DB, accountId string) error{
	err := db.Model(account).Relation("Transactions").Where("account_id= ?",accountId).Select()
	if err != nil {
		return err
	}
	return nil
}
func (account *Account) Update(db *pg.DB, accountId string) error {
	_, err := db.Model(account).Where("account_id=?", accountId).Update()
	if err != nil {
		return err
	}
	return nil
}

func Deposit(tx *pg.Tx, rawJSON map[string]interface{}) (float64, error) {

	amount := rawJSON["Amount"]

	var account Account

	account.AccountId = uint(rawJSON["AccountId"].(float64))

	err := tx.Model(&account).WherePK().Select()
	if err != nil {
		tx.Rollback()
		fmt.Println("error is ", err)
		return -1, err
	}

	account.Balance += amount.(float64)

	_, err = tx.Model(&account).WherePK().Update()
	if err != nil {
		tx.Rollback()
		fmt.Println("error is ", err)
		return account.Balance, err
	}

	return account.Balance, err
}

func Withdraw(tx *pg.Tx, rawJSON map[string]interface{}) (float64, error) {

	amount := rawJSON["Amount"].(float64)

	var account Account
	account.AccountId = uint(rawJSON["AccountId"].(float64))

	err := tx.Model(&account).WherePK().Select()
	if err != nil {
		tx.Rollback()
		fmt.Println("error is ", err)
		return -1, err
	}
	if account.Balance < amount {
		return account.Balance, errors.New("not enough Balance ")
	}
	account.Balance -= amount

	_, err = tx.Model(&account).WherePK().Update()
	if err != nil {
		tx.Rollback()
		fmt.Println("error is ", err)
		return account.Balance, err
	}

	return account.Balance, err
}
