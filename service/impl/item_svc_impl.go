package impl

import (
	"golang-web-demo/dao"
	"golang-web-demo/rao"
	"golang-web-demo/model"
	"encoding/json"
	"log"
)

func GetItemService() ItemServiceImpl {
	svc := ItemServiceImpl{

	}
	return svc
}

type ItemServiceImpl struct {
	dao dao.ItemDao
	rao rao.ItemRao
}

func (svc ItemServiceImpl) GetItem(id int) model.Item {
	var item model.Item
	val, ok := svc.rao.Get("key")
	if !ok {
		log.Fatal("itemrao get error")
		panic("itemrao get error")
	}
	if len(val) > 0 {
		json.Unmarshal([]byte(val), &item)
	} else {
		item = svc.dao.FindItem(id)
	}

	return item
}

//TODO
func (svc ItemServiceImpl) GetItems() []model.Item {
	return []model.Item{}
}
//TODO
func (svc ItemServiceImpl) PostItem(item model.Item) int {
	return 0
}
//TODO
func (svc ItemServiceImpl) PutItem(item model.Item) {

}
//TODO
func (svc ItemServiceImpl) DeleteItem(id int) int {
	return 0
}
