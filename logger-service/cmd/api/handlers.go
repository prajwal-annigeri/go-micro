package main

import (
	"log"
	"log-service/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload
	err := app.readJSON(w, r, &requestPayload)
	if (err != nil) {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	response := jsonResponse{
		Error: false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, response)
}