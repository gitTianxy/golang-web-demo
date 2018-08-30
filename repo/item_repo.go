package repo

import (
	"golang-web-demo/model"
	"fmt"
)

func getItem(idx int) (model.Item) {
	return model.Item{idx, fmt.Sprintf("item_%v", idx)}
}

func FindItems() (items []model.Item) {
	item1 := getItem(1)
	item2 := getItem(2)
	item3 := getItem(3)
	return []model.Item{item1, item2, item3}
}

// TODO
func CreateItem(item model.Item) {

}

// TODO
func UpdateItem(item model.Item) {
}

// TODO
func FindItem(id int) model.Item  {
	item := getItem(id)
	return item
}

// TODO
func DeleteItem(id int)  {

}