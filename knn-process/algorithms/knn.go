package algorithms

import (
	"fmt"
	"sort"

	"knn-process/models"
)

// KNNAlgorithm implementa o algoritmo KNN com Distância de Jaccard
type KNNAlgorithm struct{}

// NewKNNAlgorithm cria uma nova instância do algoritmo KNN
func NewKNNAlgorithm() *KNNAlgorithm {
	return &KNNAlgorithm{}
}

// ProductSimilarity representa a similaridade entre dois produtos
type ProductSimilarity struct {
	ProductID  int
	Similarity float64
}

// CalculateJaccardDistance calcula a distância de Jaccard entre dois conjuntos
// A distância de Jaccard é: 1 - (intersecção / união)
func (knn *KNNAlgorithm) CalculateJaccardDistance(setA, setB []int) float64 {
	if len(setA) == 0 && len(setB) == 0 {
		return 0.0
	}

	// Criar mapas para facilitar a busca
	mapA := make(map[int]bool)
	mapB := make(map[int]bool)

	for _, id := range setA {
		mapA[id] = true
	}
	for _, id := range setB {
		mapB[id] = true
	}

	// Calcular interseção
	intersection := 0
	for id := range mapA {
		if mapB[id] {
			intersection++
		}
	}

	// Calcular união
	union := len(mapA) + len(mapB) - intersection

	if union == 0 {
		return 0.0
	}

	// Similaridade de Jaccard = interseção / união
	// Distância de Jaccard = 1 - similaridade
	jaccardSimilarity := float64(intersection) / float64(union)
	return 1.0 - jaccardSimilarity
}

// CalculateJaccardSimilarity calcula a similaridade de Jaccard (inverso da distância)
func (knn *KNNAlgorithm) CalculateJaccardSimilarity(setA, setB []int) float64 {
	return 1.0 - knn.CalculateJaccardDistance(setA, setB)
}

// FindTopKSimilarProducts encontra os K produtos mais similares usando KNN
func (knn *KNNAlgorithm) FindTopKSimilarProducts(
	targetProductID int,
	productSets map[int][]int,
	k int,
) ([]ProductSimilarity, error) {
	
	targetSet, exists := productSets[targetProductID]
	if !exists {
		return nil, fmt.Errorf("produto alvo %d não encontrado", targetProductID)
	}

	var similarities []ProductSimilarity

	// Calcular similaridade com todos os outros produtos
	for productID, productSet := range productSets {
		if productID == targetProductID {
			continue // Pular o próprio produto
		}

		similarity := knn.CalculateJaccardSimilarity(targetSet, productSet)
		similarities = append(similarities, ProductSimilarity{
			ProductID:  productID,
			Similarity: similarity,
		})
	}

	// Ordenar por similaridade (maior primeiro)
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].Similarity > similarities[j].Similarity
	})

	// Retornar apenas os K primeiros
	if len(similarities) > k {
		similarities = similarities[:k]
	}

	return similarities, nil
}

// BuildProductSets constrói os conjuntos de produtos baseado nas vendas
func (knn *KNNAlgorithm) BuildProductSets(sales []models.Sale) map[int][]int {
	// Mapa para agrupar produtos por pedido
	orderProducts := make(map[int][]int)
	
	// Mapa para rastrear produtos únicos por pedido
	orderProductSet := make(map[int]map[int]bool)

	// Agrupar produtos por pedido
	for _, sale := range sales {
		if orderProductSet[sale.OrderID] == nil {
			orderProductSet[sale.OrderID] = make(map[int]bool)
		}
		orderProductSet[sale.OrderID][sale.ProductID] = true
	}

	// Converter para slice de IDs
	for orderID, productMap := range orderProductSet {
		var products []int
		for productID := range productMap {
			products = append(products, productID)
		}
		orderProducts[orderID] = products
	}

	// Construir conjunto de produtos que aparecem junto com cada produto
	productSets := make(map[int][]int)
	productCooccurrence := make(map[int]map[int]bool)

	// Para cada pedido, registrar co-ocorrência de produtos
	for _, products := range orderProducts {
		for i, productA := range products {
			if productCooccurrence[productA] == nil {
				productCooccurrence[productA] = make(map[int]bool)
			}
			
			// Adicionar todos os outros produtos do mesmo pedido
			for j, productB := range products {
				if i != j {
					productCooccurrence[productA][productB] = true
				}
			}
		}
	}

	// Converter para slice de IDs
	for productID, cooccurrenceMap := range productCooccurrence {
		var cooccurringProducts []int
		for cooccurringProductID := range cooccurrenceMap {
			cooccurringProducts = append(cooccurringProducts, cooccurringProductID)
		}
		productSets[productID] = cooccurringProducts
	}

	return productSets
}
