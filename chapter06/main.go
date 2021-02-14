package main

import (
	"net/http"

	"githb.com/jamesvrooney/go_web_dev_with_jc/chapter06/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", handlers.Hello)

	http.ListenAndServe(":8877", router)
}
