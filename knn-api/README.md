# KNN API

API para recomendação de produtos usando algoritmo KNN.

## 🚀 Tecnologias

- Node.js
- Express
- TypeScript
- Drizzle ORM
- MySQL
- Swagger (OpenAPI 3.0)

## 📦 Instalação

```bash
npm install
```

## ⚙️ Configuração

1. Copie o arquivo `.env.example` para `.env`:

```bash
cp .env.example .env
```

2. Configure as variáveis de ambiente no arquivo `.env`:

```
PORT=3000
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=sua_senha
DB_NAME=knn_database
```

## 🏃 Executar

### Desenvolvimento

```bash
npm run dev
```

### Produção

```bash
npm run build
npm start
```

## 📚 Documentação

A documentação completa da API está disponível via Swagger UI:

- **Swagger UI**: `http://localhost:3000/api-docs`
- **Swagger JSON**: `http://localhost:3000/api-docs.json`

Através do Swagger você pode:

- Ver todos os endpoints disponíveis
- Testar os endpoints diretamente no navegador
- Ver exemplos de requisições e respostas
- Visualizar os schemas de dados

## 📡 Endpoints

### Obter produtos relacionados

**GET** `/api/products/:id/related`

Retorna produtos relacionados ao produto especificado, ordenados por similaridade.

**Exemplo de resposta:**

```json
{
  "success": true,
  "message": "Encontrados 5 produtos casados",
  "data": {
    "parentProductId": 11,
    "parentProductName": "Gabinete Gamer RGB",
    "relatedProducts": [
      {
        "related_id": 12,
        "order": 1,
        "similarity": 0.9791666666666666,
        "related_description": "Cadeira gamer ergonômica com suporte lombar, reclinável",
        "related_price": "1299.90"
      }
    ]
  },
  "totalFound": 5,
  "timestamp": "2025-11-13T22:35:22.452Z"
}
```

## 📁 Estrutura do Projeto

```
knn-api/
├── src/
│   ├── config/         # Configurações (database, schema, swagger)
│   ├── controllers/    # Controladores
│   ├── repositories/   # Repositórios de dados
│   ├── routes/         # Rotas da API
│   ├── services/       # Lógica de negócio
│   ├── types/          # Tipos TypeScript
│   ├── utils/          # Utilitários
│   └── index.ts        # Entry point
├── .env.example
├── package.json
├── tsconfig.json
└── README.md
```

## 🔍 Schema do Banco de Dados

### Tabela: knn_products

Armazena as relações de produtos e suas similaridades calculadas pelo algoritmo KNN.

| Campo      | Tipo     | Descrição                       |
| ---------- | -------- | ------------------------------- |
| id         | int      | Chave primária (auto increment) |
| parent_id  | int      | ID do produto pai               |
| related_id | int      | ID do produto relacionado       |
| order      | int      | Ordem de relevância             |
| similarity | double   | Índice de similaridade (0-1)    |
| created_at | datetime | Data de criação                 |
| updated_at | datetime | Data de atualização             |

### Tabela: products

Armazena os dados dos produtos.

| Campo       | Tipo         | Descrição                       |
| ----------- | ------------ | ------------------------------- |
| id          | int          | Chave primária (auto increment) |
| name        | varchar(255) | Nome do produto                 |
| description | varchar(500) | Descrição do produto            |
| category    | varchar(100) | Categoria do produto            |
| price       | varchar(20)  | Preço do produto                |
| created_at  | datetime     | Data de criação                 |
| updated_at  | datetime     | Data de atualização             |
