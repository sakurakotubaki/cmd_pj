package usecase

import (
	"github.com/user/cmd_pj/internal/domain/model"
	"github.com/user/cmd_pj/internal/repository"
)

type UserUseCase struct {
	repo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(user *model.User) error {
	return u.repo.Create(user)
}

func (u *UserUseCase) GetAllUsers() ([]model.User, error) {
	return u.repo.GetAll()
}

func (u *UserUseCase) GetUserByID(id uint) (*model.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserUseCase) UpdateUser(user *model.User) error {
	return u.repo.Update(user)
}

func (u *UserUseCase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}
