package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/httpbody/response"
	"faba_traning_project/internal/services"
	mockService "faba_traning_project/mocks/services"
	"faba_traning_project/testhelpers"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	timeData := time.Date(2020, 4, 19, 10, 37, 40, 0, time.UTC)

	tests := []struct {
		name                     string
		requestUser              request.CreateUser
		givenResponseUser        response.User
		givenError               error
		expectError              error
		expectedStatusCode       int
		expectedJSONResponsePath string
	}{
		{
			name: "create user success",
			requestUser: request.CreateUser{
				Name:  "Thormas",
				Email: "thormas@gmail.com",
			},
			givenResponseUser: response.User{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "Thormas",
				Email:     "thormas@gmail.com",
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
			expectedStatusCode:       http.StatusOK,
			expectedJSONResponsePath: "testdata/user/create_user_success.json",
		},
		{
			name: "bad request missing Name field",
			requestUser: request.CreateUser{
				Email: "thormas@gmail.com",
			},
			givenResponseUser: response.User{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "Thormas",
				Email:     "thormas@gmail.com",
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
			expectedStatusCode:       http.StatusBadRequest,
			expectedJSONResponsePath: "testdata/user/create_user_err_validate_name.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			ctx := context.Background()

			// mock user service
			mockUserService := new(mockService.IUser)
			mockUserService.On(
				"Create",
				mock.Anything, // context can be mocked
				tc.requestUser,
			).Return(
				tc.givenResponseUser,
				tc.givenError,
			)

			requestBody, err := json.Marshal(&tc.requestUser)
			require.NoError(t, err)

			serviceContainer := services.Container{
				User: mockUserService,
			}

			// when
			resp, jsonResp := testhelpers.JSONRequest(ctx, t, http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody), nil,
				func(req *http.Request, rec *httptest.ResponseRecorder) {
					CreateUser(serviceContainer)(rec, req)
				},
			)

			expectedJSON := testhelpers.LoadJSONFixture(t, tc.expectedJSONResponsePath)

			// then
			require.Equal(t, tc.expectedStatusCode, resp.StatusCode)
			require.Equal(t, expectedJSON, jsonResp)
		})
	}
}
