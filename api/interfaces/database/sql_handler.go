package database

type SqlHandler interface {
	Create(interface{}) error
}
