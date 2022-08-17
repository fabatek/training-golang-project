package v1

import (
	"faba_traning_project/internal/services"

	"github.com/go-chi/chi"
)

func Routes(r chi.Router, serviceContainer services.Container) {
	r.Route("/v1", func(r chi.Router) {
		UserRoutes(r, serviceContainer)
		ProductRoutes(r, serviceContainer)
		OrderRoutes(r, serviceContainer)
	})
}
