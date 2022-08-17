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

type IProduct interface {
	Create(ctx context.Context, requestUser request.CreateProduct) (respUser response.Product, err error)
}

type Product struct {
	dbStore databases.DBStore
}

var productInstance IProduct

func NewProduct(dbStore databases.DBStore) IProduct {
	if productInstance == nil {
		productInstance = &Product{
			dbStore: dbStore,
		}
	}
	return productInstance
}

func (product *Product) Create(ctx context.Context, requestProduct request.CreateProduct) (respProduct response.Product, err error) {
	modelProduct := models.Product{
		ID:       utils.GetUUID(),
		Name:     requestProduct.Name,
		Price:    requestProduct.Price,
		Quantity: requestProduct.Quantity,
	}

	modelProduct, err = product.dbStore.Product.Create(ctx, modelProduct)
	if err != nil {
		log.Error().Err(err)
		return respProduct, err
	}

	respProduct = response.Product{
		ID:        modelProduct.ID,
		Name:      modelProduct.Name,
		Price:     modelProduct.Price,
		Quantity:  modelProduct.Quantity,
		CreateAt:  modelProduct.CreatedAt.String(),
		UpdatedAt: modelProduct.UpdatedAt.String(),
	}

	return respProduct, nil
}
