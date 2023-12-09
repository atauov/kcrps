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

func (s *InvoiceService) Create(invoice kcrps.Invoice) (int, error) {
	return s.repo.Create(invoice)
}

func (s *InvoiceService) GetAll(invoice kcrps.Invoice) ([]kcrps.Invoice, error) {
	return s.repo.GetAll(invoice)
}

func (s *InvoiceService) GetById(invoice kcrps.Invoice) (kcrps.Invoice, error) {
	return s.repo.GetById(invoice)
}

func (s *InvoiceService) SetInvoiceForCancel(invoice kcrps.Invoice) error {
	return s.repo.SetInvoiceForCancel(invoice)
}

func (s *InvoiceService) SetInvoiceForRefund(invoice kcrps.Invoice) error {
	return s.repo.SetInvoiceForRefund(invoice)
}
