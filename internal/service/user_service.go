package service

import (
	"context"

	"github.com/Sixchalice/go-cassandra-rest/internal/models"
	"github.com/Sixchalice/go-cassandra-rest/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
}

func (s *UserService) CreateUser(ctx context.Context, input CreateUserInput) error {
	return s.repo.InsertUser(ctx, input.ID, input.Name)
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, input UpdateUserInput) error {
	return s.repo.UpdateUser(ctx, id, input.Name)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
