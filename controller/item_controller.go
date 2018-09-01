package controller

import (
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"golang-web-demo/model"
	"golang-web-demo/util"
	"golang-web-demo/base"
	"golang-web-demo/service"
	"github.com/go-redis/redis"
	"golang-web-demo/dao"
	"golang-web-demo/rao"
	"fmt"
	"errors"
)

type ItemController struct {
	itemSvc *service.ItemService
}

func (this *ItemController) Init(mysqlClient *base.MySQLClient, redisClient *redis.Client)  {
	this.itemSvc = &service.ItemService{}
	itemDao := dao.ItemDao{}
	itemDao.SetMysqlClient(mysqlClient)
	this.itemSvc.Dao = &itemDao
	itemRao := rao.ItemRao{}
	itemRao.SetClient(redisClient)
	this.itemSvc.Rao = &itemRao
}

func (control ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	items, err := control.itemSvc.GetItems()
	if err != nil {
		base.LogRequestErr(r, err)
		return
	}
	ret := model.RespData{200, "ok", items, time.Now().Unix()}
	handler := base.HttpResponseHandler{ w}
	handler.HandleResult(ret)
	base.LogRequestFinish(r, "get items", time.Since(start).Seconds())
}

func (control ItemController) PostItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	handler := base.HttpResponseHandler{w}
	// parse JSON body
	var item model.Item
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		base.LogRequestErr(r, err)
		handler.Fail(err.Error())
		return
	}
	err = json.Unmarshal(body, &item)
	if err != nil {
		msg := fmt.Sprintf("request body err. body:%s, err:%s", body, err.Error())
		base.LogRequestErr(r, errors.New(msg))
		handler.Fail(msg)
		return
	}
	err = control.itemSvc.PostItem(item)
	if err != nil {
		base.LogRequestErr(r, err)
		handler.Fail(err.Error())
		return
	}

	handler.Succ("")
	base.LogRequestFinish(r, "post item", time.Since(start).Seconds())
}

func (control ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	handler := base.HttpResponseHandler{w}

	// parse query parameter
	reqHandler := base.HttpRequestHandler{r}
	id, ok := reqHandler.GetReqParam("id")
	if !ok {
		msg := fmt.Sprintf("invalid param. id:%v", id)
		base.LogRequestErr(r, errors.New(msg))
		handler.Fail(msg)
		return
	}

	item, err := control.itemSvc.GetItem(util.String2Int64(id))
	if err != nil {
		base.LogRequestErr(r, err)
		msg := fmt.Sprintf("get item fail. error:%s", err.Error())
		handler.Fail(msg)
		return
	}

	if control.itemSvc.IsNull(item) {
		handler.NotFound("")
		return
	}

	handler.HandleResult(item)
	base.LogRequestFinish(r, "get item", time.Since(start).Seconds())
}

func (control ItemController) PutItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	handler := base.HttpResponseHandler{w}

	// parse JSON body
	var item model.Item
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &item)
	if err != nil {
		base.LogRequestErr(r, err)
		msg := fmt.Sprintf("invalid request. body:%s", body)
		handler.Fail(msg)
		return
	}
	err = control.itemSvc.PutItem(item)
	if err != nil {
		base.LogRequestErr(r, err)
		handler.Fail(err.Error())
		return
	}

	handler.Succ("")
	base.LogRequestFinish(r, "update item", time.Since(start).Seconds())
}

func (control ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	respHandler := base.HttpResponseHandler{w}

	// parse query parameter
	reqHandler := base.HttpRequestHandler{r}
	id, ok := reqHandler.GetReqParam("id")
	if !ok {
		msg := fmt.Sprintf("invalid param. id=%v", id)
		base.LogRequestErr(r, errors.New(msg))
		respHandler.Fail(msg)
		return
	}

	err := control.itemSvc.DeleteItem(util.String2Int64(id))
	if err != nil {
		base.LogRequestErr(r, err)
		msg := fmt.Sprintf("delete item error. id=%v", id)
		respHandler.Fail(msg)
		return
	}

	respHandler.Succ("")
	base.LogRequestFinish(r, "delete item", time.Since(start).Seconds())
}
