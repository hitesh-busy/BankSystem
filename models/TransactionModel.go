package models

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Transaction struct {
	TransactionId uint    `pg:",pk"`
	AccountID     uint    `pg:"fk:account_id,on_delete:CASCADE"`
	Mode          string  `pg:",notnull"`
	ReceiverAccNo uint    
	Amount        float64 `pg:",notnull"`
	Account       *Account `pg:"rel:has-one"`
}

func FethchTransactions(db *pg.DB) ([]Transaction, error) {
	var transaction []Transaction
	err := db.Model(&transaction).Select()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
func (transaction *Transaction) FethchTransaction(db *pg.DB, transactionId string) error {
	err := db.Model(transaction).Where("transaction_id=?", transactionId).Select()
	if err != nil {
		fmt.Println("error is ", err)
		return err
	}
	return nil
}
func (transaction *Transaction) Deletetranaction(db *pg.DB, transactionId string) error {

	_, err := db.Model(&transaction).Where("tranaction_id=?", transactionId).Delete()
	if err != nil {
		return err
	}
	return nil
}
func (transaction *Transaction) InsertTransaction(tx *pg.Tx, rawJson map[string]interface{}) error {
	var sender Account
	var receiver Account

	transaction.AccountID = uint(rawJson["AccountID"].(float64))

	var amount = rawJson["Amount"].(float64)

	sender.AccountId = uint(rawJson["AccountID"].(float64))

	//finding if sender account exists or not
	err := tx.Model(&sender).WherePK().Select()
	if err != nil {
		return err
	}

	switch rawJson["Mode"].(string) {
	case "Deposit":
		sender.Balance += amount
		_, err = tx.Model(&sender).WherePK().Update()
		if err != nil {
			return errors.New("issue while Depositing  Money")
		}
		fmt.Println("New Balance is ", sender.Balance)

	case "Withdraw":
		if sender.Balance < amount {
			return errors.New("not enought Balance")
		}
		sender.Balance -= amount
		_, err = tx.Model(&sender).WherePK().Update()
		if err != nil {
			return errors.New("issue while Withdrawing  Money")
		}
	case "Transfer":
		if sender.Balance < amount {
			return errors.New("not enought Balance")
		}
		receiver.AccountId = uint(rawJson["ReceiverAccNo"].(float64))
		transaction.ReceiverAccNo = receiver.AccountId

		//sender.Balance -= amount
		_, err = tx.Model(&sender).WherePK().Set("balance =balance - ?", amount).Update()
		if err != nil {
			return errors.New("issue while Withdrawing  Money from Sender")
		}

		///receiver.Balance += amount
		_, err = tx.Model(&receiver).WherePK().Set("balance =balance + ?", amount).Update()
		if err != nil {
			return errors.New("issue while Deposting  Money to receiver")
		}

	}

	transaction.Amount = amount
	transaction.Mode = rawJson["Mode"].(string)

	_, err = tx.Model(transaction).Insert()
	if err != nil {
		fmt.Println(err)
		return errors.New("issue while inserting Transaction in DB")
	}
	return nil
}
