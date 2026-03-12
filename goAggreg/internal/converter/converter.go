package converter

import (
	"encoding/json"
	"goAggreg/internal/model"
	"io"
	"net/http"
)

type EventConverter interface {
	ConvertEventFromHTTPToModel(*http.Request) (*model.Event, error)
}

type EventConvert struct {
}

func NewEventConvert() *EventConvert {
	return &EventConvert{}
}

func (c *EventConvert) ConvertEventFromHTTPToModel(r *http.Request) (*model.Event, error) {
	dataByte, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	event := &model.Event{}
	err = json.Unmarshal(dataByte, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
