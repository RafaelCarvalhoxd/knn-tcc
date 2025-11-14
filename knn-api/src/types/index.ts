export interface KnnProduct {
  id: number;
  parentId: number;
  relatedId: number;
  order: number;
  similarity: number;
  createdAt: Date;
  updatedAt: Date;
}

export interface Product {
  id: number;
  name: string;
  description: string;
  category: string;
  price: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface RelatedProductData {
  related_id: number;
  order: number;
  similarity: number;
  related_name: string;
  related_description: string;
  related_category: string;
  related_price: string;
}

export interface RelatedProductsResponse {
  success: boolean;
  message: string;
  data: {
    parentProductId: number;
    parentProductName: string;
    relatedProducts: RelatedProductData[];
  };
  totalFound: number;
  timestamp: string;
}

