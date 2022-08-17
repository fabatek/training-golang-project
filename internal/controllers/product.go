package controllers

import (
	"encoding/json"
	"errors"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/services"
	"faba_traning_project/utils"
	"net/http"

	"github.com/rs/zerolog/log"
)

func CreateProduct(serviceContainer services.Container) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		argumentHttp := utils.ArgumentHttp{W: w}

		var createProductRequest request.CreateProduct
		if r.Body == nil {
			err := errors.New(string("empty body"))
			log.Error().Err(err)

			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&createProductRequest); err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		if err := utils.Validate(createProductRequest); err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		respProduct, err := serviceContainer.Product.Create(ctx, createProductRequest)
		if err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusInternalServerError, argumentHttp, err)
			return
		}

		utils.SuccessJSON(ctx, http.StatusOK, argumentHttp, respProduct)
	})
}
