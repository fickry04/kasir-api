package service

import (
	"context"
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) error {
	return s.Repo.Create(ctx, product)
}

func (s *ProductService) GetProduct(ctx context.Context) ([]model.Product, error) {
	return s.Repo.FindAll(ctx)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *model.Product) error {
	return s.Repo.Update(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}

func (s *ProductService) GetProductById(ctx context.Context, id int) (*model.Product, error) {
	return s.Repo.FindById(ctx, id)
}
