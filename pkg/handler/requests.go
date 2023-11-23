package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/atauov/kcrps"
)

const flask = "http://145.249.246.27:8080"

func sendInvoice(invoice *kcrps.Invoice) error {
	invoice.Account = invoice.Account[1:]
	jsonData, err := json.Marshal(invoice)
	if err != nil {
		return err
	}
	sendJson(flask+"/kaspi", jsonData)
	return nil
}

func sendJson(url string, data []byte) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Ответ сервера: %s", resp.Status)
}
