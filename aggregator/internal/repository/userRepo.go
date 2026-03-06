package repository

import (
	"aggregator/cmd/config"
	"aggregator/internal/logg"
	"aggregator/internal/model"
	"context"
	"errors"
	"sync"
)

type UserRepo struct {
	Store  map[int]*model.User
	mx     sync.RWMutex
	cfg    config.Config
	logger logg.Logger
}

func NewUserRepo(config config.Config, logger logg.Logger) *UserRepo {
	return &UserRepo{
		Store:  make(map[int]*model.User, 10),
		cfg:    config,
		logger: logger,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	user.ID = len(r.Store) + 1
	r.Store[user.ID] = user
	return user, nil
}

func (r *UserRepo) Get(ctx context.Context, ID int) (*model.User, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	if user, ok := r.Store[ID]; ok {
		return user, nil
	}
	return nil, errors.New("user does not exist")
}
