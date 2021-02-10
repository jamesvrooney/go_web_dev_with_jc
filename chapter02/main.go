package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello James, you big eejit! Up Leitrim</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch email: <a href=\"mailto:jamesvrooney@hotmail.com\">jamesvrooney@hotmail.com</a>")
	}
}

func main() {
	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":3000", nil)
}
