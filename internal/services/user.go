package services

import (
	"context"
	"faba_traning_project/internal/databases"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/httpbody/response"
	"faba_traning_project/internal/models"
	"faba_traning_project/utils"

	"github.com/rs/zerolog/log"
)

type IUser interface {
	Create(ctx context.Context, requestUser request.CreateUser) (respUser response.User, err error)
}

type User struct {
	dbStore databases.DBStore
}

var userInstance IUser

// Singleton
func NewUser(dbStore databases.DBStore) IUser {
	if userInstance == nil {
		userInstance = &User{
			dbStore: dbStore,
		}
	}
	return userInstance
}

func (user *User) Create(ctx context.Context, requestUser request.CreateUser) (respUser response.User, err error) {
	modelUser := models.User{
		ID:    utils.GetUUID(),
		Name:  requestUser.Name,
		Email: requestUser.Email,
	}

	modelUser, err = user.dbStore.User.Create(ctx, modelUser)
	if err != nil {
		log.Error().Err(err)
		return respUser, err
	}

	respUser = response.User{
		ID:        modelUser.ID,
		Name:      modelUser.Name,
		Email:     modelUser.Email,
		CreateAt:  modelUser.CreatedAt.String(),
		UpdatedAt: modelUser.UpdatedAt.String(),
	}

	return respUser, nil
}
