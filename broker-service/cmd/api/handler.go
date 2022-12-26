package main

import (
	"fmt"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the Broker",
	}

	fmt.Println("Method: ", r.Method)

	app.writeJSON(w, http.StatusOK, payload)

}
