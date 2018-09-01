package main


import (
	"net/http"
	"golang-web-demo/controller"
	"github.com/gorilla/mux"
	"golang-web-demo/base"
)

func main() {
	// init `MySQLClient`
	mysqlClientShared := base.GetMysqlClient()
	// init `RedisClient`
	redisClientShared := base.GetRedisClient()

	// init controllers
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index).Methods("GET")

	itemController := controller.ItemController{}
	itemController.Init(mysqlClientShared, redisClientShared)
	router.HandleFunc("/api/items", itemController.GetItems).Methods("GET")
	router.HandleFunc("/api/item", itemController.GetItem).Methods("GET")
	router.HandleFunc("/api/item", itemController.PostItem).Methods("POST")
	router.HandleFunc("/api/item", itemController.PutItem).Methods("PUT")
	router.HandleFunc("/api/item", itemController.DeleteItem).Methods("DELETE")

	// listen port
	http.ListenAndServe("localhost:8085", router)
}
