package handlers

import (
	"html/template"
	"net/http"
)

type user struct {
	Name string
}

// Hello Handles requests to /hello
func Hello(writer http.ResponseWriter, request *http.Request) {
	tmp, err := template.ParseFiles("templates/hello.html")

	if err != nil {
		panic(err)
	}

	james := user{
		Name: "James",
	}

	err = tmp.Execute(writer, james)

	if err != nil {
		panic(err)
	}
}
