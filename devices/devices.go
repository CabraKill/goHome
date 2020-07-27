package main

import (
	"encoding/json"
	"net/http"

	"github.com/gookit/color"

	"github.com/goHome/devices/driver"
	"github.com/goHome/devices/model"
	
	"github.com/goHome/devices/repository"
	"github.com/goHome/httpserver"
)

const (
	port = 8081
)

func post(w http.ResponseWriter, r *http.Request, db *repository.DevicesDataBase) {
	var device model.DeviceModel

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//log.Printf("Request: %+v.\n", device)
	color.Green.Printf("Request: %+v.\n", device)
	(*db).AddDevice(device)
	(*db).Show()

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := (*db).ToString()
	w.Write([]byte(response))
}

func main() {
	database := driver.NewJSON()
	server := httpserver.HTTPServer{
		Port: port,
		Post: post,
		DataBase: &database,
	}

	server.Start()
}
