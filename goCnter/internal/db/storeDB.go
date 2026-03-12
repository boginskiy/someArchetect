package db

import (
	"context"
	"database/sql"
	"log"
)

var conn = "postgres://username:userpassword@localhost:5432/testdb?sslmode=disable"

type StoreDB struct {
	DB *sql.DB
}

func NewStoreDB(ctx context.Context) *StoreDB {
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	return &StoreDB{
		DB: db,
	}
}

func (s *StoreDB) GetDB() **sql.DB {
	return &s.DB
}

func (s *StoreDB) Ping() bool {
	err := s.DB.Ping()
	if err != nil {
		return false
	}
	return true
}

func (s *StoreDB) Close() {
	s.DB.Close()
}
