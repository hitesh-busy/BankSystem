package models

import "github.com/go-pg/pg/v10"

/*
cascade not working deleted using transactions
*/
/*
pg:"fk:user_id":
This tag is used to define a regular foreign key relationship.
It signifies that the column (user_id in this case) is a foreign key that references the primary key of another table.

pg:"join_fk:user_id":
This tag is used in a many-to-many relationship to indicate that the column (user_id in this case) is part of a join table.
It signifies that the column is a foreign key referencing another table in a join table used to establish a many-to-many relationship.
The join_fk tag is typically used when dealing with join tables.
*/
type Branch struct {
	BranchId uint        `pg:"branch_id, type:serial,pk"`
	Name     string      `pg:"name"`
	Address  string      `pg:"address"`
	BankID   uint        `pg:",on_delete:CASCADE, on_update:CASCADE"`
	Bank     *Bank       `pg:"rel:has-one"`
	Customer []*Customer `pg:"rel:has-many"`
}

func (branch *Branch) InsertBranch(db *pg.DB) error {
	_, err := db.Model(branch).Insert()
	if err != nil {
		return err
	}
	return nil
}

func FetchBranches(db *pg.DB) ([]Branch, error) {
	var branches []Branch
	err := db.Model(&branches).Select()
	if err != nil {
		return nil, err
	}
	return branches, nil
}
func (branch *Branch) FetchBranch(db *pg.DB, branchId string) error {
	err := db.Model(branch).Where("branch_id=?", branchId).Select()
	if err != nil {
		return err
	}
	return nil
}

func (branch *Branch) Update(db *pg.DB, branchId string) error {
	_, err := db.Model(branch).Where("branch_id=?", branchId).Update(&branch)
	if err != nil {
		return err
	}
	return nil
}
func (branch *Branch) Delete(db *pg.DB, branchId string) error {
	_, err := db.Model(branch).Where("branch_id=?", branchId).Delete()
	if err != nil {
		return err
	}
	return nil
}

func GetCustomers(db *pg.DB, branchId string) ([]Customer, error){
	var customers []Customer
	err := db.Model(&customers).Where("branch_id=?",branchId).Select()
	if err != nil {
		return nil, err
	}
	return customers, nil

}