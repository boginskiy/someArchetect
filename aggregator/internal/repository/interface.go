package repository

import (
	"aggregator/internal/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Get(ctx context.Context, ID int) (*model.User, error)
}

type EventRepository interface {
	Update([]*model.AgrTypeValue)
}
