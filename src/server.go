package main

import (
	"github.com/julienschmidt/httprouter"
	"base"
	"controller"
	"net/http"
)

func main() {
	// init `MySQLClient`
	mysqlClientShared := base.GetMysqlClient()
	// init `RedisClient`
	redisClientShared := base.GetRedisClient()

	// init controllers
	router := httprouter.New()
	router.GET("/", controller.Index)

	itemController := controller.ItemController{}
	itemController.Init(mysqlClientShared, redisClientShared)
	router.GET("/api/items", itemController.GetItems)
	router.GET("/api/item/:id", itemController.GetItem)
	router.POST("/api/item", itemController.PostItem)
	router.PUT("/api/item/:id", itemController.PutItem)
	router.DELETE("/api/item/:id", itemController.DeleteItem)

	// listen port
	http.ListenAndServe("localhost:8080", router)
}
