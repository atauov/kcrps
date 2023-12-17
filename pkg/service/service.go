package service

import (
	"github.com/atauov/kcrps/models/request"
	"github.com/atauov/kcrps/models/response"
	"github.com/atauov/kcrps/pkg/repository"
)

type Authorization interface {
	CreateUser(user request.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUserIdByApiKey(api string) (int, error)
}

type Invoice interface {
	Create(invoice request.Invoice) (int, error)
	GetAll(invoice request.Invoice) ([]response.Invoice, error)
	GetById(invoice request.Invoice) (response.Invoice, error)
	SetInvoiceForCancel(invoice request.Invoice) error
	SetInvoiceForRefund(invoice request.Invoice) error
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
