package service

import (
	"context"

	models "github.com/eunice99x/dingoSuck/internal/db/pgModels"
	"github.com/eunice99x/dingoSuck/internal/repository"
)


type UserService interface {
	CreateUser(ctx context.Context, name string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, name string) (*models.User, error) {
	return s.repo.CreateUser(ctx, name)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetAllUsers(ctx)
}