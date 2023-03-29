package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_SCHEMA   = "capstone"
	DB_USERNAME = "root"
	DB_PASSWORD = "P@ss1234"
	TABLE_DISH  = "dish_tab"
	TABLE_ORDER = "order_tab"
)

type DB struct {
	db *sql.DB
}

func (d *DB) Init() error {
	d.db, err := sql.Open("sql", fmt.Sprintf("%s:%s@tcp(%s:3066)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_SCHEMA))

	return err
}