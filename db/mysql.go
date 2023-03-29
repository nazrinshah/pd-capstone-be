package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	VENDOR_TABLE = "vendor"
	DB_SCHEMA    = "restaurant"
)

type Vendor struct {
	Vendor_ID     int    `json:"vendor_id"`
	Vendor_Name   string `json:"vendor_name"`
	Status        string `json:"status"`
	Opening_Hours string `json:"opening_hours"` // hh:mm-hh:mm format, assume open 7 days a week
	//Dish_Name     []Dish
}

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

func CreateVendor(v Vendor) {
	// insert sql statement
	query := fmt.Sprintf("INSERT INTO `%s` (`vendor_name`, `status`, `opening_hours`) VALUES(?, ?, ?)", VENDOR_TABLE)

	fmt.Println("query createvendor:", query)
	_, err := DB.Exec(query, v.Vendor_Name, v.Status, v.Opening_Hours)
	if err != nil {
		panic(err.Error())
	}
}

func GetDB() {
	rows, err := DB.Query("SELECT vendor_id, vendor_name, status, opening_hours FROM %s.%s", DB_SCHEMA, VENDOR_TABLE)
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
