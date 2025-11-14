import { mysqlTable, int, double, datetime, varchar } from 'drizzle-orm/mysql-core';

export const knnProducts = mysqlTable('knn_products', {
  id: int('id').primaryKey().autoincrement(),
  parentId: int('parent_id').notNull(),
  relatedId: int('related_id').notNull(),
  order: int('order').notNull(),
  similarity: double('similarity').notNull(),
  createdAt: datetime('created_at').notNull(),
  updatedAt: datetime('updated_at').notNull(),
});

export const products = mysqlTable('products', {
  id: int('id').primaryKey().autoincrement(),
  name: varchar('name', { length: 255 }).notNull(),
  description: varchar('description', { length: 500 }).notNull(),
  category: varchar('category', { length: 100 }).notNull(),
  price: varchar('price', { length: 20 }).notNull(),
  createdAt: datetime('created_at').notNull(),
  updatedAt: datetime('updated_at').notNull(),
});

