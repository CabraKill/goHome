package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gookit/color"
	"github.com/gorilla/mux"

	"github.com/goHome/devices/driver"
	"github.com/goHome/devices/model"

	"github.com/goHome/devices/repository"
)

type devicesAPI struct {
	database *repository.DevicesDataBase
}

func (api *devicesAPI) getDevices(w http.ResponseWriter, r *http.Request) {
	color.Red.Printf("get\n")
	response := (*api.database).ToString()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func (api *devicesAPI) addDevice(w http.ResponseWriter, r *http.Request) {
	var device model.DeviceModel

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	color.Blue.Printf("post")
	color.Green.Printf(" - Request: %+v.\n", device)
	(*api.database).AddDevice(device)
	(*api.database).Show()

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := (*api.database).ToString()
	w.Write([]byte(response))
}

func (api *devicesAPI) deleteDevice(w http.ResponseWriter, r *http.Request) {
	var device model.DeviceModel

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	color.Yellow.Printf("delete")
	color.Green.Printf(" - Request: %+v.\n", device)
	(*api.database).RemoveDevice(device)
	(*api.database).Show()

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := (*api.database).ToString()
	w.Write([]byte(response))
}

const (
	port = "8081"
)

func main() {
	database := driver.NewJSON()
	api := devicesAPI{
		&database,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", api.getDevices).Methods(http.MethodGet)
	r.HandleFunc("/", api.addDevice).Methods(http.MethodPost)
	r.HandleFunc("/", api.deleteDevice).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
