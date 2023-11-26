package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
	"github.com/sirupsen/logrus"
	"io"
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

func (s *PosInvoiceService) SendInvoice(userId int, invoice kcrps.Invoice) (kcrps.Invoice, error) {
	invoiceForFlask := RequestInvoice{
		UserID:  userId,
		Account: invoice.Account,
		Amount:  invoice.Amount,
		Message: invoice.Message,
	}
	jsonData, err := json.Marshal(invoiceForFlask)
	if err != nil {
		return invoice, err
	}
	req, err := http.NewRequest(http.MethodPost, CreateInvoiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return invoice, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		var response ResponseInvoice
		if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return invoice, err
		}
		invoice.ClientName = response.ClientName
		invoice.Status = 1
		return invoice, nil
	} else if resp.StatusCode == http.StatusNotFound {
		invoice.InWork = 0
		return invoice, nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return invoice, errors.New("error on pos, please try later")
	}

	return invoice, errors.New("unknown error")
}

func (s *PosInvoiceService) CancelInvoice(userId int, invoiceId string) error {
	invoiceCancel := RequestCancelInvoice{
		UserID: userId,
		ID:     invoiceId,
	}
	jsonData, err := json.Marshal(invoiceCancel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, CancelInvoiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("error on pos, please try later")
	}

	return errors.New("unknown error")
}

func (s *PosInvoiceService) CancelPayment(userId, isToday int, invoiceId string) error {
	paymentCancel := RequestCancelPayment{
		UserID:  userId,
		IsToday: isToday,
		ID:      invoiceId,
	}
	jsonData, err := json.Marshal(paymentCancel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, CancelPaymentURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("error on pos, please try later")
	}

	return errors.New("unknown error")
}

func (s *PosInvoiceService) CheckInvoices(userId, isToday int, invoices map[string]int) (map[string]int, error) {
	var IDs []string
	for k := range invoices {
		IDs = append(IDs, k)
	}
	invoicesForCheck := RequestCheck{
		UserID:  userId,
		IsToday: isToday,
		IDs:     IDs,
	}
	//result := make(map[string]int)
	jsonData, err := json.Marshal(invoicesForCheck)
	if err != nil {
		return invoices, err
	}
	req, err := http.NewRequest(http.MethodPost, CheckInvoicesURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return invoices, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		result := make(map[string]int)
		res, _ := io.ReadAll(resp.Body)
		if err = json.Unmarshal(res, &result); err != nil {
			return result, err
		}
		return result, nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return invoices, errors.New("error on pos, please try later")
	}

	return invoices, errors.New("unknown error")
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
