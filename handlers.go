package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Health struct {
	Status string
}

func index(w http.ResponseWriter, r *http.Request) {
	// Raise an 404 if the url not matches
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Hello from the Webserver"))

	// define an slice (lists in python)
	files := []string{
		"./ui/html/index.html",
	}

	// Implementing variadic functions for the files
	// to avoid to add manually
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Return an error for the stdout
		log.Println(err.Error())
		// Return an error to the ResponseWriter function
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
