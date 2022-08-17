package v1

import (
	"faba_traning_project/internal/controllers"
	"faba_traning_project/internal/services"

	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router, serviceContainer services.Container) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", controllers.CreateUser(serviceContainer))
	})
}
