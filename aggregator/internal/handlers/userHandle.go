package handlers

import (
	"aggregatorProject/internal/converter"
	"aggregatorProject/internal/response"
	"aggregatorProject/internal/service"
	"net/http"
)

type UserHandle struct {
	Service   service.UserService
	Converter converter.UserConverter
	Response  response.Response
}

func NewUserHandle(srv service.UserService, conv converter.UserConverter, resp response.Response) *UserHandle {
	return &UserHandle{
		Service:   srv,
		Converter: conv,
		Response:  resp,
	}
}

func (h *UserHandle) Read(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandle) Create(w http.ResponseWriter, r *http.Request) {
	// Converter.
	user, err := h.Converter.ConvertBytesToUser(r)
	if err != nil {
		h.Response.SendBadRequest("user has not been created", w)
		return
	}

	// Service.
	user, err = h.Service.Create(user)
	if err != nil {
		h.Response.SendBadRequest("user has not been created", w)
		return
	}

	// Converter.
	dataBytes, err := h.Converter.ConvertUserToBytes(user)
	if err != nil {
		h.Response.SendBadRequest("user has not been created", w)
		return
	}

	// Response.
	h.Response.CreateObjJSON(dataBytes, w)
}
