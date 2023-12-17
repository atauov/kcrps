package service

import (
	"github.com/atauov/kcrps/models/request"
	"github.com/atauov/kcrps/models/response"
	"github.com/atauov/kcrps/pkg/repository"
)

type InvoiceService struct {
	repo repository.Invoice
}

func NewInvoiceService(repo repository.Invoice) *InvoiceService {
	return &InvoiceService{repo: repo}
}

func (s *InvoiceService) Create(invoice request.Invoice) (int, error) {
	return s.repo.Create(invoice)
}

func (s *InvoiceService) GetAll(invoice request.Invoice) ([]response.Invoice, error) {
	return s.repo.GetAll(invoice)
}

func (s *InvoiceService) GetById(invoice request.Invoice) (response.Invoice, error) {
	return s.repo.GetById(invoice)
}

func (s *InvoiceService) SetInvoiceForCancel(invoice request.Invoice) error {
	return s.repo.SetInvoiceForCancel(invoice)
}

func (s *InvoiceService) SetInvoiceForRefund(invoice request.Invoice) error {
	return s.repo.SetInvoiceForRefund(invoice)
}
