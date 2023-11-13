package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Invoice interface {
}

type Repository struct {
	Authorization
	Invoice
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
