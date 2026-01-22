package repository

// Berisi service category
// namun hanya statis dan tidak terhubung ke database/repository
// (untuk pembelajaran)

import (
	"context"
	"kasir-api/internal/model"
)

// Contoh data kategori namun data statis saja (untuk pembelajaran)
var categories = []model.Category{
	{
		ID:          1,
		Name:        "Makanan",
		Description: "Produk makanan siap saji dan olahan",
	},
	{
		ID:          2,
		Name:        "Minuman",
		Description: "Minuman segar, panas, dan kemasan",
	},
	{
		ID:          3,
		Name:        "Pakaian",
		Description: "Pakaian pria, wanita, dan anak-anak",
	},
	{
		ID:          4,
		Name:        "Elektronik",
		Description: "Perangkat elektronik rumah tangga dan gadget",
	},
	{
		ID:          5,
		Name:        "Alat Tulis",
		Description: "Perlengkapan sekolah dan perkantoran",
	},
}

type CategoryRepository struct{}

func (r *CategoryRepository) Create(ctx context.Context, category *model.Category) {
	category.ID = len(categories) + 1
	categories = append(categories, *category)
}

func (r *CategoryRepository) FindAll(ctx context.Context) []model.Category {
	return categories
}

func (r *CategoryRepository) Update(ctx context.Context, category *model.Category) bool {
	for i := range categories {
		if categories[i].ID == category.ID {
			categories[i] = *category
			return true
		}
	}
	return false
}

func (r *CategoryRepository) Delete(ctx context.Context, id int) bool {
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return true
		}
	}

	return false
}

func (r *CategoryRepository) FindById(ctx context.Context, id int) *model.Category {
	for _, c := range categories {
		if c.ID == id {
			return &c
		}
	}

	return nil
}
