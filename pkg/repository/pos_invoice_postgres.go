package repository

import (
	"bytes"
	"encoding/json"
	"github.com/atauov/kcrps"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

const flask = "http://localhost:8080"

type PosInvoicePostgres struct {
	db *sqlx.DB
}

func NewPosInvoicePostgres(db *sqlx.DB) *PosInvoicePostgres {
	return &PosInvoicePostgres{db: db}
}

func (r *PosInvoicePostgres) SendInvoice(userId int, invoice kcrps.Invoice) error {
	invoice.Account = invoice.Account[1:]
	jsonData, err := json.Marshal(invoice)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, flask+"/create_invoice", bytes.NewBuffer(jsonData))
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

func (r *PosInvoicePostgres) CancelInvoice(userId, invoiceId int) error {
	return nil
}

func (r *PosInvoicePostgres) CancelPayment(userId, invoiceId int) error {
	return nil
}
