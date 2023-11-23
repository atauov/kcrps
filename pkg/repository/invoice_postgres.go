package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type InvoicePostgres struct {
	db *sqlx.DB
}

func NewInvoicePostgres(db *sqlx.DB) *InvoicePostgres {
	return &InvoicePostgres{db: db}
}

func (r *InvoicePostgres) Create(userId int, invoice dashboard.Invoice) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createInvoiceQuery := fmt.Sprintf(
		"INSERT INTO %s (created_at, account, amount, client_name, status)"+
			"VALUES (NOW(), $1, $2, $3, $4) RETURNING id", invoicesTable)
	row := tx.QueryRow(createInvoiceQuery, invoice.Account, invoice.Amount, "No", 0)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersInvoicesQuery := fmt.Sprintf("INSERT INTO %s (user_id, invoice_id) VALUES ($1, $2)", usersInvoicesTable)
	_, err = tx.Exec(createUsersInvoicesQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	invoice.UUID = 100000 + id
	invoice.Message = strconv.Itoa(invoice.UUID) + " " + invoice.Message
	updateUuidQuery := fmt.Sprintf("UPDATE %s SET uuid=$1, message=$2 WHERE id=$3", invoicesTable)
	_, err = tx.Exec(updateUuidQuery, invoice.UUID, invoice.Message, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *InvoicePostgres) GetAll(userId int) ([]dashboard.Invoice, error) {
	var invoices []dashboard.Invoice

	query := fmt.Sprintf("SELECT il.id, il.uuid, il.created_at, il.account, il.amount, il.client_name, il.message,"+
		"il.status, il.in_work FROM %s il INNER JOIN %s ul on il.id =ul.invoice_id WHERE ul.user_id = $1 ORDER BY il.id DESC",
		invoicesTable, usersInvoicesTable)
	err := r.db.Select(&invoices, query, userId)

	return invoices, err
}

func (r *InvoicePostgres) GetById(userId, invoiceId int) (dashboard.Invoice, error) {
	var invoice dashboard.Invoice

	query := fmt.Sprintf("SELECT il.id,  il.uuid, il.created_at, il.account, il.amount, il.clent_name, il.message,"+
		" il.status, il.in_work FROM %s il INNER JOIN %s ul on il.id=ul.invoice_id WHERE ul.user_id = $1 AND ul.invoice_id = $2",
		invoicesTable, usersInvoicesTable)
	err := r.db.Get(&invoice, query, userId, invoiceId)

	return invoice, err
}

func (r *InvoicePostgres) Cancel(userId, invoiceId int) error {
	var invoice dashboard.Invoice
	var newStatusValue int

	query := fmt.Sprintf("SELECT il.status, il.in_work FROM %s il INNER JOIN %s ul on il.id=ul.invoice_id WHERE ul.user_id=$1 AND ul.invoice_id=$2",
		invoicesTable, usersInvoicesTable)
	if err := r.db.Get(&invoice, query, userId, invoiceId); err != nil {
		return err
	}

	if invoice.Status == 1 && invoice.InWork == 1 {
		newStatusValue = 3
	} else if invoice.Status == 2 {
		newStatusValue = 4
	} else {
		return errors.New("cant cancel this payment right now, try later please")
	}

	query = fmt.Sprintf("UPDATE %s SET status=$3, in_work=1 il USING %s ul WHERE il.id = ul.invoice_id AND ul.user_id=$1 AND ul.invoice_id=$2",
		invoicesTable, usersInvoicesTable)
	_, err := r.db.Exec(query, userId, invoiceId, newStatusValue)

	return err
}
