package handlers

import (
	"net/http"

	"githb.com/jamesvrooney/go_web_dev_with_jc/chapter06/views"
)

var contactView = views.NewView("bootstrap", "views/contact.html")

type doctor struct {
	Name string
	Age  int
}

// Contact Handles request to /contact
func Contact(writer http.ResponseWriter, request *http.Request) {

	data := doctor{
		Name: "Dr. James",
		Age:  46,
	}

	// err := contactView.Template.Execute(writer, data)
	err := contactView.Template.ExecuteTemplate(writer, contactView.Layout, data)

	if err != nil {
		panic(err)
	}
}
