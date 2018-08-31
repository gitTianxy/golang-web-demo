package dao

import (
	"golang-web-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"golang-web-demo/base"
)

type ItemDao struct {
	client base.MySQLClient
}

func (dao *ItemDao) SetMysqlClient(client base.MySQLClient)  {
	dao.client = client
}

func (dao ItemDao) FindItems() ([]model.Item) {
	pool := dao.client.GetPool()
	rows, err := pool.Query("select * from item order by id desc")
	base.CheckErr(err)
	items := make([]model.Item, 0)
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		items = append(items, model.Item{id, name})
	}
	return items
}

func (dao ItemDao) CreateItem(item model.Item) {
	pool := dao.client.GetPool()
	res, err := pool.Exec("INSERT item (`name`) values (?)", item.Name)
	base.CheckErr(err)
	id, err := res.LastInsertId()
	base.CheckErr(err)
	log.Printf("create: %T[id=%v, name=%v]", item, id, item.Name)
}

func (dao ItemDao) UpdateItem(item model.Item) int64 {
	pool := dao.client.GetPool()
	stmt, err := pool.Prepare("update item set name=? where id=?")
	base.CheckErr(err)
	res, err := stmt.Exec(item.Name, item.ID)
	base.CheckErr(err)
	num, err := res.RowsAffected()
	base.CheckErr(err)

	return num
}

func (dao ItemDao) FindItem(id int) model.Item {
	pool := dao.client.GetPool()
	stmt, err := pool.Prepare("select * from item where id=?")
	base.CheckErr(err)
	row := stmt.QueryRow(id)
	var dbId int
	var dbName string
	err = row.Scan(&dbId, &dbName)
	base.CheckErr(err)
	return model.Item{dbId, dbName}
}

func (dao ItemDao) DeleteItem(id int) int64 {
	pool := dao.client.GetPool()
	res, err := pool.Exec("DELETE FROM item WHERE id=?", id)
	base.CheckErr(err)
	num, err := res.RowsAffected()
	base.CheckErr(err)
	return num
}
