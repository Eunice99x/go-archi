package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	models "github.com/eunice99x/dingoSuck/internal/db/pgModels"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type userRepository struct {
	db boil.ContextExecutor
}

func NewUserRepository(db boil.ContextExecutor) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, name string) (*models.User, error) {
	user := &models.User{Name: name}
	err := user.Insert(ctx, r.db, boil.Infer())
	return user, err
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := models.Users().All(ctx, r.db)
	return users, err
}