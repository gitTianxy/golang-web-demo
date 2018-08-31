package main


import (
	"net/http"
	"golang-web-demo/controller"
	"github.com/gorilla/mux"
	"golang-web-demo/base"
	"golang-web-demo/dao"
)

func main() {
	// init `MySQLClient`
	mysqlClientShared := base.GetMysqlClient()
	// init `ItemController`
	itemDao := dao.ItemDao{}
	itemDao.SetMysqlClient(mysqlClientShared)

	// init web service
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index).Methods("GET")
	itemController := controller.ItemController{&itemDao}
	router.HandleFunc("/api/items", itemController.GetItems).Methods("GET")
	router.HandleFunc("/api/item", itemController.GetItem).Methods("GET")
	router.HandleFunc("/api/item", itemController.PostItem).Methods("POST")
	router.HandleFunc("/api/item", itemController.PutItem).Methods("PUT")
	router.HandleFunc("/api/item", itemController.DeleteItem).Methods("DELETE")

	// listen port
	http.ListenAndServe("localhost:8085", router)
}
