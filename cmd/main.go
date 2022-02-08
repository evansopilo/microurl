package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// struct to store configuration variable read from commandline flag during start of the server.
// port holds server port number
// env holds server environment mode
type config struct {
	port int
	env  string
}

// application struct, wraps application dependencies
type application struct {
	config  config
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {

	var cfg config

	// read comandline flag values and decode it to cfg struct
	flag.IntVar(&cfg.port, "port", 8080, "server port")
	flag.StringVar(&cfg.env, "env", "production", "Environment(development|staging|production)")
	flag.Parse()

	// custom instance of logger that logs application's info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// custom instance of logger that logs application's error
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	// instance of applications struct with application dependencies
	app := &application{
		config:  cfg,
		infoLog: infoLog,
		errLog:  errorLog,
	}

	// server with configured timeout for read, write and idle
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		ErrorLog:     app.errLog,
		Handler:      app.Router(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// start server with configuration information
	infoLog.Println("server listening on port : ", app.config.port)
	if err := srv.ListenAndServe(); err != nil {
		app.errLog.Println(err)
	}
}
