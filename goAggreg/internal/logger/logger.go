package logger

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Logger interface {
	RaiseFatal(err error)
	RaiseError(err error)
}

type Log struct {
}

func NewLog(ctx context.Context) (*Log, error) {
	return &Log{}, nil
}

func (l *Log) RaiseError(err error) {
	fmt.Fprintf(
		os.Stdin,
		"%s\n\r", fmt.Errorf("raise error: %s\n\r", err.Error()))
}

func (l *Log) RaiseFatal(err error) {
	log.Fatalf("raise fatal: %s\n\r", err.Error())
}
