package product

import (
	"context"
	"database/sql"

	"faba_traning_project/internal/databases/orm"
	"faba_traning_project/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Create(ctx context.Context, user models.Product) (models.Product, error)
	GetProductById(ctx context.Context, productID string) (models.Product, error)
	GetListProductByListID(ctx context.Context, productIDs []string) (listProducts []models.Product, err error)
}

// Management is the implementation of product
type Management struct {
	dbconn *sql.DB
}

// NewManagement initializes product
func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func translateOrmProduct(product *models.Product) orm.Product {
	ormProduct := orm.Product{
		ID:        product.ID,
		Name:      null.StringFrom(product.Name),
		Price:     null.Float32From(product.Price),
		Quantity:  null.Int64From(product.Quantity),
		CreatedAt: product.CreatedAt,
		UpdatedAt: null.TimeFrom(product.UpdatedAt),
		DeletedAt: product.DeletedAt,
	}
	return ormProduct
}

func translateProduct(ormProduct *orm.Product) models.Product {
	product := models.Product{
		ID:        ormProduct.ID,
		Name:      ormProduct.Name.String,
		Price:     ormProduct.Price.Float32,
		Quantity:  ormProduct.Quantity.Int64,
		CreatedAt: ormProduct.CreatedAt,
		UpdatedAt: *ormProduct.UpdatedAt.Ptr(),
		DeletedAt: ormProduct.DeletedAt,
	}
	return product
}

// Create product
func (management *Management) Create(ctx context.Context, product models.Product) (models.Product, error) {

	ormProduct := translateOrmProduct(&product)
	if err := ormProduct.Insert(ctx, management.dbconn, boil.Infer()); err != nil {
		return models.Product{}, err
	}
	return translateProduct(&ormProduct), nil
}

func (management *Management) GetProductById(ctx context.Context, productID string) (models.Product, error) {

	ormProduct, err := orm.Products(orm.ProductWhere.ID.EQ(productID)).One(ctx, management.dbconn)
	if err != nil {
		return models.Product{}, err
	}

	return translateProduct(ormProduct), nil
}

func (management *Management) GetListProductByListID(ctx context.Context, productIDs []string) (listProducts []models.Product, err error) {

	ormProducts, err := orm.Products(orm.ProductWhere.ID.IN(productIDs)).All(ctx, management.dbconn)
	if err != nil {
		return listProducts, err
	}

	for _, item := range ormProducts {
		listProducts = append(listProducts, translateProduct(item))
	}

	return listProducts, err
}

func (management *Management) UpdateQuantityProductById(ctx context.Context, productID string, quantity int64) (models.Product, error) {

	ormProduct, err := orm.Products(orm.ProductWhere.ID.EQ(productID)).One(ctx, management.dbconn)
	if err != nil {
		return models.Product{}, err
	}

	return translateProduct(ormProduct), nil
}
