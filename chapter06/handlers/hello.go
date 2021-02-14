package handlers

import (
	"html/template"
	"net/http"
)

type user struct {
	Name string
}

// Hello Handles request to /hello
func Hello(writer http.ResponseWriter, request *http.Request) {
	helloTemplate, err := template.ParseFiles("templates/hello.html")

	if err != nil {
		panic(err)
	}

	data := user{
		Name: "James V. Rooney",
	}

	err = helloTemplate.Execute(writer, data)

	if err != nil {
		panic(err)
	}
}
