#!/bin/bash

# Script de build para a aplicação KNN
echo "Compilando aplicação KNN..."

# Verificar se Go está instalado
if ! command -v go &> /dev/null; then
    echo "Erro: Go não está instalado"
    exit 1
fi

# Baixar dependências
echo "Baixando dependências..."
go mod tidy

# Compilar aplicação
echo "Compilando..."
go build -o knn-process main.go

if [ $? -eq 0 ]; then
    echo "Compilação concluída com sucesso!"
    echo "Execute com: ./knn-process"
else
    echo "Erro na compilação"
    exit 1
fi
