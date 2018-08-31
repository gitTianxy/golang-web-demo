package controller

import (
	"net/http"
	"time"
	"encoding/json"
	"log"
	"golang-web-demo/model"
	"golang-web-demo/dao"
	"golang-web-demo/util"
	"golang-web-demo/base"
	"io/ioutil"
)

type ItemController struct {
	ItemDao *dao.ItemDao
}

func (control ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	items := control.ItemDao.FindItems()
	ret := model.RespData{200, "ok", items, time.Now().Unix()}
	handler := base.HttpResponseHandler{ w}
	handler.HandleResult(ret)

	log.Printf(
		"%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		"get items",
		time.Since(start),
	)
}

func (control ItemController) PostItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// parse JSON body
	var item model.Item
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &item)

	control.ItemDao.CreateItem(item)

	handler := base.HttpResponseHandler{w}
	handler.Succ()


	log.Printf(
		"%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		"post item",
		time.Since(start),
	)
}

func (control ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// parse query parameter
	reqHandler := base.HttpRequestHandler{r}
	id := reqHandler.GetParamVal("id")

	item := control.ItemDao.FindItem(util.String2Int(id))
	handler := base.HttpResponseHandler{w}
	handler.HandleResult(item)

	log.Printf(
		"%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		"get item",
		time.Since(start),
	)
}

func (control ItemController) PutItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// parse JSON body
	var item model.Item
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &item)
	control.ItemDao.UpdateItem(item)
	handler := base.HttpResponseHandler{w}
	handler.Succ()

	log.Printf(
		"%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		"update item",
		time.Since(start),
	)
}

func (control ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// parse query parameter
	reqHandler := base.HttpRequestHandler{r}
	id := reqHandler.GetParamVal("id")

	control.ItemDao.DeleteItem(util.String2Int(id))

	handler := base.HttpResponseHandler{w}
	handler.Succ()

	log.Printf(
		"%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		"delete item",
		time.Since(start),
	)
}
