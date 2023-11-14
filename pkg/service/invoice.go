package service

import (
	"dashboard"
	"dashboard/pkg/repository"
)

type InvoiceService struct {
	repo repository.Invoice
}

func NewInvoiceService(repo repository.Invoice) *InvoiceService {
	return &InvoiceService{repo: repo}
}

func (s *InvoiceService) Create(userId int, invoice dashboard.Invoice) (int, error) {
	return s.repo.Create(userId, invoice)
}

func (s *InvoiceService) GetAll(userId int) ([]dashboard.Invoice, error) {
	return s.repo.GetAll(userId)
}
func (s *InvoiceService) GetById(userId, invoiceId int) (dashboard.Invoice, error) {
	return s.repo.GetById(userId, invoiceId)
}
