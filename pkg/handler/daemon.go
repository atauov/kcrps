package handler

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const TimeOutSec = 10

//TODO service for checking status
//TODO service for check db for invoices with in_work == 1
//send all invoices to Flask
//TODO service for canceling invoices with date > 24 hours

func (h *Handler) Daemon(posIDs []int) {
	logrus.Println("daemon started")

	for {
		for _, posID := range posIDs {
			// logrus.Printf("unit of pos: %d daemon started", posID)
			if _, exists := h.mutexes[posID]; !exists {
				h.mutexes[posID] = &sync.Mutex{}
			}
			go h.allOperations(posID)
		}
	}
}

func (h *Handler) allOperations(posID int) {
	mutex := h.mutexes[posID]
	mutex.Lock()
	defer mutex.Unlock()
	invoices, err := h.services.GetInWorkInvoices(posID)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(invoices)
	var forCheck []string
	for _, invoice := range invoices {
		switch invoice.Status {
		case 0:
			if err = h.services.SendInvoice(posID, invoice); err != nil {
				logrus.Error(err)
			}
		case 1:
			forCheck = append(forCheck, strconv.Itoa(invoice.UUID))
		case 3:
			if err = h.services.CancelInvoice(posID, invoice.Id); err != nil {
				logrus.Error(err)
			}
		case 4:
			amount, err := h.services.GetInvoiceAmount(invoice.Id)
			if err != nil {
				logrus.Error()
				continue
			}
			if err = h.services.CancelPayment(posID, amount, 1, invoice.Id); err != nil {
				logrus.Error(err)
			}
		}
	}
	fmt.Println(forCheck)
	if len(forCheck) > 0 {
		if err = h.services.CheckInvoices(posID, 1, forCheck); err != nil {
			logrus.Error(err)
		}
	}

	time.Sleep(TimeOutSec * time.Second)
}
