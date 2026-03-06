package logg

import "log"

type Logger interface {
	RaiseFatal(msg string)
}

type Logg struct {
}

func NewLogg() *Logg {
	return &Logg{}
}

func (l *Logg) RaiseFatal(msg string) {
	log.Fatalf("%s", msg)
}
