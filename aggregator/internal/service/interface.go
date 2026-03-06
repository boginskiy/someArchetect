package service

import "aggregatorProject/internal/model"

type UserService interface {
	Create(*model.User) (*model.User, error)
	Get(ID int) (*model.User, error)
}
