package models

import "time"

// Sale representa uma venda no sistema
type Sale struct {
	ID        int       `json:"id" db:"id"`
	SaleDate  time.Time `json:"sale_date" db:"sale_date"`
	ProductID int       `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Price     float64   `json:"price" db:"price"`
	OrderID   int       `json:"order_id" db:"order_id"`
}

// Product representa um produto
type Product struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Category    string `json:"category" db:"category"`
	Price       float64 `json:"price" db:"price"`
}

// KNNProduct representa o resultado da análise KNN
type KNNProduct struct {
	ID           int     `json:"id" db:"id"`
	ParentID     int     `json:"parent_id" db:"parent_id"`
	RelatedID    int     `json:"related_id" db:"related_id"`
	Order        int     `json:"order" db:"order"`
	Similarity   float64 `json:"similarity" db:"similarity"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// ProductSet representa um conjunto de produtos vendidos juntos
type ProductSet struct {
	ProductIDs []int
	Frequency  int
}
