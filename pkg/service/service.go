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
	SetInvoiceForCancel(userId, invoiceId int) error
	SetInvoiceForRefund(userId, invoiceId int) error
}

type PosInvoice interface {
	SendInvoice(userId int, invoice kcrps.Invoice) error
	CancelInvoice(userId, invoiceId int) error
	CancelPayment(userId, amount, isToday, invoiceId int) error
	CheckInvoices(userId, isToday int, invoices []string) error
	UpdateStatus(id, status, inWork int) error
	UpdateClientName(invoiceId int, clientName string) error
	GetInWorkInvoices(userId int) ([]kcrps.Invoice, error)
	GetInvoiceAmount(invoiceId int) (int, error)
}

type Service struct {
	Authorization
	Invoice
	PosInvoice
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Invoice:       NewInvoiceService(repos.Invoice),
		PosInvoice:    NewPosInvoiceService(repos.PosInvoice),
	}
}
