package handler

import (
	"goAggreg/internal/converter"
	"goAggreg/internal/service"
	"net/http"
)

type Handler interface {
	Send(w http.ResponseWriter, r *http.Request)
}

type Handle struct {
	Converter converter.EventConverter
	Servicer  service.EventServicer
}

func NewHandle(converter converter.EventConverter, service service.EventServicer) *Handle {
	return &Handle{
		Converter: converter,
		Servicer:  service,
	}
}

func (h *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/event" && r.Method == "POST":
		h.Send(w, r)
	case r.URL.Path == "/event" && r.Method == "GET":
		// h.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handle) Send(w http.ResponseWriter, r *http.Request) {
	event, err := h.Converter.ConvertEventFromHTTPToModel(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.Servicer.Send(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
