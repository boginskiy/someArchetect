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

const CntWorkers = 3

type AggregatorServi struct {
	Repo      repository.EventRepository
	Converter converter.EventConverter

	store  maps.CountMaper
	limit  *int64
	cfg    config.Config
	logger logg.Logger
}

func NewAggregatorServi(
	ctx context.Context,
	config config.Config,
	logger logg.Logger,
	repo repository.EventRepository,
	eventCh <-chan *model.Event,
	converter converter.EventConverter,
) *AggregatorServi {

	tmpEvent := &AggregatorServi{
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

func (s *AggregatorServi) Reciver(ctx context.Context, eventCh <-chan *model.Event) {
	s.workers(CntWorkers, eventCh)
}

func (s *AggregatorServi) workers(cnt int, eventCh <-chan *model.Event) {
	done := make(chan bool)

	for i := 0; i < cnt; i++ {
		go func(workerID int) {
			for event := range eventCh {
				s.Procc(event)
			}
			done <- true
		}(i)
	}

	// Stop
	for i := 0; i < cnt; i++ {
		<-done
	}
}

func (s *AggregatorServi) Procc(event *model.Event) {
	// Сохраняем определенное количество данных и далее отправляем в слой Репо.

	if atomic.LoadInt64(s.limit) < 3 {
		s.store.Put(event.Type, int64(event.Value))
		atomic.AddInt64(s.limit, 1)

	} else {
		snapShot := s.store.SnapShot()
		listOfAgr := s.Converter.ConvertMapToAgrTypeValue(snapShot)
		s.Repo.Set(listOfAgr)
		s.limit = new(int64)
	}
}

func (s *AggregatorServi) pprint() {
	fmt.Printf("Result: %v\n\r", s.store)
}
