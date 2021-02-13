package handlers

import (
	"fmt"
	"net/http"
)

// Hello Handles requests to /hello
func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Chapter 04</h1>")
}
