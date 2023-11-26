package service

import "github.com/atauov/kcrps"

type ResponseFromFlask struct {
	Status     string `json:"status"`
	ClientName string `json:"client-name"`
	Invoices   []kcrps.Invoice
}
