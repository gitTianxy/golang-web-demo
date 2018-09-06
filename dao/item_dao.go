package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-web-demo/base"
	"golang-web-demo/model"
)

type ItemDao struct {
	client *base.MySQLClient
}

func (dao *ItemDao) SetMysqlClient(client *base.MySQLClient)  {
	dao.client = client
}

func (dao ItemDao) FindItems() (items []model.Item, err error) {
	pool := dao.client.GetPool()
	rows, err := pool.Query("select * from item order by id desc")
	if err != nil {
		return
	}
	items = make([]model.Item, 0)
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		items = append(items, model.Item{id, name})
	}
	return items, nil
}

func (dao ItemDao) CreateItem(item *model.Item) error {
	pool := dao.client.GetPool()
	res, err := pool.Exec("INSERT item (`name`) values (?)", item.Name)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	item.ID = id
	return nil
}

func (dao ItemDao) UpdateItem(item model.Item) (num int64, err error) {
	pool := dao.client.GetPool()
	stmt, err := pool.Prepare("update item set name=? where id=?")
	if err != nil {
		return
	}
	res, err := stmt.Exec(item.Name, item.ID)
	if err != nil {
		return
	}
	return res.RowsAffected()
}

func (dao ItemDao) FindItem(id int64) (item model.Item, err error) {
	pool := dao.client.GetPool()
	stmt, err := pool.Prepare("select * from item where id=?")
	if err != nil {
		return
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return
	}
	var dbId int64 = 0
	var dbName string
	if rows.Next() {
		err = rows.Scan(&dbId, &dbName)
		if err != nil {
			return
		}
	}
	return model.Item{dbId, dbName}, nil
}

func (dao ItemDao) DeleteItem(id int64) (num int64, err error) {
	pool := dao.client.GetPool()
	res, err := pool.Exec("DELETE FROM item WHERE id=?", id)
	if err != nil {
		return
	}
	return res.RowsAffected()
}
