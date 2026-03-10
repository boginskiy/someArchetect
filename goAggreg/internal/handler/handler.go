package handler

import "net/http"

// type Handler interface {
// 	Send(w http.ResponseWriter, r *http.Request)
// 	ServerHTTP(w http.ResponseWriter, r *http.Request)
// }

type Handle struct {
}

func NewHandle() *Handle {
	return &Handle{}
}

func (h *Handle) ServerHTTP(w http.ResponseWriter, r *http.Request) {
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

}
