# Fluxograma - Sistema de Recomendação de Produtos (KNN)

## Como Funciona o Sistema de Recomendação

Este fluxograma explica de forma simples como o sistema identifica produtos que costumam ser comprados juntos, como quando você vai ao supermercado e o caixa sugere "clientes que compraram isso também compraram aquilo".

---

## Fluxograma Principal

```mermaid
flowchart TD
    A[🚀 INÍCIO: Sistema de Recomendação] --> B[📊 Conectar ao Banco de Dados]
    B --> C{✅ Conexão OK?}
    C -->|❌ Não| D[⚠️ Erro: Não foi possível conectar]
    C -->|✅ Sim| E[🔍 Buscar Vendas dos Últimos 3 Meses]

    E --> F{📦 Existem Vendas?}
    F -->|❌ Não| G[ℹ️ Nenhuma venda encontrada<br/>Processo encerrado]
    F -->|✅ Sim| H[📋 Construir Grupos de Produtos<br/>Quais produtos foram comprados juntos?]

    H --> I[🧹 Limpar Recomendações Antigas]
    I --> J[📚 Buscar Todos os Produtos Cadastrados]
    J --> K[🔄 Para Cada Produto]

    K --> L{🛒 Produto tem vendas<br/>nos últimos 3 meses?}
    L -->|❌ Não| M[⏭️ Pular este produto]
    L -->|✅ Sim| N[🔎 Encontrar os 5 Produtos Mais Similares<br/>Usando algoritmo KNN]

    N --> O[💾 Salvar Recomendações no Banco]
    O --> P{🔄 Mais produtos<br/>para analisar?}
    P -->|✅ Sim| K
    P -->|❌ Não| Q[📊 Exibir Estatísticas]

    M --> P
    Q --> R[✅ FIM: Análise Concluída!]
    D --> S[🛑 ERRO: Processo Interrompido]
    G --> R

    style A fill:#e1f5ff
    style R fill:#d4edda
    style S fill:#f8d7da
    style D fill:#f8d7da
    style N fill:#fff3cd
    style Q fill:#d1ecf1
```

---

## Exemplo Prático: Como Funciona na Prática

### Cenário: Você compra um Notebook

```mermaid
flowchart LR
    A[👤 Cliente compra<br/>📱 Notebook] --> B[🛒 Sistema registra:<br/>Notebook + Mouse + Teclado<br/>no mesmo pedido]
    B --> C[📊 Após 3 meses,<br/>sistema analisa padrões]
    C --> D[🔍 Algoritmo KNN identifica:<br/>Quem compra Notebook<br/>também compra Mouse]
    D --> E[💡 Recomendação criada:<br/>Notebook → Mouse<br/>Similaridade: 85%]
    E --> F[🛍️ Próximo cliente que<br/>ver Notebook verá:<br/>"Clientes também compraram: Mouse"]

    style A fill:#e1f5ff
    style E fill:#fff3cd
    style F fill:#d4edda
```

---

## Detalhamento do Algoritmo KNN

```mermaid
flowchart TD
    A[🎯 Produto Alvo:<br/>Notebook] --> B[📦 Buscar todos os pedidos<br/>que contêm Notebook]
    B --> C[📋 Listar produtos que<br/>aparecem junto com Notebook<br/>em cada pedido]
    C --> D[🔢 Calcular Similaridade<br/>Distância de Jaccard]
    D --> E[📊 Comparar com TODOS<br/>os outros produtos]
    E --> F[🏆 Selecionar os 5 produtos<br/>com maior similaridade]
    F --> G[💾 Salvar no banco:<br/>Notebook → Mouse 85%<br/>Notebook → Teclado 80%<br/>Notebook → Mousepad 75%<br/>Notebook → Carregador 70%<br/>Notebook → Fone 65%]

    style A fill:#e1f5ff
    style D fill:#fff3cd
    style F fill:#d4edda
    style G fill:#d1ecf1
```

---

## Fluxo de Dados Simplificado

```mermaid
flowchart LR
    A[💾 Banco de Dados<br/>Vendas] --> B[🔍 Análise KNN]
    B --> C[📊 Produtos Similares]
    C --> D[💾 Banco de Dados<br/>Recomendações]
    D --> E[🛍️ Sistema de Vendas<br/>Mostra Sugestões]

    style A fill:#e1f5ff
    style B fill:#fff3cd
    style C fill:#d4edda
    style D fill:#d1ecf1
    style E fill:#f8d7da
```

---

## Analogia: Como um Vendedor Experiente

Imagine que você tem um vendedor muito experiente que:

1. **Observa** todas as compras dos últimos 3 meses 👀
2. **Identifica padrões**: "Quem compra X também compra Y" 📝
3. **Aprende** quais produtos combinam bem juntos 🧠
4. **Sugere** produtos relacionados quando você escolhe algo 🗣️
5. **Atualiza** suas sugestões a cada 3 meses para manter relevância 🔄

O sistema KNN faz exatamente isso, mas de forma automática e precisa!

---

## Resultado Final

Após a análise, o sistema terá:

- ✅ **Recomendações salvas** para cada produto
- ✅ **Top 5 produtos similares** para cada item
- ✅ **Percentual de similaridade** entre produtos
- ✅ **Dados atualizados** a cada execução

**Exemplo de resultado:**

```
Produto: Notebook
├── Recomendação 1: Mouse (85% similaridade)
├── Recomendação 2: Teclado (80% similaridade)
├── Recomendação 3: Mousepad (75% similaridade)
├── Recomendação 4: Carregador (70% similaridade)
└── Recomendação 5: Fone de Ouvido (65% similaridade)
```

---

## Estatísticas Exibidas

Ao final do processo, o sistema mostra:

- 📊 Total de vendas analisadas
- 🏷️ Quantidade de produtos únicos encontrados
- ⏱️ Tempo total de processamento

---

_Este fluxograma foi criado para facilitar o entendimento do sistema de recomendação baseado em KNN (K-Nearest Neighbors)._
