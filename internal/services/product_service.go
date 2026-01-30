package services

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *model.Product) error {
	return s.repo.Create(data)
}

func (s *ProductService) GetById(id int) (*model.Product, error) {
	return s.repo.GetById(id)
}

func (s *ProductService) Update(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
