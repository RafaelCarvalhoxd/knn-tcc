import { Router } from 'express';
import productRoutes from './productRoutes';

const router = Router();

router.use('/api', productRoutes);

export default router;

