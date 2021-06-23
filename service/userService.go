package service

import (
	"mysql_rest_api_users/domain"
)

type DefaultUserService struct {
	repo domain.UserRepository
}

type UserService interface {
	GetAll() ([]domain.UserStruct, error)
	GetBuyId(id string) (*domain.UserStruct, error)
	Insert(newUser *domain.UserStruct) (string, error)
	Update(newUser *domain.UserStruct) (string, error)
	Delete(id string) (string, error)
}

func (s DefaultUserService) GetAll() ([]domain.UserStruct, error) {
	return s.repo.GetAllUsersDB()
}

func (s DefaultUserService) GetBuyId(id string) (*domain.UserStruct, error) {
	return s.repo.GetUserById(id)
}

func (s DefaultUserService) Insert(newUser *domain.UserStruct) (string, error) {
	return s.repo.InsertUserDB(newUser)
}

func (s DefaultUserService) Update(newUser *domain.UserStruct) (string, error) {
	return s.repo.UpdateUserDB(newUser)
}

func (s DefaultUserService) Delete(id string) (string, error) {
	return s.repo.DeleteUserDB(id)
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repository}
}
