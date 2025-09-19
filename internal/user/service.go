package user

import (
	"context"
)

// type IService interface {
// 	GetUsers(ctx context.Context) ([]User, error)
// 	CreateUser(ctx context.Context, u User) (int64, error)
// }

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, u User) (int64, error) {
	return s.repo.Create(ctx, u)
}

func (s *Service) UpdateUser(ctx context.Context, u User) (*User, error) {
	return s.repo.Update(ctx, u)
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
