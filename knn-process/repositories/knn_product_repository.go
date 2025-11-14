package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"knn-process/models"
)

// KNNProductRepository gerencia operações com resultados KNN
type KNNProductRepository struct {
	db *sql.DB
}

// NewKNNProductRepository cria uma nova instância do repositório KNN
func NewKNNProductRepository(db *sql.DB) *KNNProductRepository {
	return &KNNProductRepository{db: db}
}

// SaveKNNProduct salva um resultado KNN
func (r *KNNProductRepository) SaveKNNProduct(knnProduct *models.KNNProduct) error {
	query := `
		INSERT INTO knn_products (parent_id, related_id, ` + "`order`" + `, similarity, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	knnProduct.CreatedAt = now
	knnProduct.UpdatedAt = now

	result, err := r.db.Exec(query,
		knnProduct.ParentID,
		knnProduct.RelatedID,
		knnProduct.Order,
		knnProduct.Similarity,
		knnProduct.CreatedAt,
		knnProduct.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("erro ao salvar resultado KNN: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("erro ao obter ID inserido: %w", err)
	}

	knnProduct.ID = int(id)
	return nil
}

// ClearKNNProducts limpa todos os resultados KNN existentes
func (r *KNNProductRepository) ClearKNNProducts() error {
	query := `DELETE FROM knn_products`
	
	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao limpar resultados KNN: %w", err)
	}

	return nil
}

// GetKNNProductsByParentID busca resultados KNN por produto pai
func (r *KNNProductRepository) GetKNNProductsByParentID(parentID int) ([]models.KNNProduct, error) {
	query := `
		SELECT id, parent_id, related_id, ` + "`order`" + `, similarity, created_at, updated_at
		FROM knn_products 
		WHERE parent_id = ?
		ORDER BY ` + "`order`" + `
	`

	rows, err := r.db.Query(query, parentID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar resultados KNN: %w", err)
	}
	defer rows.Close()

	var knnProducts []models.KNNProduct
	for rows.Next() {
		var knnProduct models.KNNProduct
		err := rows.Scan(
			&knnProduct.ID,
			&knnProduct.ParentID,
			&knnProduct.RelatedID,
			&knnProduct.Order,
			&knnProduct.Similarity,
			&knnProduct.CreatedAt,
			&knnProduct.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear resultado KNN: %w", err)
		}
		knnProducts = append(knnProducts, knnProduct)
	}

	return knnProducts, nil
}
