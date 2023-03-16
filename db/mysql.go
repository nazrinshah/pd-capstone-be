package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:Jason@123@tcp(127.0.0.1:3306)/restaurant")
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("==========Successfully connected to the database==========")
}

func CloseDB() {
	DB.Close()
}

func InsertDB() {
	insert, err := DB.Query("INSERT INTO restaurant.vendor VALUES ( 5, 'vendor_name', 'test', '10:00-15:00')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func GetDB() {
	rows, err := DB.Query("SELECT vendor_id, vendor_name, status, opening_hours FROM restaurant.vendor")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var vendor_id int
		var vendor_name string
		var status string
		var opening_hours string
		err := rows.Scan(&vendor_id, &vendor_name, &status, &opening_hours)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Status: %s, Opening_hours:%s\n", vendor_id, vendor_name, status, opening_hours)
	}
}
