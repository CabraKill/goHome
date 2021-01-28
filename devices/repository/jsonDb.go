package repository

import (
	"encoding/json"
	"fmt"

	model "github.com/goHome/devices/model"
)

//JSONRepo contains the variables and functions to sustain the database.
type JSONRepo struct {
	devices []model.DeviceModel
}

//GetDevices returns the devices list from the json list.
func (j *JSONRepo) GetDevices() []model.DeviceModel {
	return j.devices
}

//AddDevice adds a given device into the json list.
//If the ip is already known the properties are overriden
func (j *JSONRepo) AddDevice(device model.DeviceModel) {
	for i, v := range j.devices {
		if v.Ip == device.Ip && v.Name == device.Name {
			j.devices[i] = device
			return
		}
	}
	j.devices = append(j.devices, device)

}

//RemoveDevice removes a given device from the json list.
func (j *JSONRepo) RemoveDevice(device model.DeviceModel) {
	for i, v := range j.devices {
		if v.Ip == device.Ip && v.Name == device.Name {
			j.devices = append(j.devices[:i], j.devices[i+1:]...)
			break
		}
	}
}

//Show prints all devices into the json list.
func (j *JSONRepo) Show() {
	for i := 0; i < len(j.devices); i++ {
		fmt.Printf("%+v\n", j.devices[i])
	}

}

//ToString returns the json as string after Marshal.
func (j *JSONRepo) ToString() string {
	//text := ""
	for i := 0; i < len(j.devices); i++ {
		//text += fmt.Sprintf("%+v,",j.devices[i])

	}
	//return text
	text, _ := json.Marshal(j.devices)
	return string(text)
}
