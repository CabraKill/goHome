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

//PostFunction hgfhgf
type PostFunction func(http.ResponseWriter, *http.Request, *repository.DevicesDataBase)

//HTTPServer asdawdawd
type HTTPServer struct {
	Port     int64
	Post     PostFunction
	DataBase *repository.DevicesDataBase
}

//PostHandler kjhkjh
func (s HTTPServer) PostHandler(w http.ResponseWriter, r *http.Request) {
	//handler.SetupResponse(&w, r)
	//handler.OptionsSendOk(&w, r)

	s.Post(w, r, s.DataBase)
}

func (s HTTPServer) mainHandler(w http.ResponseWriter, r *http.Request) {
	handler.SetupResponse(&w, r)
	handler.OptionsSendOk(&w, r)

	switch method := r.Method; method {
	case "GET":
		//response := databaseDriver.DataBase.ToString()
		response := "ola"
		w.Write([]byte(response))
		return
	case "POST":
		s.PostHandler(w, r)
		return
	default:
		response := "Method not implemented ;-;\nSorry, try again."
		w.Write([]byte(response))
	}

}

//Start A function to start the server with its port and HTTP methods
func (s HTTPServer) Start() {
	//.NewDataBaseJson()

	//fmt.Printf("................Server on................\n")
	color.Cyan.Println("................................Server on................................")
	//http.Handle("/", http.FileServer(http.Dir("./page")))
	http.HandleFunc("/", s.mainHandler)

	http.HandleFunc("/post", s.PostHandler)
	if err := http.ListenAndServe(":"+strconv.FormatInt(s.Port, 10), nil); err != nil {
		log.Fatal(err)
	}
}
