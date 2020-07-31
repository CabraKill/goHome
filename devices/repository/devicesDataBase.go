// go:binary-only-package

package repository

import (
	model "github.com/goHome/devices/model"
)

//DevicesDataBase Interface for Devices' Repository
type DevicesDataBase interface {
	GetDevices() []model.DeviceModel
	AddDevice(device model.DeviceModel)
	RemoveDevice(device model.DeviceModel)
	ToString() string
	Show()
}
