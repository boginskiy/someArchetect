package db

import "database/sql"

type DataBase interface {
	Close()
	GetDB() *sql.DB
	Ping() bool
}
