package main

import (
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"golang-web-demo/base"
	"golang-web-demo/controller"
	"net/http"
)

var mysqlClientShared *base.MySQLClient
var redisClientShared *redis.Client
var router *httprouter.Router

/**
 * init resources
 */
func init()  {
	// init `MySQLClient`
	mysqlClientShared = base.GetMysqlClient()
	// init `RedisClient`
	redisClientShared = base.GetRedisClient()
	// init httprouter
	router = httprouter.New()
}

func main() {
	// define controllers
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
