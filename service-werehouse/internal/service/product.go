package service

import (
	"context"
	"service-werehouse/internal/dto/request"
	apperrors "service-werehouse/internal/errors"
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
	if err := s.repo.CreateProduct(ctx, req); err != nil {
		return apperrors.Wrap("create product failed", err)
	}
	return nil
}
