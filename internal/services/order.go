package services

import (
	"context"
	"errors"
	"faba_traning_project/internal/databases"
	"faba_traning_project/internal/httpbody/request"
	"faba_traning_project/internal/httpbody/response"
	"faba_traning_project/internal/models"
	"faba_traning_project/utils"
	"fmt"

	"github.com/rs/zerolog/log"
)

type IOrder interface {
	Create(ctx context.Context, requestUser request.CreateOrder) (respUser response.Order, err error)
}

type Order struct {
	dbStore databases.DBStore
}

var orderInstance IOrder

// Singleton
func NewOrder(dbStore databases.DBStore) IOrder {
	if orderInstance == nil {
		orderInstance = &Order{
			dbStore: dbStore,
		}
	}
	return orderInstance
}

func (order *Order) Create(ctx context.Context, requestOrder request.CreateOrder) (respOrder response.Order, err error) {
	mapIdAndQuantity, ids := getProductIdFromOrderItem(requestOrder.OrderItem)

	listProduct, err := order.dbStore.Product.GetListProductByListID(ctx, ids)
	if err != nil {
		log.Error().Err(err).Send()
		return respOrder, err
	}

	if err := verifyOrderQuantity(listProduct, mapIdAndQuantity); err != nil {
		log.Error().Err(err).Send()
		return respOrder, err
	}

	modelOrder := models.Order{
		ID:     utils.GetUUID(),
		UserID: requestOrder.UserID,
	}

	tx, err := order.dbStore.DBconn.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Send()
		return respOrder, err
	}
	defer tx.Rollback()

	modelOrder, err = order.dbStore.Order.Create(ctx, tx, modelOrder)
	if err != nil {
		log.Error().Err(err).Send()
		return respOrder, err
	}

	listOrderItems := []models.OrderItem{}
	for _, item := range requestOrder.OrderItem {
		listOrderItems = append(listOrderItems, models.OrderItem{
			ID:        utils.GetUUID(),
			OrderID:   modelOrder.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	if err := order.dbStore.OrderItem.CreateListOrderItem(ctx, tx, listOrderItems); err != nil {
		log.Error().Err(err).Send()
		return respOrder, err
	}

	if err = tx.Commit(); err != nil {
		log.Error().Err(err).Send()
		return respOrder, err
	}

	// TODO minus the number of products in stock

	respListOrderItems := []response.OrderItem{}
	for _, item := range listOrderItems {
		respListOrderItems = append(respListOrderItems, response.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	respOrder = response.Order{
		ID:        modelOrder.ID,
		UserID:    modelOrder.UserID,
		OrderItem: respListOrderItems,
		CreateAt:  modelOrder.CreatedAt.String(),
		UpdatedAt: modelOrder.UpdatedAt.String(),
	}

	return respOrder, err
}

func getProductIdFromOrderItem(requestOrderItem []request.OrderItem) (mapIdAndQuantity map[string]int64, ids []string) {
	mapIdAndQuantity = make(map[string]int64)
	for _, item := range requestOrderItem {
		ids = append(ids, item.ProductID)
		mapIdAndQuantity[item.ProductID] = item.Quantity
	}
	return mapIdAndQuantity, ids
}

func verifyOrderQuantity(listProduct []models.Product, mapIdAndQuantity map[string]int64) error {
	for _, item := range listProduct {
		if mapIdAndQuantity[item.ID] > item.Quantity {
			err := errors.New(fmt.Sprintf("The product %v is in excess of the allowed quantity", item.Name))
			log.Error().Err(err).Send()
			return err
		}
	}
	return nil
}
