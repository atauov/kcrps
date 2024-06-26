package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/repository"
	"github.com/sirupsen/logrus"
)

const (
	CreateInvoiceURL    = "http://localhost:8080/create-invoice"
	CancelInvoiceURL    = "http://localhost:8080/cancel-invoice"
	CancelPaymentURL    = "http://localhost:8080/cancel-payment"
	CheckInvoicesURL    = "http://localhost:8080/check-invoices"
	WebHookURL          = "http://192.168.0.5:1111/webhook"
	StatusInvoiceOk     = "Invoice successful sent"
	StatusNoAccount     = "No kaspi account on number"
	StatusPaymentOk     = "Payment successful"
	StatusInvoiceCancel = "Invoice has been cancelled"
	StatusPaymentRefund = "Refund successful"
)

type PosInvoiceService struct {
	repo repository.PosInvoice
}

type WebHook struct {
	Id         int    `json:"id"`
	Status     string `json:"status"`
	Account    string `json:"account"`
	ClientName string `json:"client-name"`
}

func NewPosInvoiceService(repo repository.PosInvoice) *PosInvoiceService {
	return &PosInvoiceService{repo: repo}
}

func (s *PosInvoiceService) SendInvoice(userId int, invoice kcrps.Invoice) error {
	invoiceForFlask := RequestInvoice{
		UserID:  userId,
		Account: invoice.Account[1:],
		Amount:  invoice.Amount,
		Message: invoice.Message,
	}

	jsonData, err := json.Marshal(invoiceForFlask)
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
		return err
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
			return err
		}
		if err = s.repo.UpdateClientName(invoice.Id, response.ClientName); err != nil {
			return err
		}
		if err = s.repo.UpdateStatus(invoice.Id, 1, 1); err != nil {
			return err
		}

		webhook(invoice.Id, StatusInvoiceOk, invoice.Account, response.ClientName)

		return nil
	} else if resp.StatusCode == http.StatusNotFound {
		invoice.InWork = 0

		webhook(invoice.Id, StatusNoAccount, invoice.Account, "unknown")

		return nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("error on pos, please try later")
	}

	return errors.New("unknown error")
}

func (s *PosInvoiceService) CancelInvoice(userId, invoiceId int) error {
	invoiceCancel := RequestCancelInvoice{
		UserID: userId,
		ID:     strconv.Itoa(invoiceId),
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
		if err = s.repo.UpdateStatus(invoiceId, 3, 0); err != nil {
			return err
		}

		webhook(invoiceId, StatusInvoiceCancel, "", "")

		return nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("error on pos, please try later")
	}

	return errors.New("unknown error")
}

func (s *PosInvoiceService) CancelPayment(userId, amount, isToday, invoiceId int) error {

	paymentCancel := RequestCancelPayment{
		UserID:  userId,
		IsToday: isToday,
		Amount:  amount,
		ID:      strconv.Itoa(invoiceId),
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
		if err = s.repo.UpdateStatus(invoiceId, 4, 0); err != nil {
			return err
		}

		webhook(invoiceId, StatusPaymentRefund, "", "")

		return nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("error on pos, please try later")
	}

	return errors.New("unknown error")
}

func (s *PosInvoiceService) CheckInvoices(userId, isToday int, IDs []string) error {
	invoicesForCheck := RequestCheck{
		UserID:  userId,
		IsToday: isToday,
		IDs:     IDs,
	}
	jsonData, err := json.Marshal(invoicesForCheck)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, CheckInvoicesURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при выполнении запроса: %v", err)
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
			return err
		}
		for k, v := range result {
			uuid, _ := strconv.Atoi(k)
			invoiceId := uuid - 100000
			if err = s.UpdateStatus(invoiceId, v, 0); err != nil {
				return err
			}
			switch v {
			case 2:
				webhook(invoiceId, StatusPaymentOk, "", "")
			case 1:
				webhook(invoiceId, StatusInvoiceCancel, "", "")
			}
		}

		return nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("error on pos, please try later")
	}

	return errors.New("unknown error")
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

func (s *PosInvoiceService) GetInvoiceAmount(invoiceId int) (int, error) {
	return s.repo.GetInvoiceAmount(invoiceId)
}

func webhook(invoiceId int, status string, account string, clientName string) {
	jsonWebHook, _ := json.Marshal(WebHook{
		Id:         invoiceId,
		Status:     status,
		Account:    account,
		ClientName: clientName,
	})
	client := &http.Client{}
	resp, _ := http.NewRequest(http.MethodPost, WebHookURL, bytes.NewBuffer(jsonWebHook))
	if _, err := client.Do(resp); err != nil {
		logrus.Error(err)
	}
}
