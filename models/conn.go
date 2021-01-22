package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connet() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/go_restapi")
	return db, err
}
