package repository

import (
	"dashboard"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user dashboard.User) (int, error)
	GetUser(username, password string) (dashboard.User, error)
}

type Invoice interface {
}

type Repository struct {
	Authorization
	Invoice
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
