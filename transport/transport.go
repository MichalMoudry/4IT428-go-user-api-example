package transport

import (
	"net/http"
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

	handler.Mux.Get("/health", health)
	handler.Mux.Route("/users", func(r chi.Router) {
		r.Get("/", handler.ListUsers)
		r.Post("/", handler.CreateUser)
		r.Route("/{email}", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Delete("/", handler.DeleteUser)
			r.Patch("/", handler.UpdateUser)
		})
	})

	return handler
}

func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
