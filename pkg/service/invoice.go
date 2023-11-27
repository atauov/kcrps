package service

import (
	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
)

type InvoiceService struct {
	repo repository.Invoice
}

func NewInvoiceService(repo repository.Invoice) *InvoiceService {
	return &InvoiceService{repo: repo}
}

func (s *InvoiceService) Create(userId int, invoice kcrps.Invoice) (int, error) {
	return s.repo.Create(userId, invoice)
}

func (s *InvoiceService) GetAll(userId int) ([]kcrps.Invoice, error) {
	return s.repo.GetAll(userId)
}

func (s *InvoiceService) GetById(userId, invoiceId int) (kcrps.Invoice, error) {
	return s.repo.GetById(userId, invoiceId)
}

func (s *InvoiceService) SetInvoiceForCancel(userId, invoiceId int) error {
	return s.repo.SetInvoiceForCancel(userId, invoiceId)
}

func (s *InvoiceService) SetInvoiceForRefund(userId, invoiceId int) error {
	return s.repo.SetInvoiceForRefund(userId, invoiceId)
}
