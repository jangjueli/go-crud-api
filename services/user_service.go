package services

import (
	"go-crud-api/models"
	"go-crud-api/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(user *models.User) error {
	return s.Repo.Delete(user)
}
