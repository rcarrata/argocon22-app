package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// All incoming HTTP requests are server in their own goroutine

	// Define a new command-line flag with the name 'addr', the default value 3000, and the
	// short text defining the text
	addr := flag.String("addr", ":3000", "HTTP network address")

	// Parse the command line and assigns it to the addr variable
	flag.Parse()

	// Use log.New() to create a logger for writing information messages.
	// the destination to write the logs to (os.Stdout)
	// prefix for message (INFO followed by a tab), and flags to indicate
	// additional information to include (local date and time).
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// writing error messages using the stderr as the output
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new servemux
	mux := http.NewServeMux()

	// Register a / as the index
	mux.HandleFunc("/", index)
	// this is equal to
	// http.HandlerFunc works adding a ServerHTTP() to the index function
	// mux.HandleFunc("/", http.HandleFunc(index))

	// Register a /template for show the existent template
	mux.HandleFunc("/rck", rckStatus)

	// Register a /template for show the existent template
	mux.HandleFunc("/template", showTemplate)

	// Register a /template/create for create a new template
	mux.HandleFunc("/template/create", createTemplate)

	// path /static/ subtree path pattern (wildcard at the end)
	// Serves files from ./ui/static
	fileServer := http.FileServer(http.Dir("ui/static"))

	// Register an tge file server as the handler that maps all URL
	// paths that start /static/. Then it will strip /static and load the fileServer
	// to the / index
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Log to the standard output the message
	// log.Println("Starting Server on", *addr)
	infoLog.Printf("Starting Server on %s", *addr)

	// Starting a new web server on 3000 port and log the errors if any
	// mux in ListenAndServe passes the response itself on to a second handler

	// The value returned from flag.String is a pointer to the function value
	// not the value itself. We need to deference the pointer (*addr) before to using it.
	// Dereferencing a pointer gives us access to the value the pointer points to.
	err := http.ListenAndServe(*addr, mux)

	// log.Fatal(err)
	errorLog.Fatal(err)
}
