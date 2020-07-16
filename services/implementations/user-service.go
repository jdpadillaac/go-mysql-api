package implementations

import "github.com/jdpadillaac/go-mysql-api/services/interfaces"

type UserService struct {
	storage interfaces.UserStorage
}

func NewUserService(s interfaces.UserStorage) *UserService{
	return &UserService{s}
}