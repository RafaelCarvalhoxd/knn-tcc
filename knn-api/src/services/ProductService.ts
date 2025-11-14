import { ProductRepository } from "../repositories/ProductRepository";
import { KnnProductRepository } from "../repositories/KnnProductRepository";
import { RelatedProductData } from "../types";

export class ProductService {
  private productRepository: ProductRepository;
  private knnProductRepository: KnnProductRepository;

  constructor() {
    this.productRepository = new ProductRepository();
    this.knnProductRepository = new KnnProductRepository();
  }

  async getRelatedProducts(parentId: number): Promise<{
    parentProductId: number;
    parentProductName: string;
    relatedProducts: RelatedProductData[];
    totalFound: number;
  } | null> {
    const parentProduct = await this.productRepository.findById(parentId);

    if (!parentProduct) {
      return null;
    }

    const knnRelations = await this.knnProductRepository.findRelatedByParentId(
      parentId
    );

    if (knnRelations.length === 0) {
      return {
        parentProductId: parentProduct.id,
        parentProductName: parentProduct.name,
        relatedProducts: [],
        totalFound: 0,
      };
    }

    const relatedProductsData: RelatedProductData[] = [];

    for (const relation of knnRelations) {
      const relatedProduct = await this.productRepository.findById(
        relation.relatedId
      );

      if (relatedProduct) {
        relatedProductsData.push({
          related_id: relation.relatedId,
          order: relation.order,
          similarity: relation.similarity,
          related_name: relatedProduct.name,
          related_description: relatedProduct.description,
          related_category: relatedProduct.category,
          related_price: relatedProduct.price,
        });
      }
    }

    return {
      parentProductId: parentProduct.id,
      parentProductName: parentProduct.name,
      relatedProducts: relatedProductsData,
      totalFound: relatedProductsData.length,
    };
  }
}
