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
	Create(invoice kcrps.Invoice) (int, error)
	GetAll(invoice kcrps.Invoice) ([]kcrps.Invoice, error)
	GetById(invoice kcrps.Invoice) (kcrps.Invoice, error)
	SetInvoiceForCancel(invoice kcrps.Invoice) error
	SetInvoiceForRefund(invoice kcrps.Invoice) error
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
