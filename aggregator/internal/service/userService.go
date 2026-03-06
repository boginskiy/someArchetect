package service

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/internal/logg"
	"aggregatorProject/internal/model"
	"aggregatorProject/internal/repository"
	"context"
)

type UserServi struct {
	Repo   repository.UserRepository
	cfg    config.Config
	logger logg.Logger
	ctx    context.Context
}

func NewUserServi(
	ctx context.Context,
	config config.Config,
	logger logg.Logger,
	repo repository.UserRepository) *UserServi {

	return &UserServi{
		Repo:   repo,
		cfg:    config,
		logger: logger,
		ctx:    ctx,
	}
}

func (u *UserServi) Create(user *model.User) (*model.User, error) {
	return u.Repo.Create(u.ctx, user)
}

func (u *UserServi) Get(ID int) (*model.User, error) {
	return nil, nil
}
