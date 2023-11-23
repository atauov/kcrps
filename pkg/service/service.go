package service

import (
	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
)

type Authorization interface {
	CreateUser(user kcrps.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Invoice interface {
	Create(userId int, invoice kcrps.Invoice) (int, error)
	GetAll(userId int) ([]kcrps.Invoice, error)
	GetById(userId, invoiceId int) (kcrps.Invoice, error)
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
