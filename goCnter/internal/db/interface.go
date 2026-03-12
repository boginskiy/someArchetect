package db

type DataBase[T any] interface {
	Close()
	GetDB() T
	Ping() bool
}
