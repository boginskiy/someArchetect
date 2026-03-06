package repository

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/internal/logg"
	"aggregatorProject/internal/model"
	"sync"
)

type EventRepo struct {
	store  map[string]*model.AgrTypeValue
	cfg    config.Config
	logger logg.Logger
	mx     sync.RWMutex
}

func NewEventRepo(config config.Config, logger logg.Logger) *EventRepo {
	return &EventRepo{
		cfg:    config,
		logger: logger,
		store:  make(map[string]*model.AgrTypeValue, 10),
		mx:     sync.RWMutex{},
	}
}

func (r *EventRepo) Update(dataArr []*model.AgrTypeValue) {
	r.mx.Lock()
	defer r.mx.Unlock()

	for _, v := range dataArr {
		r.store[v.Type] = v
	}
}
