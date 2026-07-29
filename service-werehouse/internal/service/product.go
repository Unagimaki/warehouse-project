package service

import (
	"context"
	"service-werehouse/internal/dto/request"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, req request.CreateProductRequest) error {
	return s.repo.CreateProduct(ctx, req)
}
