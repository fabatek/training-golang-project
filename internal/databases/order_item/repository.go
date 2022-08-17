package order_item

import (
	"context"
	"database/sql"
	"faba_traning_project/internal/databases/orm"
	"faba_traning_project/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Create(ctx context.Context, orderItem models.OrderItem) (models.OrderItem, error)
	CreateListOrderItem(ctx context.Context, tx *sql.Tx, listOrderItem []models.OrderItem) (err error)
}

// Management is the implementation of order item
type Management struct {
	dbconn *sql.DB
}

// NewManagement initializes order item
func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func translateOrmOrderItem(orderItem *models.OrderItem) orm.OrderItem {
	ormOrderItem := orm.OrderItem{
		ID:        orderItem.ID,
		ProductID: null.StringFrom(orderItem.ProductID),
		Quantity:  null.Int64From(orderItem.Quantity),
		OrderID:   null.StringFrom(orderItem.OrderID),
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: null.TimeFrom(orderItem.UpdatedAt),
		DeletedAt: orderItem.DeletedAt,
	}
	return ormOrderItem
}

func translateOrderItem(ormOrderItem *orm.OrderItem) models.OrderItem {
	orderItem := models.OrderItem{
		ID:        ormOrderItem.ID,
		ProductID: ormOrderItem.ProductID.String,
		Quantity:  ormOrderItem.Quantity.Int64,
		OrderID:   ormOrderItem.OrderID.String,
		CreatedAt: ormOrderItem.CreatedAt,
		UpdatedAt: *ormOrderItem.UpdatedAt.Ptr(),
		DeletedAt: ormOrderItem.DeletedAt,
	}
	return orderItem
}

// Create order item
func (management *Management) Create(ctx context.Context, orderItem models.OrderItem) (models.OrderItem, error) {

	ormOrderItem := translateOrmOrderItem(&orderItem)
	if err := ormOrderItem.Insert(ctx, management.dbconn, boil.Infer()); err != nil {
		return models.OrderItem{}, err
	}
	return translateOrderItem(&ormOrderItem), nil
}

func (management *Management) CreateListOrderItem(ctx context.Context, tx *sql.Tx, listOrderItem []models.OrderItem) (err error) {

	sqlStr := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES ",
		orm.TableNames.OrderItems,
		orm.OrderItemColumns.ID,
		orm.OrderItemColumns.ProductID,
		orm.OrderItemColumns.Quantity,
		orm.OrderItemColumns.OrderID,
		orm.OrderItemColumns.CreatedAt,
		orm.OrderItemColumns.UpdatedAt,
	)
	vals := []interface{}{}

	var i int = 1
	for _, item := range listOrderItem {
		item.CreatedAt = time.Now().UTC()
		item.UpdatedAt = time.Now().UTC()

		sqlStr += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d),", i, i+1, i+2, i+3, i+4, i+5)
		vals = append(vals, item.ID, item.ProductID, item.Quantity, item.OrderID, item.CreatedAt, item.UpdatedAt)
		i += 6
	}

	//trim the last ,
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	//prepare the statement
	var stmt *sql.Stmt
	if tx != nil {
		stmt, err = tx.Prepare(sqlStr)
		if err != nil {
			return err
		}
	} else {
		stmt, err = management.dbconn.Prepare(sqlStr)
		if err != nil {
			return err
		}
	}

	//format all vals at once
	if _, err := stmt.Exec(vals...); err != nil {
		return err
	}

	stmt.Close()
	return err
}
