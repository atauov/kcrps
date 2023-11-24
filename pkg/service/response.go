package service

import "github.com/atauov/kcrps"

type ResponseFromFlask struct {
	Status   int `json:"status"`
	Invoices []kcrps.Invoice
}
