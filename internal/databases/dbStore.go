package databases

import (
	"database/sql"
	"faba_traning_project/internal/databases/order"
	orderitem "faba_traning_project/internal/databases/order_item"
	"faba_traning_project/internal/databases/product"
	"faba_traning_project/internal/databases/user"
)

type DBStore struct {
	User      user.Repository
	Product   product.Repository
	Order     order.Repository
	OrderItem orderitem.Repository
	DBconn    *sql.DB
}
