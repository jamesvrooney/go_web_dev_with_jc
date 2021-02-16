package handlers

import (
	"net/http"

	"githb.com/jamesvrooney/go_web_dev_with_jc/chapter06/views"
)

var homeView = views.NewView("bootstrap", "views/hello.html")

type user struct {
	Name string
}

// Hello Handles request to /hello
func Hello(writer http.ResponseWriter, request *http.Request) {

	data := user{
		Name: "James V. Rooney",
	}

	// err := homeView.Template.Execute(writer, data)
	err := homeView.Template.ExecuteTemplate(writer, homeView.Layout, data)

	if err != nil {
		panic(err)
	}
}
