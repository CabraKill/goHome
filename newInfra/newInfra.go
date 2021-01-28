package main

import (
	"log"
	"net/http"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
	// "github.com/goHome/devices/driver"
)

type httpConnectionExample struct {
	dataBase string
}

func (z httpConnectionExample) ListDevices(w http.ResponseWriter, r *http.Request) {
	color.Red.Println("get")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "` + z.dataBase + `"}`))
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
	con := httpConnectionExample{
		"testando",
	}
	r := mux.NewRouter()
	r.HandleFunc("/", con.ListDevices).Methods(http.MethodGet)
	r.HandleFunc("/", addDevice).Methods(http.MethodPost)
	r.HandleFunc("/", updateDevice).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
