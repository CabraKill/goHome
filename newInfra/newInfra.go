package main

import (
	"log"
	"net/http"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
	// "github.com/goHome/devices/driver"
)

func listDevices(w http.ResponseWriter, r *http.Request) {
	color.Red.Println("get")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func addDevice(w http.ResponseWriter, r *http.Request) {
	color.Cyan.Println("post")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "post called"}`))
}

func updateDevice(w http.ResponseWriter, r *http.Request) {
	color.Cyan.Println("put")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "put called"}`))
}

const (
	port = "8080"
)

func main() {
	// database := driver.NewJson()
	r := mux.NewRouter()
	r.HandleFunc("/", listDevices).Methods(http.MethodGet)
	r.HandleFunc("/", addDevice).Methods(http.MethodPost)
	r.HandleFunc("/", updateDevice).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
