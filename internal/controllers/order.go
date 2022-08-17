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

func CreateOrder(serviceContainer services.Container) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		argumentHttp := utils.ArgumentHttp{W: w}

		var createOrderRequest request.CreateOrder
		if r.Body == nil {
			err := errors.New(string("empty body"))
			log.Error().Err(err)

			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&createOrderRequest); err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		if err := utils.Validate(createOrderRequest); err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		respOrder, err := serviceContainer.Order.Create(ctx, createOrderRequest)
		if err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusInternalServerError, argumentHttp, err)
			return
		}

		utils.SuccessJSON(ctx, http.StatusOK, argumentHttp, respOrder)
	})
}
