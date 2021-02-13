package handlers

import (
	"fmt"
	"net/http"
)

// Contact Handles requests to /contact
func Contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	fmt.Fprint(writer, "James - To get in touch email: <a href=\"mailto:jamesvrooney@hotmail.com\">jamesvrooney@hotmail.com</a>")
}
