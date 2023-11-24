package repository

import (
	"github.com/atauov/kcrps"
	"github.com/jmoiron/sqlx"
)

type PosInvoicePostgres struct {
	db *sqlx.DB
}

func NewPosInvoicePostgres(db *sqlx.DB) *PosInvoicePostgres {
	return &PosInvoicePostgres{db: db}
}

func (r *PosInvoicePostgres) SendInvoice(userId int, invoice kcrps.Invoice) error {
	return nil
}
func (r *PosInvoicePostgres) CancelInvoice(userId, invoiceId int) error {
	return nil
}
func (r *PosInvoicePostgres) CancelPayment(userId, invoiceId int) error {
	return nil
}
