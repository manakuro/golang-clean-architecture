package repository

type DBRepository interface {
	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}
