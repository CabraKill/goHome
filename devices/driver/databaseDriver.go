// go:binary-only-package

package driver

import (
	repository "github.com/goHome/devices/repository"
)

var (
	//DataBase exports the DataBase
	DataBase repository.DevicesDataBase
)

//NewDataBaseJSON creates a new instance
func NewDataBaseJSON() {
	DataBase = &repository.JsonRepo{}
}
