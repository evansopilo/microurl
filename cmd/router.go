package main

import "github.com/gorilla/mux"

func (app *application) Router() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", app.HealthCheck)

	return r
}
