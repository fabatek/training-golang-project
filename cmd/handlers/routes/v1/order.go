package v1

import (
	"faba_traning_project/internal/controllers"
	"faba_traning_project/internal/services"

	"github.com/go-chi/chi"
)

func OrderRoutes(r chi.Router, serviceContainer services.Container) {
	r.Route("/orders", func(r chi.Router) {
		r.Post("/", controllers.CreateOrder(serviceContainer))
	})
}
