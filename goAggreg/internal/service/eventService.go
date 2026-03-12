package service

import (
	"goAggreg/internal/model"
	"sync"
	"sync/atomic"
)

type EventServicer interface {
	Send(event *model.Event) error
}

type EventService struct {
	store map[string]int
	cnt   int64
	rMx   sync.RWMutex
}

func NewEventService() *EventService {
	return &EventService{}
}

func (s *EventService) Send(event *model.Event) error {
	if len(s.store) == 10 {
		// Делаем список Event и отправляем его в репозиторий
	}

	s.rMx.Lock()
	s.store[event.Name] += event.Value
	s.rMx.Unlock()
	atomic.AddInt64(&s.cnt, 1)
	return nil
}
