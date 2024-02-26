package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Bank struct {
	Name     string    `pg:"name"`
	BankId   uint      `pg:"bank_id,pk"`
	Branches []*Branch `pg:"branches,rel:has-many"`
}

func (bank *Bank) InsertBank(db *pg.DB) error {
	_, err := db.Model(bank).Insert()
	if err != nil {
		return err
	}
	return nil
}

func FetchBanks(db *pg.DB) ([]Bank, error) {
	var banks []Bank
	err := db.Model(&banks).Select()
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (bank *Bank) FetchBank(db *pg.DB, bankId string) error {

	err := db.Model(bank).Relation("Branches").Where("bank_id=?", bankId).Select()
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	return nil
}

func (bank *Bank) Delete(db *pg.DB, bankId string) error {

	_, err := db.Model(bank).Where("bank_id=?", bankId).Delete()
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	return nil
}

func (bank *Bank) Update(db *pg.DB, bankId string) error {

	_, err := db.Model(bank).Where("bank_id=?", bankId).Set("bank_id=?", bank.BankId).Update()
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	return nil
}
func BankBranches(db *pg.DB, bankId string) ([]*Branch, error) {
	//var branches []*Branch
	var bank Bank
	// via join query method.
	/*
		it tranalates to
				SELECT branches.*
				FROM branches AS branches
				JOIN banks ON banks.bank_id = branches.bank_id
				WHERE branches.bank_id = <bankId>;
		ismein explictly banana padta h branches []*Branch to capture the solution
	*/
	// err := db.Model(&branches).
	// 	ColumnExpr("banks.*, branches.*").
	// 	Join("JOIN banks ON banks.bank_id = branches.bank_id").
	// 	TableExpr("branches AS branches").
	// 	Where("branches.bank_id = ?", bankId).
	// 	Select()

	//Relation method
	/*
		SELECT bank.*, branches.*
		FROM bank AS bank
		LEFT JOIN branches AS branches ON branches.bank_id = bank.bank_id
		WHERE bank.bank_id = <bankId>;
	
	REMEMBER:- Relation ke andar woh FIELD VALUE jaati h jo Bank model ke andar h. it is not the Model name nor the tables name
	*/
	err := db.Model(&bank).Relation("Branches").Where("bank_id=?", bankId).Select()
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return bank.Branches, nil
}
