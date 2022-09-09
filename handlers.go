package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func rckStatus(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/rck.html",
	}

	// URL to https://github.com/rcarrata/rck-api microservice
	url := "http://localhost:8080/healthz"
	dat := getApiRequest(url)
	rck_name := "rck"
	log.Printf("Request to -> %s microservice", rck_name)

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

	err = ts.Execute(w, dat)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

// Function to manage the /template?id=XXX
func showTemplate(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Check if the id can't be converted to integer and if it's negative
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// use the Fprintf to write to the variable of the w ResponseWriter
	fmt.Fprintf(w, "Display a specific template with ID %d...", id)
}

func createTemplate(w http.ResponseWriter, r *http.Request) {
	// Raise a Method Not Allowed or 405
	// If the method is POST
	if r.Method == "POST" {

		w.Write([]byte("Create a new template"))

		// Elif method is PUT
	} else if r.Method == "PUT" {

		w.Write([]byte("Update a new template"))

		// For the rest of the Methods
	} else {

		// Use the Header().Set() method to add an 'Allow: POST' header to the response header map
		w.Header().Set("Allow", "POST, PUT")

		// Use instead the http.Error()
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		//Use the http.Error() function to send a 405 status code
		http.Error(w, "Method Not Allowed. Please check Allowed Methods", 405)
		return

	}

}
