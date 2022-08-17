package routes

import (
	v1 "faba_traning_project/cmd/handlers/routes/v1"
	"faba_traning_project/internal/controllers"
	"faba_traning_project/internal/services"
	"net/http"

	"github.com/go-chi/chi"
)

// Route function
func Route(serviceContainer services.Container) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/health-check", controllers.HealthCheck)

		v1.Routes(r, serviceContainer)
	})

	return r
}
