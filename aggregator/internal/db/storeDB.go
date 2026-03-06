package db

import (
	"aggregator/cmd/config"
	"aggregator/internal/logg"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type StoreDB struct {
	cfg    config.Config
	logger logg.Logger
	DB     *sql.DB
}

func NewDB(ctx context.Context, config config.Config, logger logg.Logger) *StoreDB {
	db, err := sql.Open("postgres", config.GetConnToDB())

	if err != nil {
		logger.RaiseFatal(err.Error())
		return nil
	}

	return &StoreDB{
		cfg:    config,
		logger: logger,
		DB:     db,
	}
}

func (s *StoreDB) Close() {
	s.DB.Close()
}

func (s *StoreDB) GetDB() *sql.DB {
	return s.DB
}

func (s *StoreDB) Ping() bool {
	err := s.DB.Ping()
	if err != nil {
		return false
	}
	return true
}
