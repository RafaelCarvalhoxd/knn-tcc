package services

import (
	"database/sql"
	"fmt"
	"log"

	"knn-process/algorithms"
	"knn-process/models"
	"knn-process/repositories"
)

// KNNAnalysisService gerencia a análise KNN de produtos casados
type KNNAnalysisService struct {
	saleRepo        *repositories.SaleRepository
	productRepo     *repositories.ProductRepository
	knnProductRepo  *repositories.KNNProductRepository
	knnAlgorithm    *algorithms.KNNAlgorithm
}

// NewKNNAnalysisService cria uma nova instância do serviço
func NewKNNAnalysisService(db *sql.DB) *KNNAnalysisService {
	return &KNNAnalysisService{
		saleRepo:       repositories.NewSaleRepository(db),
		productRepo:    repositories.NewProductRepository(db),
		knnProductRepo: repositories.NewKNNProductRepository(db),
		knnAlgorithm:   algorithms.NewKNNAlgorithm(),
	}
}

// AnalyzeProductAssociations executa a análise completa de associações de produtos
func (s *KNNAnalysisService) AnalyzeProductAssociations() error {
	log.Println("Iniciando análise de associações de produtos...")

	// 1. Buscar vendas dos últimos 3 meses
	log.Println("Buscando vendas dos últimos 3 meses...")
	sales, err := s.saleRepo.GetSalesLast3Months()
	if err != nil {
		return fmt.Errorf("erro ao buscar vendas: %w", err)
	}

	if len(sales) == 0 {
		log.Println("Nenhuma venda encontrada nos últimos 3 meses")
		return nil
	}

	log.Printf("Encontradas %d vendas nos últimos 3 meses", len(sales))

	// 2. Construir conjuntos de produtos
	log.Println("Construindo conjuntos de produtos...")
	productSets := s.knnAlgorithm.BuildProductSets(sales)
	log.Printf("Conjuntos de produtos construídos para %d produtos", len(productSets))

	// 3. Limpar resultados anteriores
	log.Println("Limpando resultados KNN anteriores...")
	err = s.knnProductRepo.ClearKNNProducts()
	if err != nil {
		return fmt.Errorf("erro ao limpar resultados anteriores: %w", err)
	}

	// 4. Buscar todos os produtos para análise
	log.Println("Buscando produtos para análise...")
	products, err := s.productRepo.GetAllProducts()
	if err != nil {
		return fmt.Errorf("erro ao buscar produtos: %w", err)
	}

	log.Printf("Analisando %d produtos...", len(products))

	// 5. Para cada produto, encontrar os 5 mais similares
	processedCount := 0
	for _, product := range products {
		// Verificar se o produto tem vendas nos últimos 3 meses
		if _, hasSales := productSets[product.ID]; !hasSales {
			continue
		}

		log.Printf("Analisando produto %d: %s", product.ID, product.Name)

		// Encontrar os 5 produtos mais similares
		similarities, err := s.knnAlgorithm.FindTopKSimilarProducts(
			product.ID,
			productSets,
			5,
		)
		if err != nil {
			log.Printf("Erro ao analisar produto %d: %v", product.ID, err)
			continue
		}

		// Salvar resultados
		for order, similarity := range similarities {
			knnProduct := &models.KNNProduct{
				ParentID:   product.ID,
				RelatedID:  similarity.ProductID,
				Order:      order + 1, // Ordem começa em 1
				Similarity: similarity.Similarity,
			}

			err = s.knnProductRepo.SaveKNNProduct(knnProduct)
			if err != nil {
				log.Printf("Erro ao salvar resultado KNN para produto %d: %v", product.ID, err)
				continue
			}
		}

		processedCount++
		if processedCount%10 == 0 {
			log.Printf("Processados %d produtos...", processedCount)
		}
	}

	log.Printf("Análise concluída! Processados %d produtos", processedCount)
	return nil
}

// GetProductRecommendations busca recomendações para um produto específico
func (s *KNNAnalysisService) GetProductRecommendations(productID int) ([]models.KNNProduct, error) {
	return s.knnProductRepo.GetKNNProductsByParentID(productID)
}

// GetAnalysisStats retorna estatísticas da análise
func (s *KNNAnalysisService) GetAnalysisStats() (map[string]interface{}, error) {
	// Esta função pode ser expandida para incluir mais estatísticas
	stats := make(map[string]interface{})
	
	// Buscar vendas dos últimos 3 meses para estatísticas
	sales, err := s.saleRepo.GetSalesLast3Months()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar estatísticas de vendas: %w", err)
	}

	stats["total_sales"] = len(sales)
	
	// Contar produtos únicos
	uniqueProducts := make(map[int]bool)
	for _, sale := range sales {
		uniqueProducts[sale.ProductID] = true
	}
	stats["unique_products"] = len(uniqueProducts)

	return stats, nil
}
