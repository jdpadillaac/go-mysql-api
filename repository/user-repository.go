package repository

import (
	"database/sql"
	"github.com/jdpadillaac/go-mysql-api/models"
)

type UserRepository struct {
	db *sql.DB
}



func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) Save(user models.User) (models.User, error) {
	var userCreated models.User
	stmt, err := u.db.Prepare("INSERT INTO user (names, surnames, email, password) VALUES(?,?, ?, ?)")
	if err != nil {
		return userCreated, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Names, user.Surnames, user.Email, user.Password)
	if err != nil {
		return userCreated, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return userCreated, err
	}

	userCreated.ID = uint(id)
	return userCreated, nil

}

func (u *UserRepository) FindAll() ([]models.User, error) {
	panic("implement me")
}
