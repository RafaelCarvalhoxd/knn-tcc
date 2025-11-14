import { db } from '../config/database';
import { knnProducts } from '../config/schema';
import { eq, asc } from 'drizzle-orm';
import { KnnProduct } from '../types';

export class KnnProductRepository {
  async findRelatedByParentId(parentId: number): Promise<KnnProduct[]> {
    const result = await db
      .select()
      .from(knnProducts)
      .where(eq(knnProducts.parentId, parentId))
      .orderBy(asc(knnProducts.order));

    return result.map(item => ({
      id: item.id,
      parentId: item.parentId,
      relatedId: item.relatedId,
      order: item.order,
      similarity: item.similarity,
      createdAt: item.createdAt,
      updatedAt: item.updatedAt,
    }));
  }
}

