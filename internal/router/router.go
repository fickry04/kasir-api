package router

import (
	"kasir-api/internal/handlers"
	"net/http"
)

func ProductRegisterRoutes(productHandler *handlers.ProductHandler) {
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetAll(w, r)
		case "POST":
			productHandler.Create(w, r)
		}
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetByID(w, r)
		case "PUT":
			productHandler.Update(w, r)
		case "DELETE":
			productHandler.Delete(w, r)
		}
	})
}

func CategoryRegisterRoutes(categoryHandler *handlers.CategoryHandler) {
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			categoryHandler.GetAll(w, r)
		case "POST":
			categoryHandler.Create(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			categoryHandler.GetByID(w, r)
		case "PUT":
			categoryHandler.Update(w, r)
		case "DELETE":
			categoryHandler.Delete(w, r)
		}
	})
}

func TransactionRegisterRoutes(transactionHandler *handlers.TransactionHandler) {
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)
	http.HandleFunc("/api/report/hari-ini", transactionHandler.SummaryToday)
	http.HandleFunc("/api/report", transactionHandler.Summary)
}
