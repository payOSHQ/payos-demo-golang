package controllers

import (
	"io"
	"log"
	"net/http"
	"payos-demo/config"
	"strconv"
	"time"

	"github.com/payOSHQ/payos-lib-golang"
)

func init() {
	payos.Key(config.PAYOS_CLIENT_ID, config.PAYOS_API_KEY, config.PAYOS_CHECKSUM_KEY)
}

func GenerateNumber() int {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	millisStr := strconv.FormatInt(millis, 10)
	number, _ := strconv.Atoi(millisStr[len(millisStr)-6:])
	return number
}

func CreatePaymentLink(w http.ResponseWriter, r *http.Request) {
	body := payos.CheckoutRequestType{
		OrderCode: GenerateNumber(),
		Amount:    2000,
		Items: []payos.Item{
			{
				Name:     "My tom Hao Hao ly",
				Price:    2000,
				Quantity: 1,
			},
		},
		Description: "Thanh toan don hang",
		CancelUrl:   config.YOUR_DOMAIN + "/cancel/",
		ReturnUrl:   config.YOUR_DOMAIN + "/success/",
	}

	data, err := payos.CreatePaymentLink(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, data.CheckoutUrl, http.StatusFound)
}

func GetPaymentLinkInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("orderId")
	data, err := payos.GetPaymentLinkInformation(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf(data.Status)
}

func CancelPaymentLink(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("orderId")
	data, err := payos.CancelPaymentLink(orderId, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf(data.Status)
}

func ConfirmWebhook(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	webhookUrl := string(body)

	webhookUrlRes, err := payos.ConfirmWebhook(webhookUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Print(webhookUrlRes)
}
