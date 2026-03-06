package handlers

import (
	"aggregator/internal/converter"
	"aggregator/internal/model"
	"context"
	"log"
	"net/http"
)

type EventHandle struct {
	Converter converter.EventConverter
	ctx       context.Context
	eventCh   chan<- *model.Event
}

func NewEventHandle(
	ctx context.Context,
	eventCh chan<- *model.Event,
	conv converter.EventConverter) *EventHandle {

	tmpEvent := &EventHandle{
		Converter: conv,
		ctx:       ctx,
		eventCh:   eventCh,
	}

	return tmpEvent
}

func (h *EventHandle) Listen(w http.ResponseWriter, r *http.Request) {
	event, err := h.Converter.ConvertBytesToEvent(r)
	if err != nil {
		log.Println(err)
		return
	}

	h.eventCh <- event

	// Потребляем данные по HTTP без ответа
}
