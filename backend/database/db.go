package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var DB *sql.DB
var err error

func InitDB() error {
	var db *sql.DB
	db, err = sql.Open("mysql",
		"root:@(127.0.0.1:3306)/go_users?charset=utf8&parseTime=true")
	if err != nil {
		return err
	}
	DB = db
	//	defer DB.Close()
	return nil
}
