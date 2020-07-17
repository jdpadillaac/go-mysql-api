package implementations

import (
	"github.com/jdpadillaac/go-mysql-api/models"
	"github.com/jdpadillaac/go-mysql-api/repository"
)

type UserService interface {
	Create(user models.User) (models.User, bool)
}

type userService struct {}

var (
	repo repository.UserRepository
)

func NewUserService(repository repository.UserRepository) UserService{
	repo = repository
	return  &userService{}
}

func (u *userService) Create(user models.User) (models.User, bool) {
	result,_ := repo.Save(user)
	return result, true
}


