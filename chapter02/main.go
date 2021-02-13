package main

import (
	"fmt"
	"net/http"

	"githb.com/jamesvrooney/go_web_dev_with_jc/chapter02/handlers"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello James, you big eejit! Up Leitrim</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch email: <a href=\"mailto:jamesvrooney@hotmail.com\">jamesvrooney@hotmail.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email if you keep getting an invalid page.</p>")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Hello)
	router.HandleFunc("/contact", handlers.Contact)

	http.ListenAndServe(":3000", router)
}
