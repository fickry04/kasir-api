package router

import (
	"kasir-api/internal/handler"
	"net/http"
)

func ProductRegisterRoutes(productHandler *handler.ProductHandler) {
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetProduct(w, r)
		case "POST":
			productHandler.CreateProduct(w, r)
		}
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetProductByID(w, r)
		case "PUT":
			productHandler.UpdateProduct(w, r)
		case "DELETE":
			productHandler.DeleteProduct(w, r)
		}
	})
}

func CategoryRegisterRoutes(categoryHandler *handler.CategoryHandler) {
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			categoryHandler.GetCategory(w, r)
		case "POST":
			categoryHandler.CreateCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			categoryHandler.GetCategoryByID(w, r)
		case "PUT":
			categoryHandler.UpdateCategory(w, r)
		case "DELETE":
			categoryHandler.DeleteCategory(w, r)
		}
	})
}
