package model

import "time"

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Value     int64     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type AgrTypeValue struct {
	Type  string
	Value int64
}
