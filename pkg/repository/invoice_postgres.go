package repository

import (
	"dashboard"
	"fmt"
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
	createInvoiceQuery := fmt.Sprintf("INSERT INTO %s (amount, account, message) VALUES ($1, $2, $3) RETURNING id",
		invoicesTable)
	row := tx.QueryRow(createInvoiceQuery, invoice.Amount, invoice.Account, invoice.Message)
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

	return id, tx.Commit()
}

func (r *InvoicePostgres) GetAll(userId int) ([]dashboard.Invoice, error) {
	var invoices []dashboard.Invoice

	query := fmt.Sprintf("SELECT tl.id, tl.amount, tl.account, tl.message FROM %s tl INNER JOIN %s ul on "+
		"tl.id =ul.invoice_id WHERE ul.user_id = $1", invoicesTable, usersInvoicesTable)
	err := r.db.Select(&invoices, query, userId)

	return invoices, err
}

func (r *InvoicePostgres) GetById(userId, invoiceId int) (dashboard.Invoice, error) {
	var invoice dashboard.Invoice

	query := fmt.Sprintf("SELECT tl.id, tl.amount, tl.account, tl.message FROM %s tl INNER JOIN %s ul on "+
		"tl.id =ul.invoice_id WHERE ul.user_id = $1 AND ul.invoice_id = $2", invoicesTable, usersInvoicesTable)
	err := r.db.Get(&invoice, query, userId, invoiceId)

	return invoice, err
}