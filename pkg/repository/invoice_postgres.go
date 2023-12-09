package repository

import (
	"errors"
	"fmt"
	"github.com/atauov/kcrps"
	"github.com/jmoiron/sqlx"
)

type InvoicePostgres struct {
	db *sqlx.DB
}

func NewInvoicePostgres(db *sqlx.DB) *InvoicePostgres {
	return &InvoicePostgres{db: db}
}

func (r *InvoicePostgres) Create(invoice kcrps.Invoice) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var uuid int

	query := fmt.Sprintf("SELECT COALESCE(MAX(uuid), 100000) + 1 AS next_invoice_id FROM %s "+
		"WHERE pos_id=$1", invoicesTable)
	if err = tx.QueryRow(query, invoice.PosID).Scan(&uuid); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (created_at, account, amount, message, client_name, status, pos_id, uuid, user_id)"+
		"VALUES (NOW(), $1, $2, $3, $4, $5, $6, $7, $8)", invoicesTable)
	tx.QueryRow(query, invoice.Account, invoice.Amount, invoice.Message, "Unknown", STATUS1, invoice.PosID, uuid, invoice.UserID)
	if err = tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			return 0, err
		}

		return 0, err
	}

	return uuid, err
}

func (r *InvoicePostgres) GetAll(invoice kcrps.Invoice) ([]kcrps.Invoice, error) {
	var result []kcrps.Invoice

	query := fmt.Sprintf("SELECT uuid, created_at, account, amount, client_name, message, status FROM %s "+
		"WHERE pos_id=$1 AND user_id=$2 ORDER BY uuid DESC", invoicesTable)
	err := r.db.Select(&result, query, invoice.PosID, invoice.UserID)
	return result, err
}

func (r *InvoicePostgres) GetById(invoice kcrps.Invoice) (kcrps.Invoice, error) {
	var result kcrps.Invoice

	query := fmt.Sprintf("SELECT uuid, created_at, account, amount, client_name, message, status "+
		"FROM %s WHERE uuid=$1 AND pos_id=$2 AND user_id=$3", invoicesTable)
	err := r.db.Get(&result, query, invoice.ID, invoice.PosID, invoice.UserID)

	return result, err
}

func (r *InvoicePostgres) SetInvoiceForCancel(invoice kcrps.Invoice) error {
	var invoiceExist bool
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE uuid=$1 AND pos_id=$2 AND user_id=$3 AND status=$4)`,
		invoicesTable)
	if err := r.db.Get(&invoiceExist, query, invoice.ID, invoice.PosID, invoice.UserID, STATUS2); err != nil {
		return err
	}

	if !invoiceExist {
		return errors.New("cant find invoice for cancel")
	}

	query = fmt.Sprintf("UPDATE %s SET status=$1 WHERE uuid=$2 AND pos_id=$3 AND user_id=$4", invoicesTable)
	_, err := r.db.Exec(query, STATUS4, invoice.ID, invoice.PosID, invoice.UserID)

	return err
}

func (r *InvoicePostgres) SetInvoiceForRefund(invoice kcrps.Invoice) error {
	var invoiceExist bool
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE uuid = $1 AND pos_id=$2 AND user_id=$3 AND status=$4)`,
		invoicesTable)
	if err := r.db.Get(&invoiceExist, query, invoice.ID, invoice.PosID, invoice.UserID, STATUS9); err != nil {
		return err
	}

	if !invoiceExist {
		return errors.New("cant find invoice for refund")
	}

	query = fmt.Sprintf("UPDATE %s SET status=$1 WHERE uuid=$2 AND pos_id=$3 AND user_id=$4", invoicesTable)
	_, err := r.db.Exec(query, STATUS10, invoice.ID, invoice.PosID, invoice.UserID)

	return err
}
