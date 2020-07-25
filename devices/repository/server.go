// go:binary-only-package

package repository

//Server interface for methods implementation
type Server interface {
	GetMethod()
	PostMethod()
}
