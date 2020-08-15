package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Application-wide dependancies struct
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	// Check comand-line for port number. The default value is 4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse() // Use Parse() to read the provided value and assign to addr

	// Custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// Initialize a new http.Server struct where we can add Addr and Handler
	// fileds that the server uses the same network address and routes as
	// before.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// The value returned from flag.String() function is a pointer to the
	// function value, not the value itself. So we need to dereference the
	// pointer.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
