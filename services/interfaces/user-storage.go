package interfaces

import "github.com/jdpadillaac/go-mysql-api/models"

type UserStorage interface {
	Save(user models.User)(models.User, error)
	FindAll() ([]models.User, error)
}
