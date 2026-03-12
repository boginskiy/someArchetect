package repository

import (
	"database/sql"
	"goCnter/internal/db"
)

type Repository interface {
}

type CounterRepo struct {
	DB db.DataBase[*sql.DB]
}

// Бахаем репо слой дальше ...

func UpdateCounter(db *sql.DB, id string, newCount int) error {
	query := "UPDATE counters SET count=$1 WHERE id=$2;"
	_, err := db.Exec(query, newCount, id)
	return err
}

func InsertOrUpdateCounter(db *sql.DB, id string, initialValue int) error {
	query := "INSERT INTO counters (id, count) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET count=EXCLUDED.count;"
	_, err := db.Exec(query, id, initialValue)
	return err
}
