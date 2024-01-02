package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/payOSHQ/payos-lib-golang"
)

func VerifyPaymentWebhookData(w http.ResponseWriter, r *http.Request) {
	var webhookDataReq payos.WebhookType

	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &webhookDataReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webhookData, err := payos.VerifyPaymentWebhookData(webhookDataReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// use webhookdata

	log.Print(webhookData.OrderCode)
}
