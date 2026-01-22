package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/internal/model"
	"kasir-api/internal/service"
)

type ProductHandler struct {
	Service *service.ProductService
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	h.Service.CreateProduct(r.Context(), &product)

	w.Header().Set("Content-Type", "appication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	products, _ := h.Service.GetProduct(r.Context())
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateProduk model.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduk)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateProduk.ID = id
	h.Service.UpdateProduct(r.Context(), &updateProduk)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProduk)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	product, err := h.Service.GetProductById(r.Context(), id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteProduct(r.Context(), id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sukses delete",
	})
}
