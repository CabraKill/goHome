// go:binary-only-package

package model

//DeviceModel with the info about it
//such as name and actions.
type DeviceModel struct {
	Name    string        `json:"name"`
	Ip      string        `json:"ip"`
	Type    int           `json:"type"`
	Actions []ActionModel `json:"actions"`
}
