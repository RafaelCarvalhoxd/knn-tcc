import { db } from '../config/database';
import { products } from '../config/schema';
import { eq } from 'drizzle-orm';
import { Product } from '../types';

export class ProductRepository {
  async findById(id: number): Promise<Product | undefined> {
    const result = await db
      .select()
      .from(products)
      .where(eq(products.id, id))
      .limit(1);

    return result[0];
  }

  async findByIds(ids: number[]): Promise<Product[]> {
    if (ids.length === 0) {
      return [];
    }

    const result = await db
      .select()
      .from(products)
      .where(eq(products.id, ids[0]));

    return result;
  }
}

