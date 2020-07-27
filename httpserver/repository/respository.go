package repository

import (
	"net/http"

	"github.com/goHome/devices/repository"
)

//ServerFunctions awdawd
type ServerFunctions interface {
	PostHandler(w http.ResponseWriter, r *http.Request, db *repository.DevicesDataBase)
}
