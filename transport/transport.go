package transport

import (
	"user-api/service"
	"user-api/transport/model/ioc"

	"github.com/go-chi/chi"
)

type Handler struct {
	Port    int
	Mux     *chi.Mux
	Service ioc.IService
}

func Initialize(port int) *Handler {
	handler := &Handler{
		Port:    port,
		Mux:     chi.NewRouter(),
		Service: service.CreateService(),
	}

	handler.Mux.Route("/users", func(r chi.Router) {
		//r.Get("/", handler.)
	})

	return handler
}
