package v1

import (
	"faba_traning_project/internal/controllers"
	"faba_traning_project/internal/services"

	"github.com/go-chi/chi"
)

func ProductRoutes(r chi.Router, serviceContainer services.Container) {
	r.Route("/products", func(r chi.Router) {
		r.Post("/", controllers.CreateProduct(serviceContainer))
	})
}
