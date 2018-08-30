package database

type SqlHandler interface {
	Create(*interface{}) (Result, error)
	Find(*interface{}) (Result, error)
	Update(*interface{}) (Result, error)
	Where(*interface{}) (Result, error)
	Destroy(*interface{}) error
}

type Result *interface{}
