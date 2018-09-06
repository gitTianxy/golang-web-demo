package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"golang-web-demo/base"
	"golang-web-demo/dao"
	"golang-web-demo/model"
	"golang-web-demo/rao"
	"golang-web-demo/service"
	"golang-web-demo/util"
	"io/ioutil"
	"net/http"
	"time"
)

type ItemController struct {
	itemSvc *service.ItemService
	logHandler *base.RestfulLogHandler
}

func (this *ItemController) Init(mysqlClient *base.MySQLClient, redisClient *redis.Client)  {
	this.itemSvc = &service.ItemService{}
	itemDao := dao.ItemDao{}
	itemDao.SetMysqlClient(mysqlClient)
	this.itemSvc.Dao = &itemDao
	itemRao := rao.ItemRao{}
	itemRao.SetClient(redisClient)
	this.itemSvc.Rao = &itemRao

	this.logHandler = base.GetRestfulLogger()
}

func (this ItemController) GetItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()
	handler := base.HttpResponseHandler{ w}
	items, err := this.itemSvc.GetItems()
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		handler.Fail("get items fail.")
		return
	}
	ret := model.RespData{200, "ok", items, time.Now().Unix()}

	handler.HandleResult(ret)
	this.logHandler.LogRequestFinish(r, "get items", time.Since(start))
}

func (this ItemController) PostItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	handler := base.HttpResponseHandler{w}
	// parse JSON body
	var item model.Item
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		handler.Fail("read request body fail.")
		return
	}
	err = json.Unmarshal(body, &item)
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		handler.Fail("parse request body fail")
		return
	}
	err = this.itemSvc.PostItem(item)
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		handler.Fail("post item fail")
		return
	}

	handler.Succ("post item success")
	this.logHandler.LogRequestFinish(r, "post item", time.Since(start))
}

func (this ItemController) GetItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	handler := base.HttpResponseHandler{w}

	// parse query parameter
	id := ps.ByName("id")
	if len(id) == 0 {
		msg := fmt.Sprintf("invalid param. id:%v", id)
		this.logHandler.LogRequestErr(r, errors.New(msg))
		handler.Fail(msg)
		return
	}

	item, err := this.itemSvc.GetItem(util.String2Int64(id))
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		msg := fmt.Sprintf("get item fail. error:%s", err.Error())
		handler.Fail(msg)
		return
	}

	if this.itemSvc.IsNull(item) {
		msg := fmt.Sprintf("item not found. id:%v", id)
		handler.NotFound(msg)
		return
	}

	handler.HandleResult(item)
	this.logHandler.LogRequestFinish(r, "get item", time.Since(start))
}

func (this ItemController) PutItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	handler := base.HttpResponseHandler{w}

	// parse JSON body
	var item model.Item
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &item)
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		msg := fmt.Sprintf("invalid request. body:%s", body)
		handler.Fail(msg)
		return
	}
	err = this.itemSvc.PutItem(item)
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		handler.Fail(err.Error())
		return
	}

	handler.Succ("")
	this.logHandler.LogRequestFinish(r, "update item", time.Since(start))
}

func (this ItemController) DeleteItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	respHandler := base.HttpResponseHandler{w}

	// parse query parameter
	id := ps.ByName("id")
	if len(id) == 0 {
		msg := fmt.Sprintf("invalid param. id=%v", id)
		this.logHandler.LogRequestErr(r, errors.New(msg))
		respHandler.Fail(msg)
		return
	}

	err := this.itemSvc.DeleteItem(util.String2Int64(id))
	if err != nil {
		this.logHandler.LogRequestErr(r, err)
		msg := fmt.Sprintf("delete item error. id=%v", id)
		respHandler.Fail(msg)
		return
	}

	respHandler.Succ("")
	this.logHandler.LogRequestFinish(r, "delete item", time.Since(start))
}
