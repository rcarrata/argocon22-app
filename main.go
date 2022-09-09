package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// All incoming HTTP requests are server in their own goroutine
	addr := flag.String("addr", ":8080", "HTTP network address")

	// Parse the command line and assigns it to the addr variable
	flag.Parse()

	// Use log.New() to create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// writing error messages using the stderr as the output
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new servemux
	mux := http.NewServeMux()

	// Register a / as the index
	mux.HandleFunc("/", index)

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

	// Starting a new web server on 8080 port and log the errors if any
	// mux in ListenAndServe passes the response itself on to a second handler

	err := http.ListenAndServe(*addr, mux)

	errorLog.Fatal(err)
}
