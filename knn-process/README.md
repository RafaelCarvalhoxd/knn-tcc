# Sistema de Análise KNN para Produtos Casados

Este projeto implementa um sistema de análise de produtos casados usando o algoritmo KNN (K-Nearest Neighbors) com Distância de Jaccard para identificar produtos que são frequentemente vendidos juntos.

## Funcionalidades

- **Análise de Vendas**: Processa vendas dos últimos 3 meses
- **Algoritmo KNN**: Implementa KNN com Distância de Jaccard
- **Produtos Casados**: Identifica os 5 produtos mais similares para cada produto
- **Armazenamento**: Salva resultados em tabela `knn_products`

## Estrutura do Projeto

```
knn-process/
├── algorithms/           # Implementação do algoritmo KNN
│   └── knn.go
├── models/              # Modelos de dados
│   └── models.go
├── repositories/        # Camada de acesso a dados
│   ├── sale_repository.go
│   ├── product_repository.go
│   └── knn_product_repository.go
├── services/           # Lógica de negócio
│   └── knn_analysis_service.go
├── database/           # Scripts de banco de dados
│   ├── schema.sql
│   └── cleanup.sql
├── main.go            # Ponto de entrada da aplicação
├── go.mod             # Dependências do Go
└── config.env.example # Exemplo de configuração
```

## Pré-requisitos

- Go 1.21 ou superior
- MySQL 5.7 ou superior
- Git

## Instalação

1. **Clone o repositório**:

```bash
git clone <url-do-repositorio>
cd knn-process
```

2. **Instale as dependências**:

```bash
go mod tidy
```

3. **Configure o banco de dados**:

```bash
# Copie o arquivo de exemplo
cp config.env.example .env

# Edite o arquivo .env com suas configurações
nano .env
```

4. **Execute o script de criação do banco**:

```bash
mysql -u root -p < database/schema.sql
```

## Configuração

Crie um arquivo `.env` baseado no `config.env.example`:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=sua_senha
DB_NAME=knn_analysis
```

## Como Usar

1. **Execute a aplicação**:

```bash
go run main.go
```

2. **A aplicação irá**:
   - Conectar ao banco de dados
   - Buscar vendas dos últimos 3 meses
   - Aplicar o algoritmo KNN com Distância de Jaccard
   - Salvar os resultados na tabela `knn_products`

## Algoritmo KNN com Distância de Jaccard

### Como Funciona

1. **Coleta de Dados**: Busca vendas dos últimos 3 meses
2. **Construção de Conjuntos**: Para cada produto, cria um conjunto de produtos que aparecem junto nas vendas
3. **Cálculo de Similaridade**: Usa a Distância de Jaccard para calcular similaridade entre produtos
4. **Seleção KNN**: Para cada produto, encontra os 5 produtos mais similares
5. **Armazenamento**: Salva os resultados ordenados por similaridade

### Fórmula da Distância de Jaccard

```
Distância Jaccard = 1 - (|A ∩ B| / |A ∪ B|)
```

Onde:

- A = conjunto de produtos vendidos junto com produto X
- B = conjunto de produtos vendidos junto com produto Y
- |A ∩ B| = interseção dos conjuntos
- |A ∪ B| = união dos conjuntos

## Estrutura do Banco de Dados

### Tabela `products`

- `id`: ID único do produto
- `name`: Nome do produto
- `description`: Descrição
- `category`: Categoria
- `price`: Preço

### Tabela `sales`

- `id`: ID único da venda
- `sale_date`: Data da venda
- `product_id`: ID do produto
- `quantity`: Quantidade
- `price`: Preço
- `order_id`: ID do pedido

### Tabela `knn_products`

- `id`: ID único do resultado
- `parent_id`: ID do produto pai
- `related_id`: ID do produto relacionado
- `order`: Ordem de similaridade (1-5)
- `similarity`: Valor de similaridade (0-1)
- `created_at`: Data de criação
- `updated_at`: Data de atualização

## Exemplo de Uso

Após executar a aplicação, você pode consultar os resultados:

```sql
-- Buscar produtos relacionados ao produto ID 1
SELECT
    p.name as produto_pai,
    pr.name as produto_relacionado,
    kp.order as ordem,
    kp.similarity as similaridade
FROM knn_products kp
JOIN products p ON kp.parent_id = p.id
JOIN products pr ON kp.related_id = pr.id
WHERE kp.parent_id = 1
ORDER BY kp.order;
```

## Logs e Monitoramento

A aplicação fornece logs detalhados durante a execução:

- Número de vendas processadas
- Produtos analisados
- Tempo de execução
- Estatísticas finais

## Limpeza de Dados

Para remover dados de teste:

```bash
mysql -u root -p < database/cleanup.sql
```

## Desenvolvimento

### Adicionando Novos Recursos

1. **Modelos**: Adicione novos modelos em `models/`
2. **Repositórios**: Implemente acesso a dados em `repositories/`
3. **Serviços**: Adicione lógica de negócio em `services/`
4. **Algoritmos**: Implemente novos algoritmos em `algorithms/`

### Testes

Para executar testes (quando implementados):

```bash
go test ./...
```

## Troubleshooting

### Problemas Comuns

1. **Erro de Conexão**: Verifique as configurações do banco no arquivo `.env`
2. **Tabelas Não Existem**: Execute o script `database/schema.sql`
3. **Dados Insuficientes**: Certifique-se de ter vendas dos últimos 3 meses

### Logs de Debug

Para mais informações de debug, adicione logs adicionais nos serviços.

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para detalhes.
