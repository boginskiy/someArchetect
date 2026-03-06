package converter

import (
	"aggregatorProject/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

type EventConvert struct {
}

func NewEventConvert() *EventConvert {
	return &EventConvert{}
}

func (c *EventConvert) ConvertBytesToEvent(req *http.Request) (*model.Event, error) {
	dataBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	event := &model.Event{}

	err = json.Unmarshal(dataBytes, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (c *EventConvert) ConvertMapToAgrTypeValue(data map[string]*int64) []*model.AgrTypeValue {
	result := make([]*model.AgrTypeValue, len(data))
	i := 0
	for k, v := range data {
		result[i] = &model.AgrTypeValue{Type: k, Value: *v}
		i++
	}
	return result
}
