package service

import (
	"github.com/atauov/kcrps/pkg/repository"
)

type Authorization interface {
	CreateUser(user dashboard.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Invoice interface {
	Create(userId int, invoice dashboard.Invoice) (int, error)
	GetAll(userId int) ([]dashboard.Invoice, error)
	GetById(userId, invoiceId int) (dashboard.Invoice, error)
	Cancel(userId, invoiceId int) error
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
