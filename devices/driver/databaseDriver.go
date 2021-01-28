package driver

import (
	repository "github.com/goHome/devices/repository"
)

var (
	//DataBase exports the DataBase
	DataBase repository.DevicesDataBase
)

//NewJSON returns a new instance
func NewJSON() repository.DevicesDataBase {
	return &repository.JSONRepo{}
}
