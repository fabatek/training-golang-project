package user

import (
	"context"
	"database/sql"
	"faba_traning_project/internal/databases/orm"
	"faba_traning_project/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
}

// Management is the implementation of user
type Management struct {
	dbconn *sql.DB
}

// NewManagement initializes user
func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func translateOrmUser(user *models.User) orm.User {
	ormUser := orm.User{
		ID:        user.ID,
		Name:      null.StringFrom(user.Name),
		Email:     null.StringFrom(user.Email),
		CreatedAt: user.CreatedAt,
		UpdatedAt: null.TimeFrom(user.UpdatedAt),
		DeletedAt: user.DeletedAt,
	}
	return ormUser
}

func translateUser(ormUser *orm.User) models.User {
	user := models.User{
		ID:        ormUser.ID,
		Name:      ormUser.Name.String,
		Email:     ormUser.Email.String,
		CreatedAt: ormUser.CreatedAt,
		UpdatedAt: *ormUser.UpdatedAt.Ptr(),
		DeletedAt: ormUser.DeletedAt,
	}
	return user
}

// Create product
func (management *Management) Create(ctx context.Context, user models.User) (models.User, error) {

	ormUser := translateOrmUser(&user)
	if err := ormUser.Insert(ctx, management.dbconn, boil.Infer()); err != nil {
		return models.User{}, err
	}
	return translateUser(&ormUser), nil
}
