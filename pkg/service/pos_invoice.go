package service

import (
	"bytes"
	"encoding/json"
	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
	"log"
	"net/http"
)

const CreateInvoiceURL = "http://localhost:8080/create-invoice"
const CancelInvoiceURL = "http://localhost:8080/cancel-invoice"
const CancelPaymentURL = "http://localhost:8080/cancel-payment"
const CheckInvoicesURL = "http://localhost:8080/check-invoices"

type PosInvoiceService struct {
	repo repository.PosInvoice
}

func NewPosInvoiceService(repo repository.PosInvoice) *PosInvoiceService {
	return &PosInvoiceService{repo: repo}
}

func (s *PosInvoiceService) SendInvoice(userId int, invoice kcrps.Invoice) error {
	invoice.Account = invoice.Account[1:]
	jsonData, err := json.Marshal(invoice)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, CreateInvoiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Ответ сервера: %s", resp.Status)
	return nil
}

func (s *PosInvoiceService) CancelInvoice(userId, invoiceId int) error {
	return nil
}

func (s *PosInvoiceService) CancelPayment(userId, invoiceId int) error {
	return nil
}

func (s *PosInvoiceService) UpdateStatus(id, status, inWork int) error {
	return s.repo.UpdateStatus(id, status, inWork)
}

func (s *PosInvoiceService) UpdateClientName(invoiceId int, clientName string) error {
	return s.repo.UpdateClientName(invoiceId, clientName)

}

func (s *PosInvoiceService) GetInWorkInvoices(userId int) ([]kcrps.Invoice, error) {
	return s.repo.GetInWorkInvoices(userId)
}
