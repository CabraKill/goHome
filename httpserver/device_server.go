// go:binary-only-package

package httpserver

import (
	"log"
	"net/http"
	"strconv"

	"github.com/goHome/devices/repository"
	"github.com/goHome/httpserver/handler"
	"github.com/gookit/color"
)

//HTTPMethod implements a function that receives the network arguments and the database
type HTTPMethod func(http.ResponseWriter, *http.Request, *repository.DevicesDataBase)

//DevicesServer struct with the variables and functions of the Server
type DevicesServer struct {
	Port     int64
	Post     HTTPMethod
	Delete   HTTPMethod
	DataBase *repository.DevicesDataBase
}

func (s DevicesServer) mainHandler(w http.ResponseWriter, r *http.Request) {
	handler.SetupResponse(&w, r)
	handler.OptionsSendOk(&w, r)

	switch method := r.Method; method {
	case "GET":
		color.Yellow.Printf("GET\n")
		response := (*s.DataBase).ToString()
		w.Write([]byte(response))
		return
	case "POST":
		s.Post(w, r, s.DataBase)
		return
	case "DELETE":
		s.Delete(w, r, s.DataBase)
		return
	default:
		response := "Method not implemented ;-;\nSorry, try again."
		w.Write([]byte(response))
	}

}

//Start A function to start the server with its port and HTTP methods
func (s DevicesServer) Start() {
	color.Cyan.Println(".............................................Server on.............................................")
	//http.Handle("/", http.FileServer(http.Dir("./page")))
	http.HandleFunc("/", s.mainHandler)
	if err := http.ListenAndServe(":"+strconv.FormatInt(s.Port, 10), nil); err != nil {
		log.Fatal(err)
	}
}
