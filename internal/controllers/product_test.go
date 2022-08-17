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

func TestCreateProduct(t *testing.T) {
	timeData := time.Date(2020, 4, 19, 10, 37, 40, 0, time.UTC)

	tests := []struct {
		name                     string
		requestProduct           request.CreateProduct
		givenResponseProduct     response.Product
		givenError               error
		expectError              error
		expectedStatusCode       int
		expectedJSONResponsePath string
	}{
		{
			name: "create user success",
			requestProduct: request.CreateProduct{
				Name:     "truck",
				Price:    90000,
				Quantity: 100,
			},
			givenResponseProduct: response.Product{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "truck",
				Price:     90000,
				Quantity:  100,
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
			expectedStatusCode:       http.StatusOK,
			expectedJSONResponsePath: "testdata/product/create_product_success.json",
		},
		{
			name: "bad request missing Name field",
			requestProduct: request.CreateProduct{
				Name:     "truck",
				Quantity: 100,
			},
			givenResponseProduct: response.Product{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "truck",
				Price:     90000,
				Quantity:  100,
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
			expectedStatusCode:       http.StatusBadRequest,
			expectedJSONResponsePath: "testdata/product/create_product_err_validate_price.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			ctx := context.Background()

			// mock product service
			mockProductService := new(mockService.IProduct)
			mockProductService.On(
				"Create",
				mock.Anything, // context can be mocked
				tc.requestProduct,
			).Return(
				tc.givenResponseProduct,
				tc.givenError,
			)

			requestBody, err := json.Marshal(&tc.requestProduct)
			require.NoError(t, err)

			serviceContainer := services.Container{
				Product: mockProductService,
			}

			// when
			resp, jsonResp := testhelpers.JSONRequest(ctx, t, http.MethodPost, "/api/v1/product", bytes.NewBuffer(requestBody), nil,
				func(req *http.Request, rec *httptest.ResponseRecorder) {
					CreateProduct(serviceContainer)(rec, req)
				},
			)

			expectedJSON := testhelpers.LoadJSONFixture(t, tc.expectedJSONResponsePath)

			// then
			require.Equal(t, tc.expectedStatusCode, resp.StatusCode)
			require.Equal(t, expectedJSON, jsonResp)
		})
	}
}
