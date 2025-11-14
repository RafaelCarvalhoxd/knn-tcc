@echo off
REM Script de build para Windows

echo Compilando aplicacao KNN...

REM Verificar se Go esta instalado
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo Erro: Go nao esta instalado
    exit /b 1
)

REM Baixar dependencias
echo Baixando dependencias...
go mod tidy

REM Compilar aplicacao
echo Compilando...
go build -o knn-process.exe main.go

if %errorlevel% equ 0 (
    echo Compilacao concluida com sucesso!
    echo Execute com: knn-process.exe
) else (
    echo Erro na compilacao
    exit /b 1
)
