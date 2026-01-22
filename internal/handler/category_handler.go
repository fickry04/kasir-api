package handler

// Berisi service category
// namun hanya statis dan tidak terhubung ke database/repository
// (untuk pembelajaran)

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/internal/model"
	"kasir-api/internal/service"
)

type CategoryHandler struct {
	Service *service.CategoryService
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	json.NewDecoder(r.Body).Decode(&category)
	h.Service.CreateCategory(r.Context(), &category)

	w.Header().Set("Content-Type", "appication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	categories := h.Service.GetCategory(r.Context())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory model.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateCategory.ID = id
	h.Service.UpdateCategory(r.Context(), &updateCategory)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateCategory)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	category := h.Service.GetCategoryById(r.Context(), id)
	if category == nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var deleteStatus bool
	deleteStatus = h.Service.DeleteCategory(r.Context(), id)
	if !deleteStatus {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sukses delete",
	})
}
