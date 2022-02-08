package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var health = struct {
		Status  string `json:"status,omitempty"`
		Version string `json:"version,omitempty"`
	}{
		Status:  "Healthy",
		Version: "1.0.0",
	}

	if err := json.NewEncoder(w).Encode(&health); err != nil {
		app.errLog.Println(err)
		return
	}

}
