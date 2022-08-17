package order

import (
	"context"
	"database/sql"
	"faba_traning_project/internal/databases/orm"
	"faba_traning_project/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, order models.Order) (models.Order, error)
}

// Management is the implementation of order
type Management struct {
	dbconn *sql.DB
}

// NewManagement initializes order
func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func translateOrmOrder(order *models.Order) orm.Order {
	ormOrder := orm.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		CreatedAt: order.CreatedAt,
		UpdatedAt: null.TimeFrom(order.UpdatedAt),
		DeletedAt: order.DeletedAt,
	}
	return ormOrder
}

func translateOrder(ormOrder *orm.Order) models.Order {
	order := models.Order{
		ID:        ormOrder.ID,
		UserID:    ormOrder.UserID,
		CreatedAt: ormOrder.CreatedAt,
		UpdatedAt: *ormOrder.UpdatedAt.Ptr(),
		DeletedAt: ormOrder.DeletedAt,
	}
	return order
}

// Create order item
func (management *Management) Create(ctx context.Context, tx *sql.Tx, order models.Order) (models.Order, error) {
	ormOrder := translateOrmOrder(&order)

	if tx != nil {
		if err := ormOrder.Insert(ctx, tx, boil.Infer()); err != nil {
			return models.Order{}, err
		}
	} else {
		if err := ormOrder.Insert(ctx, management.dbconn, boil.Infer()); err != nil {
			return models.Order{}, err
		}
	}

	return translateOrder(&ormOrder), nil
}
