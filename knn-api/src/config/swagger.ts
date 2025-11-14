import swaggerJsdoc from "swagger-jsdoc";

const options: swaggerJsdoc.Options = {
  definition: {
    openapi: "3.0.0",
    info: {
      title: "KNN API - Recomendação de Produtos",
      version: "1.0.0",
      description:
        "API para recomendação de produtos usando algoritmo KNN (K-Nearest Neighbors)",
      contact: {
        name: "API Support",
      },
    },
    servers: [
      {
        url: "http://localhost:3000",
        description: "Servidor de Desenvolvimento",
      },
    ],
    tags: [
      {
        name: "Products",
        description: "Endpoints relacionados a produtos e recomendações",
      },
    ],
    components: {
      schemas: {
        RelatedProduct: {
          type: "object",
          properties: {
            related_id: {
              type: "integer",
              description: "ID do produto relacionado",
              example: 12,
            },
            order: {
              type: "integer",
              description: "Ordem de relevância do produto relacionado",
              example: 1,
            },
            similarity: {
              type: "number",
              format: "double",
              description: "Índice de similaridade (0-1)",
              example: 0.9791666666666666,
            },
            related_name: {
              type: "string",
              description: "Nome do produto relacionado",
              example: "Cadeira Gamer Pro",
            },
            related_description: {
              type: "string",
              description: "Descrição do produto relacionado",
              example:
                "Cadeira gamer ergonômica com suporte lombar, reclinável",
            },
            related_category: {
              type: "string",
              description: "Categoria do produto relacionado",
              example: "Móveis",
            },
            related_price: {
              type: "string",
              description: "Preço do produto relacionado",
              example: "1299.90",
            },
          },
        },
        RelatedProductsData: {
          type: "object",
          properties: {
            parentProductId: {
              type: "integer",
              description: "ID do produto pai",
              example: 11,
            },
            parentProductName: {
              type: "string",
              description: "Nome/descrição do produto pai",
              example: "Gabinete Gamer RGB",
            },
            relatedProducts: {
              type: "array",
              items: {
                $ref: "#/components/schemas/RelatedProduct",
              },
            },
          },
        },
        SuccessResponse: {
          type: "object",
          properties: {
            success: {
              type: "boolean",
              example: true,
            },
            message: {
              type: "string",
              example: "Encontrados 5 produtos casados",
            },
            data: {
              $ref: "#/components/schemas/RelatedProductsData",
            },
            totalFound: {
              type: "integer",
              description: "Total de produtos relacionados encontrados",
              example: 5,
            },
            timestamp: {
              type: "string",
              format: "date-time",
              example: "2025-11-13T22:35:22.452Z",
            },
          },
        },
        ErrorResponse: {
          type: "object",
          properties: {
            success: {
              type: "boolean",
              example: false,
            },
            message: {
              type: "string",
              example: "Produto não encontrado",
            },
            timestamp: {
              type: "string",
              format: "date-time",
              example: "2025-11-13T22:35:22.452Z",
            },
          },
        },
      },
    },
  },
  apis: ["./src/routes/*.ts", "./src/controllers/*.ts"],
};

export const swaggerSpec = swaggerJsdoc(options);
