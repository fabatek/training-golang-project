package services

import (
	"context"
	"errors"
	"faba_traning_project/internal/databases"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/httpbody/response"
	"faba_traning_project/internal/models"
	mockOrder "faba_traning_project/mocks/databases/order"
	mockOrderItem "faba_traning_project/mocks/databases/order_item"
	mockProduct "faba_traning_project/mocks/databases/product"
	"faba_traning_project/utils"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	timeData := time.Date(2020, 4, 19, 10, 37, 40, 0, time.UTC)

	tests := []struct {
		name                   string
		requestCreateOrder     request.CreateOrder
		expectedGetListProduct []models.Product
		expectCreateOrder      models.Order
		givenError             error
		expectError            error
		expectedResponse       interface{}
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
			expectedGetListProduct: []models.Product{
				{
					ID:       "3722144d-e3f5-4f59-9700-518e0b736522",
					Name:     "car",
					Price:    2000,
					Quantity: 10,
				},
				{
					ID:       "3722144d-e3f5-4f59-9700-518e0b736523",
					Name:     "truck",
					Price:    3000,
					Quantity: 50,
				},
			},
			expectCreateOrder: models.Order{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				UserID:    "82e6bec0-1704-4b98-84c6-98b097665f4e",
				CreatedAt: timeData,
				UpdatedAt: timeData,
			},
			expectedResponse: response.Order{
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
		},
		{
			name: "order quantity cannot be larger than stock",
			requestCreateOrder: request.CreateOrder{
				UserID: "82e6bec0-1704-4b98-84c6-98b097665f4e",
				OrderItem: []request.OrderItem{
					{
						ProductID: "3722144d-e3f5-4f59-9700-518e0b736522",
						Quantity:  50,
					},
				},
			},
			expectedGetListProduct: []models.Product{
				{
					ID:       "3722144d-e3f5-4f59-9700-518e0b736522",
					Name:     "car",
					Price:    2000,
					Quantity: 10,
				},
			},
			expectCreateOrder: models.Order{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				UserID:    "82e6bec0-1704-4b98-84c6-98b097665f4e",
				CreatedAt: timeData,
				UpdatedAt: timeData,
			},
			expectError: errors.New("The product car is in excess of the allowed quantity"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			utils.GetUUID = func() string { return tc.expectCreateOrder.ID }

			// given
			mockOrderRepository := new(mockOrder.Repository)
			mockOrderItemRepository := new(mockOrderItem.Repository)
			mockProductRepository := new(mockProduct.Repository)

			mockOrderRepository.On(
				"Create",
				mock.Anything, // context can be mocked
				mock.Anything,
				mock.Anything,
			).Return(
				tc.expectCreateOrder,
				tc.givenError,
			)

			mockOrderItemRepository.On(
				"CreateListOrderItem",
				mock.Anything, // context can be mocked
				mock.Anything,
				mock.Anything,
			).Return(
				tc.givenError,
			)

			mockProductRepository.On(
				"GetListProductByListID",
				mock.Anything, // context can be mocked
				mock.Anything,
			).Return(
				tc.expectedGetListProduct,
				tc.givenError,
			)

			// mock db
			dbmock, mock, _ := sqlmock.New()
			mock.ExpectBegin()
			mock.ExpectCommit()
			defer dbmock.Close()

			dbStore := databases.DBStore{
				Order:     mockOrderRepository,
				OrderItem: mockOrderItemRepository,
				Product:   mockProductRepository,
				DBconn:    dbmock,
			}

			// when
			c := &Order{dbStore}
			response, err := c.Create(context.Background(), tc.requestCreateOrder)
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
