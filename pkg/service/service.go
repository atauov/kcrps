package service

import (
	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
)

type Authorization interface {
	CreateUser(user kcrps.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUserIdByApiKey(api string) (int, error)
}

type Invoice interface {
	Create(invoice kcrps.Invoice) (int, error)
	GetAll(invoice kcrps.Invoice) ([]kcrps.Invoice, error)
	GetById(invoice kcrps.Invoice) (kcrps.Invoice, error)
	SetInvoiceForCancel(invoice kcrps.Invoice) error
	SetInvoiceForRefund(invoice kcrps.Invoice) error
}

type Service struct {
	Authorization
	Invoice
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Invoice:       NewInvoiceService(repos.Invoice),
	}
}
