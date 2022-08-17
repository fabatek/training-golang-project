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

func CreateUser(serviceContainer services.Container) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		argumentHttp := utils.ArgumentHttp{W: w}

		var createUserRequest request.CreateUser
		if r.Body == nil {
			err := errors.New(string("empty body"))
			log.Error().Err(err)

			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		if err := utils.Validate(createUserRequest); err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusBadRequest, argumentHttp, err)
			return
		}

		respUser, err := serviceContainer.User.Create(ctx, createUserRequest)
		if err != nil {
			log.Error().Err(err)
			utils.ErrorJSON(ctx, http.StatusInternalServerError, argumentHttp, err)
			return
		}

		utils.SuccessJSON(ctx, http.StatusOK, argumentHttp, respUser)
	})
}
