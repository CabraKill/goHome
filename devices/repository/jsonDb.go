// go:binary-only-package

package repository

import (
	"encoding/json"
	"fmt"

	model "github.com/goHome/devices/model"
)

type JsonRepo struct {
	devices []model.DeviceModel
}

func (j *JsonRepo) GetDevices() []model.DeviceModel {
	return j.devices
	/*return []model.DeviceModel{
		model.DeviceModel{
			Name: "pc",
		},
	}*/
}

func (j *JsonRepo) AddDevice(device model.DeviceModel) {
	j.devices = append(j.devices, device)

}

func (j *JsonRepo) RemoveDevice(device model.DeviceModel) {
	for i, v := range j.devices {
		if v.Ip == device.Ip && v.Name == device.Name {
			j.devices = append(j.devices[:i], j.devices[i+1:]...)
			break
		}
	}
}

func (j *JsonRepo) Show() {
	for i := 0; i < len(j.devices); i++ {
		fmt.Printf("%+v\n", j.devices[i])
	}

}

func (j *JsonRepo) ToString() string {
	//text := ""
	for i := 0; i < len(j.devices); i++ {
		//text += fmt.Sprintf("%+v,",j.devices[i])

	}
	//return text
	text, _ := json.Marshal(j.devices)
	return string(text)
}
