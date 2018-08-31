package service

import "golang-web-demo/model"

type ItemService interface {
	GetItem(id int) model.Item
	GetItems() []model.Item
	PostItem(item model.Item) int
	PutItem(item model.Item)
	DeleteItem(id int) int
}