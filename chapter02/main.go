package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Another request")
	fmt.Fprint(w, "<h1>Hello James, you big eejit! Up Leitrim</h1>")
}

func main() {
	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":3000", nil)
}
