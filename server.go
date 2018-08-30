package main


import (
	"net/http"
	"golang-web-demo/controller"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index).Methods("GET")
	router.HandleFunc("/api/items", controller.GetItems).Methods("GET")
	router.HandleFunc("/api/item", controller.GetItem).Methods("GET")
	router.HandleFunc("/api/item", controller.PostItem).Methods("POST")
	router.HandleFunc("/api/item", controller.PutItem).Methods("PUT")
	router.HandleFunc("/api/item", controller.DeleteItem).Methods("DELETE")
	http.ListenAndServe("localhost:8085", router)
}
