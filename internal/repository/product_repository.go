package repository

import (
	"context"
	"database/sql"
	"kasir-api/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) Create(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO products (nama, harga, stok) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRowContext(ctx, query, product.Nama, product.Harga, product.Stok).Scan(&product.ID)
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	query := `SELECT * FROM products`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		rows.Scan(&product.ID, &product.Nama, &product.Harga, &product.Stok)
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *model.Product) error {
	query := `UPDATE products SET nama=$1, harga=$2, stok=$3 WHERE id=$4`
	_, err := r.DB.ExecContext(ctx, query, product.Nama, product.Harga, product.Stok, product.ID)

	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.DB.ExecContext(ctx, query, id)

	return err
}

func (r *ProductRepository) FindById(ctx context.Context, id int) (*model.Product, error) {
	query := `SELECT * FROM products WHERE id=$1`
	var product model.Product
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Nama, &product.Harga, &product.Stok)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
