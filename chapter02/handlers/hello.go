package handlers

import (
	"fmt"
	"net/http"
)

// Hello Handles requests to /hello
func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	fmt.Fprint(writer, "<h1>Hello James2, you big eejit! Up Leitrim</h1>")
}
