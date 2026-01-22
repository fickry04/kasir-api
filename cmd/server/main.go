package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kasir-api/internal/config"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/router"
	"kasir-api/internal/service"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Gagal menyambung ke database")
	}

	// CRUD Product menggunakan database serverless neon
	productRepo := &repository.ProductRepository{DB: db}
	productService := &service.ProductService{Repo: productRepo}
	productHandler := &handler.ProductHandler{Service: productService}

	// Route Product : /api/product dan /api/product/
	router.ProductRegisterRoutes(productHandler)

	// CRUD Category menggunakan data statis
	categoryRepo := &repository.CategoryRepository{}
	categoryService := &service.CategoryService{Repo: categoryRepo}
	categoryHandler := &handler.CategoryHandler{Service: categoryService}

	// Route Category : /api/categories dan /api/categories/
	router.CategoryRegisterRoutes(categoryHandler)

	// Health Check: http://localhost:8080/api/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Api Running",
		})
	})
	fmt.Println("Server running di localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}

