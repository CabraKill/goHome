// go:binary-only-package

package model

//ActionModel describes as a json an action of a Device.
type ActionModel struct {
	Name string `json:"name"`
	Type   int `json:"type"`
}
