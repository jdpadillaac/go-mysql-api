package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}