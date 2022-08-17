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

func TestCreateOrder(t *testing.T) {
	timeData := time.Date(2020, 4, 19, 10, 37, 40, 0, time.UTC)

	tests := []struct {
		name                     string
		requestCreateOrder       request.CreateOrder
		expectedResponseOrder    response.Order
		givenError               error
		expectError              error
		expectedStatusCode       int
		expectedJSONResponsePath string
	}{
		{
			name: "create order success",
			requestCreateOrder: request.CreateOrder{
				UserID: "82e6bec0-1704-4b98-84c6-98b097665f4e",
				OrderItem: []request.OrderItem{
					{
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736522",
						Quantity:  3,
					},
					{
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736523",
						Quantity:  4,
					},
				},
			},
			expectedResponseOrder: response.Order{
				ID:     "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				UserID: "82e6bec0-1704-4b98-84c6-98b097665f4e",
				OrderItem: []response.OrderItem{
					{
						ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736522",
						Quantity:  3,
					},
					{
						ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736523",
						Quantity:  4,
					},
				},
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
			expectedStatusCode:       http.StatusOK,
			expectedJSONResponsePath: "testdata/order/create_order_success.json",
		},
		{
			name: "bad request missing UserID field",
			requestCreateOrder: request.CreateOrder{
				OrderItem: []request.OrderItem{
					{
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736522",
						Quantity:  3,
					},
					{
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736523",
						Quantity:  4,
					},
				},
			},
			expectedStatusCode:       http.StatusBadRequest,
			expectedJSONResponsePath: "testdata/order/create_order_failed_validate.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			ctx := context.Background()

			// mock order service
			mockOrderService := new(mockService.IOrder)
			mockOrderService.On(
				"Create",
				mock.Anything, // context can be mocked
				tc.requestCreateOrder,
			).Return(
				tc.expectedResponseOrder,
				tc.givenError,
			)

			requestBody, err := json.Marshal(&tc.requestCreateOrder)
			require.NoError(t, err)

			serviceContainer := services.Container{
				Order: mockOrderService,
			}

			// when
			resp, jsonResp := testhelpers.JSONRequest(ctx, t, http.MethodPost, "/api/v1/orders", bytes.NewBuffer(requestBody), nil,
				func(req *http.Request, rec *httptest.ResponseRecorder) {
					CreateOrder(serviceContainer)(rec, req)
				},
			)

			expectedJSON := testhelpers.LoadJSONFixture(t, tc.expectedJSONResponsePath)

			// then
			require.Equal(t, tc.expectedStatusCode, resp.StatusCode)
			require.Equal(t, expectedJSON, jsonResp)
		})
	}
}
