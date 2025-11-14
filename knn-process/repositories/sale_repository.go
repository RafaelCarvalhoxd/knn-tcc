package repositories

import (
	"database/sql"
	"fmt"

	"knn-process/models"
)

// SaleRepository gerencia operações com vendas
type SaleRepository struct {
	db *sql.DB
}

// NewSaleRepository cria uma nova instância do repositório de vendas
func NewSaleRepository(db *sql.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

// GetSalesLast3Months busca vendas dos últimos 3 meses
func (r *SaleRepository) GetSalesLast3Months() ([]models.Sale, error) {
	query := `
		SELECT id, sale_date, product_id, quantity, price, order_id
		FROM sales 
		WHERE sale_date >= DATE_SUB(NOW(), INTERVAL 3 MONTH)
		ORDER BY sale_date DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar vendas: %w", err)
	}
	defer rows.Close()

	var sales []models.Sale
	for rows.Next() {
		var sale models.Sale
		err := rows.Scan(
			&sale.ID,
			&sale.SaleDate,
			&sale.ProductID,
			&sale.Quantity,
			&sale.Price,
			&sale.OrderID,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear venda: %w", err)
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

// GetSalesByOrderID busca vendas por ID do pedido
func (r *SaleRepository) GetSalesByOrderID(orderID int) ([]models.Sale, error) {
	query := `
		SELECT id, sale_date, product_id, quantity, price, order_id
		FROM sales 
		WHERE order_id = ?
		ORDER BY product_id
	`

	rows, err := r.db.Query(query, orderID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar vendas por pedido: %w", err)
	}
	defer rows.Close()

	var sales []models.Sale
	for rows.Next() {
		var sale models.Sale
		err := rows.Scan(
			&sale.ID,
			&sale.SaleDate,
			&sale.ProductID,
			&sale.Quantity,
			&sale.Price,
			&sale.OrderID,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear venda: %w", err)
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

// GetDistinctOrderIDs busca IDs únicos de pedidos dos últimos 3 meses
func (r *SaleRepository) GetDistinctOrderIDs() ([]int, error) {
	query := `
		SELECT DISTINCT order_id
		FROM sales 
		WHERE sale_date >= DATE_SUB(NOW(), INTERVAL 3 MONTH)
		ORDER BY order_id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pedidos únicos: %w", err)
	}
	defer rows.Close()

	var orderIDs []int
	for rows.Next() {
		var orderID int
		err := rows.Scan(&orderID)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear pedido: %w", err)
		}
		orderIDs = append(orderIDs, orderID)
	}

	return orderIDs, nil
}
