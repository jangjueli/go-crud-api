package auth

import (
	"context"
	"go-crud-api/internal/user"
)

type Service struct {
	repo        Repository
	userService user.Service
}

func NewService(repo Repository, us user.Service) *Service {
	return &Service{
		repo:        repo,
		userService: us,
	}
}

func (s *Service) GetUsers(ctx context.Context) ([]user.User, error) {
	return s.userService.GetUsers(ctx)
}
