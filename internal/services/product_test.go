package services

import (
	"context"
	"errors"
	"faba_traning_project/internal/databases"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/httpbody/response"
	"faba_traning_project/internal/models"
	mockProduct "faba_traning_project/mocks/databases/product"
	"faba_traning_project/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	timeData := time.Date(2020, 4, 19, 10, 37, 40, 0, time.UTC)

	tests := []struct {
		name                 string
		requestCreateProduct request.CreateProduct
		givenModelProduct    models.Product
		expectModelProduct   models.Product
		givenError           error
		expectError          error
		expectedResponse     interface{}
	}{
		{
			name: "create product success",
			requestCreateProduct: request.CreateProduct{
				Name:     "truck",
				Price:    90000,
				Quantity: 100,
			},
			givenModelProduct: models.Product{
				ID:       "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:     "truck",
				Price:    90000,
				Quantity: 100,
			},
			expectModelProduct: models.Product{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "truck",
				Price:     90000,
				Quantity:  100,
				CreatedAt: timeData,
				UpdatedAt: timeData,
			},
			expectedResponse: response.Product{
				ID:        "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:      "truck",
				Price:     90000,
				Quantity:  100,
				CreateAt:  timeData.String(),
				UpdatedAt: timeData.String(),
			},
		},
		{
			name: "create user failed with internal server error",
			requestCreateProduct: request.CreateProduct{
				Name:     "truck",
				Price:    90000,
				Quantity: 100,
			},
			givenModelProduct: models.Product{
				ID:       "8a7a084e-00ce-4bf4-9cdf-9329868f4af1",
				Name:     "truck",
				Price:    90000,
				Quantity: 100,
			},
			givenError:  errors.New("db connect rejected"),
			expectError: errors.New("db connect rejected"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			utils.GetUUID = func() string { return tc.givenModelProduct.ID }

			// given
			mockProductRepository := new(mockProduct.Repository)
			mockProductRepository.On(
				"Create",
				mock.Anything, // context can be mocked
				tc.givenModelProduct,
			).Return(
				tc.expectModelProduct,
				tc.givenError,
			)

			dbStore := databases.DBStore{
				Product: mockProductRepository,
			}

			// when
			c := &Product{dbStore}
			response, err := c.Create(context.Background(), tc.requestCreateProduct)
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
