package converter

import (
	"aggregatorProject/internal/model"
	"net/http"
)

type UserConverter interface {
	ConvertBytesToUser(*http.Request) (*model.User, error)
	ConvertUserToBytes(*model.User) ([]byte, error)
}

type EventConverter interface {
	ConvertMapToAgrTypeValue(map[string]*int64) []*model.AgrTypeValue
	ConvertBytesToEvent(req *http.Request) (*model.Event, error)
}
