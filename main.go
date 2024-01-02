package main

import (
	"html/template"
	"log"
	"net/http"
	"payos-demo/controllers"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request, template string) {
	err := templates.ExecuteTemplate(w, template+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "index")
	})

	http.HandleFunc("/cancel/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "cancel")
	})

	http.HandleFunc("/success/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "success")
	})

	http.HandleFunc("/create-payment-link", controllers.CreatePaymentLink)
	http.HandleFunc("/payment-link-info", controllers.GetPaymentLinkInfo)
	http.HandleFunc("/cancel-payment-link", controllers.CancelPaymentLink)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
