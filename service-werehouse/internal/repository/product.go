package repository

import (
	"context"
	"database/sql"
	"fmt"
	"service-werehouse/internal/dto/request"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, req request.CreateProductRequest) error {
	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query := `
		INSERT INTO products (name, brand, category, gender, volume_ml, description, barcode)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		req.ItemNumber,
		req.Brand,
		req.Type,
		req.Meta.Gender,
		req.Meta.VolumeML,
		req.Meta.Description,
		req.Meta.Barcode,
	)
	if err != nil {
		return fmt.Errorf("insert product: %w", err)
	}

	return nil
}
