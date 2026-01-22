package service

// Berisi service category
// namun hanya statis dan tidak terhubung ke database/repository
// (untuk pembelajaran)

import (
	"context"
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *model.Category) {
	s.Repo.Create(ctx, category)
}

func (s *CategoryService) GetCategory(ctx context.Context) []model.Category {
	return s.Repo.FindAll(ctx)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category *model.Category) bool {
	return s.Repo.Update(ctx, category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) bool {
	return s.Repo.Delete(ctx, id)
}

func (s *CategoryService) GetCategoryById(ctx context.Context, id int) *model.Category {
	return s.Repo.FindById(ctx, id)
}
