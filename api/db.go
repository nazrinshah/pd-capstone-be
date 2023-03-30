package api

import (
	"database/sql"
	"encoding/json"
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

func (d *DB) CreateOrder(order Order) (Order, error) {

	itemStr, _ := json.Marshal(order.Items)
	sql := fmt.Sprintf("INSERT INTO %s(orderstatus, items, subtotal, platformfee, deliveryfee) VALUES (%d, '%s', %.2f, %.2f, %.2f)", TABLE_ORDER, order.OrderStatus, itemStr, order.Subtotal, order.PlatformFee, order.DeliveryFee)
	res, err := d.db.Exec(sql)

	fmt.Println(fmt.Sprintf("%+v", res))

	i, err := res.LastInsertId()
	order.Id = uint64(i)

	return order, err
}

func (d *DB) RetrieveDishes() ([]Dish, error) {
	res := []Dish{}

	query := fmt.Sprintf("SELECT * FROM %s", TABLE_DISH)

	rows, err := d.db.Query(query)
	defer rows.Close()

	for rows.Next() {
		temp := Dish{}
		err = rows.Scan(&temp.Id, &temp.VendorId, &temp.Name, &temp.Status, &temp.Price, &temp.Description, &temp.Currency, &temp.ImageName)

		if err != nil {
			break
		}

		res = append(res, temp)
	}

	return res, err
}

func (d *DB) RetrieveDishById(id uint64) (Dish, error) {
	res := Dish{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %d = ?", TABLE_DISH, id)
	err := d.db.QueryRow(query, id).Scan(&res.Id, &res.VendorId, &res.Name, &res.Status, &res.Price, &res.Description, &res.Currency, &res.ImageName)

	if err != nil {
		fmt.Println("Error in RetrieveDishById: ", err)
	}

	return res, err
}

func (d *DB) Init() error {
	var err error
	d.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_SCHEMA))

	return err
}

func (d *DB) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
