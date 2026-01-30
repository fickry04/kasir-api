package services

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *model.Category) error {
	return s.repo.Create(data)
}

func (s *CategoryService) GetById(id int) (*model.Category, error) {
	return s.repo.GetById(id)
}

func (s *CategoryService) Update(data *model.Category) error {
	return s.repo.Update(data)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
