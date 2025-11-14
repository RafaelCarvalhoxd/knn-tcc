package repositories

import (
	"database/sql"
	"fmt"

	"knn-process/models"
)

// ProductRepository gerencia operações com produtos
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository cria uma nova instância do repositório de produtos
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetProductByID busca um produto por ID
func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := `
		SELECT id, name, description, category, price
		FROM products 
		WHERE id = ?
	`

	var product models.Product
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Category,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("produto não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar produto: %w", err)
	}

	return &product, nil
}

// GetAllProducts busca todos os produtos
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `
		SELECT id, name, description, category, price
		FROM products 
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Category,
			&product.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear produto: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}
