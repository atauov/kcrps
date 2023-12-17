package repository

import (
	"github.com/atauov/kcrps/models/request"
	"github.com/atauov/kcrps/models/response"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user request.User) (int, error)
	GetUser(username, password string) (request.User, error)
	GetUserIdByApiKey(api string) (int, error)
}

type Invoice interface {
	Create(invoice request.Invoice) (int, error)
	GetAll(invoice request.Invoice) ([]response.Invoice, error)
	GetById(invoice request.Invoice) (response.Invoice, error)
	SetInvoiceForCancel(invoice request.Invoice) error
	SetInvoiceForRefund(invoice request.Invoice) error
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
