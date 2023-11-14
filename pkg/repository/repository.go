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
	Create(userId int, invoice dashboard.Invoice) (int, error)
	GetAll(userId int) ([]dashboard.Invoice, error)
	GetById(userId, invoiceId int) (dashboard.Invoice, error)
}

type Repository struct {
	Authorization
	Invoice
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Invoice:       NewInvoicePostgres(db),
	}
}
