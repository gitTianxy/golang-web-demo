package service

import (
	"dao"
	"rao"
	"model"
	"log"
)

type ItemService struct {
	Dao *dao.ItemDao
	Rao *rao.ItemRao
}

func (svc ItemService) IsNull(item model.Item) bool  {
	return item.ID <= 0
}

func (svc ItemService) GetItem(id int64) (item model.Item, err error) {
	item, err = svc.Rao.Get(id)
	if err != nil {
		log.Fatal("itemrao get error.", err)
		return
	}
	if item.ID > 0 {
		return item, nil
	}
	item, err = svc.Dao.FindItem(id)
	if err != nil {
		log.Print("itemdao get error.", err)
		return
	}
	if item.ID > 0 {
		svc.Rao.Set(item)
	}
	return item, nil
}

func (svc ItemService) GetItems() ([]model.Item, error) {
	return svc.Dao.FindItems()
}

func (svc ItemService) PostItem(item model.Item) error {
	err := svc.Dao.CreateItem(&item)
	if err != nil {
		return err
	}
	err = svc.Rao.Set(item)
	if err != nil {
		return err
	}
	return nil
}

func (this ItemService) PutItem(item model.Item) error {
	num, err := this.Dao.UpdateItem(item)
	if err != nil {
		return err
	}
	if num > 0{
		this.Rao.Set(item)
	}
	return nil
}

func (this ItemService) DeleteItem(id int64) error {
	num, err := this.Dao.DeleteItem(id)
	if err != nil {
		return err
	}
	if num > 0 {
		return this.Rao.Del(id)
	}
	return nil
}
