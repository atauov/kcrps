package service

import (
	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
)

type PosInvoiceService struct {
	repo repository.PosInvoice
}

func NewPosInvoiceService(repo repository.PosInvoice) *PosInvoiceService {
	return &PosInvoiceService{repo: repo}
}

func (s *PosInvoiceService) SendInvoice(userId int, invoice kcrps.Invoice) error {
	return s.repo.SendInvoice(userId, invoice)
}

func (s *PosInvoiceService) CancelInvoice(userId, invoiceId int) error {
	return s.repo.CancelInvoice(userId, invoiceId)
}

func (s *PosInvoiceService) CancelPayment(userId, invoiceId int) error {
	return s.repo.CancelPayment(userId, Invoice)
}
