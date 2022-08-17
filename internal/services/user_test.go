package services

import (
	"context"
	"errors"
	"faba_traning_project/internal/databases"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/httpbody/response"
	"faba_traning_project/internal/models"
	mockUser "faba_traning_project/mocks/databases/user"
	"faba_traning_project/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	timeData := time.Date(2020, 4, 19, 10, 37, 40, 0, time.UTC)

	tests := []struct {
		name              string
		requestCreateUser request.CreateUser
		givenModelUser    models.User
		expectModelUser   models.User
		givenError        error
		expectError       error
		expectedResponse  interface{}
	}{
		{
			name: "create user success",
			requestCreateUser: request.CreateUser{
				Name:  "Thormas",
				Email: "thormas@gmail.com",
			},
			givenModelUser: models.User{
				ID:    "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:  "Thormas",
				Email: "thormas@gmail.com",
			},
			expectModelUser: models.User{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "Thormas",
				Email:     "thormas@gmail.com",
				CreatedAt: timeData,
				UpdatedAt: timeData,
			},
			expectedResponse: response.User{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "Thormas",
				Email:     "thormas@gmail.com",
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
		},
		{
			name: "create user failed with internal server error",
			requestCreateUser: request.CreateUser{
				Name:  "Thormas",
				Email: "thormas@gmail.com",
			},
			givenModelUser: models.User{
				ID:    "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:  "Thormas",
				Email: "thormas@gmail.com",
			},
			givenError:  errors.New("db connect rejected"),
			expectError: errors.New("db connect rejected"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			utils.GetUUID = func() string { return tc.givenModelUser.ID }

			// given
			mockUserRepository := new(mockUser.Repository)
			mockUserRepository.On(
				"Create",
				mock.Anything, // context can be mocked
				tc.givenModelUser,
			).Return(
				tc.expectModelUser,
				tc.givenError,
			)

			dbStore := databases.DBStore{
				User: mockUserRepository,
			}

			// when
			c := &User{dbStore}
			response, err := c.Create(context.Background(), tc.requestCreateUser)
			if err != nil {
				// then
				require.Equal(t, tc.expectError, err)
			} else {
				// then
				require.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
