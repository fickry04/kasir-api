package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(nameFilter string) ([]models.Product_View, error) {
	query := `SELECT
	p.id as product_id,
	p.name as product_name,
	p.price as price,
	p.stock as stock,
	c.name as category
	FROM products p JOIN categories c ON p.category_id = c.id
	`

	args := []any{}
	if nameFilter != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product_View, 0)
	for rows.Next() {
		var p models.Product_View
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := `INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryId).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (r *ProductRepository) GetById(id int) (*models.Product_View, error) {
	query := `SELECT 
	p.id as product_id,
	p.name as name,
	p.price as price,
	p.stock as stock,
	c.name as category
	FROM products p JOIN categories c ON p.category_id = c.id
	WHERE p.id = $1;`
	var p models.Product_View
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category)

	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepository) Update(product *models.Product) error {
	query := `UPDATE products SET name=$1, price=$2, stock=$3, category_id=$4 WHERE id=$5`
	result, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryId, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}
