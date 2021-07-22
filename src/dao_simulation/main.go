package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Staff struct {
	ID   int
	Name string
}

func (stf *Staff) GetById(id int) *Staff {
	// Execute the query
	err := DB.QueryRow("SELECT id, name FROM staff where id = ?", id).Scan(&stf.ID, &stf.Name)
	if err == sql.ErrNoRows {
		stf.ID = 0
		stf.Name = ""
		return stf
	}
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	return stf
}
func (stf *Staff) IsExist() bool {
	return !(stf.ID == 0 && stf.Name == "")
}

var DB *sql.DB

func main() {
	// Connect to DB
	db, err := sql.Open("mysql", "souukou:12345678@tcp(db4free.net:3306)/souukoudb")
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	defer db.Close()

	// Query
	var the_staff Staff

	//the_staff = the_staff.GetById(3)
	for id := 1; id <= 3; id++ {
		the_staff.GetById(id)
		if the_staff.IsExist() {
			fmt.Println(the_staff.ID, the_staff.Name)
		} else {
			fmt.Println(id, "not exist")
		}
	}

}
