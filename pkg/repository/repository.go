package repository

import (
	"github.com/atauov/kcrps"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user kcrps.User) (int, error)
	GetUser(username, password string) (kcrps.User, error)
}

type Invoice interface {
	Create(userId int, invoice kcrps.Invoice) (int, error)
	GetAll(userId int) ([]kcrps.Invoice, error)
	GetById(userId, invoiceId int) (kcrps.Invoice, error)
	Cancel(userId, invoiceId int) error
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
