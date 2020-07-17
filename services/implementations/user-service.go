package implementations

import (
	"github.com/jdpadillaac/go-mysql-api/models"
	"github.com/jdpadillaac/go-mysql-api/services/interfaces"
)

type UserService interface {
	Create(user models.User) (models.User, bool)
	GetByID(id string) (models.User, error)
}

type userService struct {
	storage interfaces.UserStorage
}


func NewUserService(repository interfaces.UserStorage) UserService{
	return  &userService{repository}
}

func (u *userService) Create(user models.User) (models.User, bool) {
	result,_ := u.storage.Save(user)
	return result, true
}


func (u *userService) GetByID(id string) (models.User, error) {
	panic("implement me")
}

