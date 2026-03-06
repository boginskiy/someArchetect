package response

import "net/http"

type Response interface {
	CreateObjJSON([]byte, http.ResponseWriter)
	SendBadRequest(msg string, res http.ResponseWriter)
}

type Resp struct {
}

func NewResp() *Resp {
	return &Resp{}
}

func (r *Resp) CreateObjJSON(data []byte, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(data)
}

func (r *Resp) SendBadRequest(msg string, res http.ResponseWriter) {
	http.Error(res, msg, http.StatusBadRequest)
}
