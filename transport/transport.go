package transport

import "github.com/go-chi/chi"

type Handler struct {
	Port int
	Mux  *chi.Mux
}

func Initialize(port int) *Handler {
	handler := &Handler{
		Port: port,
		Mux:  chi.NewRouter(),
	}
	return handler
}
