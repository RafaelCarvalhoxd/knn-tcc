import { Request, Response } from 'express';
import { ProductService } from '../services/ProductService';
import { successResponse, errorResponse } from '../utils/responseHandler';

export class ProductController {
  private productService: ProductService;

  constructor() {
    this.productService = new ProductService();
  }

  getRelatedProducts = async (req: Request, res: Response): Promise<void> => {
    try {
      const productId = parseInt(req.params.id);

      if (isNaN(productId)) {
        errorResponse(res, 'ID do produto inválido', 400);
        return;
      }

      const result = await this.productService.getRelatedProducts(productId);

      if (!result) {
        errorResponse(res, 'Produto não encontrado', 404);
        return;
      }

      const message = `Encontrados ${result.totalFound} produtos casados`;
      
      successResponse(res, message, result);
    } catch (error) {
      console.error('Erro ao buscar produtos relacionados:', error);
      errorResponse(res, 'Erro interno do servidor', 500);
    }
  };
}

