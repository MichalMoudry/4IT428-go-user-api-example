package transport

import (
	"net/http"

	"github.com/go-chi/chi"
)

func getEmailFromURL(r *http.Request) string {
	email := chi.URLParam(r, "email")

	return email
}

func (handler *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	//var user dto.User
}
