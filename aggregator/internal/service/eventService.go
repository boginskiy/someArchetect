package service

import (
	"aggregator/cmd/config"
	"aggregator/internal/converter"
	"aggregator/internal/logg"
	"aggregator/internal/model"
	"aggregator/internal/repository"
	"aggregator/pkg/maps"
	"sync/atomic"

	"context"
	"fmt"
)

type EventServi struct {
	Repo      repository.EventRepository
	Converter converter.EventConverter

	store  maps.CountMaper
	limit  *int64
	cfg    config.Config
	logger logg.Logger
}

func NewEventServi(
	ctx context.Context,
	config config.Config,
	logger logg.Logger,
	repo repository.EventRepository,
	eventCh <-chan *model.Event,
	converter converter.EventConverter,
) *EventServi {

	tmpEvent := &EventServi{
		Repo:      repo,
		Converter: converter,
		cfg:       config,
		logger:    logger,
		limit:     new(int64),
		store:     maps.NewCounterMap(),
	}

	go tmpEvent.Reciver(ctx, eventCh)

	return tmpEvent
}

func (s *EventServi) Reciver(ctx context.Context, eventCh <-chan *model.Event) {
	for {
		select {
		case <-ctx.Done():
			s.pprint()
			return

		case event := <-eventCh:
			s.Procc(event)
		}
	}
}

func (s *EventServi) Procc(event *model.Event) {
	// Сохраняем определенное количество данных и далее отправляем в слой Репо.

	if atomic.LoadInt64(s.limit) < 3 {
		s.store.Put(event.Type, int64(event.Value))
		atomic.AddInt64(s.limit, 1)

	} else {
		snapShot := s.store.SnapShot()
		listOfAgr := s.Converter.ConvertMapToAgrTypeValue(snapShot)
		s.Repo.Update(listOfAgr)
		s.limit = new(int64)
	}
}

func (s *EventServi) pprint() {
	fmt.Printf("Result: %v\n\r", s.store)
}

// Kafka далее
// Проработка архитектуры Козырев
