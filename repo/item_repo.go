package repo

import (
	"golang-web-demo/model"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const DB_HOST = "localhost"
const DB_PORT = 3306
const DB_NAME = "golang_web_demo"
const DB_USER = "kevin"
const DB_PWD = "1234"
const DB_CHARSET = "utf8"

func getDb() (*sql.DB, error) {
	dbUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v",
		DB_USER, DB_PWD, DB_HOST, DB_PORT, DB_NAME, DB_CHARSET)
	return sql.Open("mysql", dbUrl)
}

func getItem(idx int) (model.Item) {
	return model.Item{idx, fmt.Sprintf("item_%v", idx)}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FindItems() ([]model.Item) {
	db, err := getDb()
	checkErr(err)

	rows, err := db.Query("select * from item order by id desc")
	checkErr(err)
	items := make([]model.Item, 0)
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		items = append(items, model.Item{id, name})
	}
	return items
}

func CreateItem(item model.Item) {
	db, err := getDb()
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare("INSERT item (`name`) values (?)")
	checkErr(err)
	res, err := stmt.Exec(item.Name)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	log.Printf("create: %T[id=%v, name=%v]", item, id, item.Name)
}

func UpdateItem(item model.Item) int64 {
	db, err := getDb()
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("update item set name=? where id=?")
	checkErr(err)
	res, err := stmt.Exec(item.Name, item.ID)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)

	return num
}

func FindItem(id int) model.Item {
	db, err := getDb()
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("select * from item where id=?")
	checkErr(err)
	rows, err := stmt.Query(id)
	checkErr(err)
	var item model.Item
	if rows.Next()  {
		var id int
		var name string
		rows.Scan(&id, &name)
		item = model.Item{id, name}
	}
	return item
}

func DeleteItem(id int) int64 {
	db, err := getDb()
	checkErr(err)

	defer db.Close()
	stmt, err := db.Prepare(`DELETE FROM item WHERE id=?`)
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	return num
}
