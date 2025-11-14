import { Router } from "express";
import { ProductController } from "../controllers/ProductController";

const router = Router();
const productController = new ProductController();

router.get("/products/:id/related", productController.getRelatedProducts);

export default router;
