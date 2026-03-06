package service

import "aggregator/internal/model"

type UserService interface {
	Create(*model.User) (*model.User, error)
	Get(ID int) (*model.User, error)
}
