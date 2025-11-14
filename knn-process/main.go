package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"knn-process/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Configurar logger para saída imediata (tempo real)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}

	// Conectar ao banco de dados
	log.Println("Conectando ao banco de dados...")
	db, err := connectDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()
	log.Printf("Conexão com banco de dados estabelecida com sucesso! Host: %s, Database: %s", 
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Criar serviço de análise KNN
	knnService := services.NewKNNAnalysisService(db)

	// Executar análise KNN
	fmt.Println("Iniciando análise KNN de produtos casados...")
	startTime := time.Now()
	
	err = knnService.AnalyzeProductAssociations()
	if err != nil {
		log.Fatal("Erro na análise KNN:", err)
	}

	duration := time.Since(startTime)
	fmt.Printf("Análise KNN concluída com sucesso em %v!\n", duration)

	// Exibir estatísticas
	stats, err := knnService.GetAnalysisStats()
	if err != nil {
		log.Printf("Erro ao obter estatísticas: %v", err)
	} else {
		fmt.Printf("Estatísticas da análise:\n")
		fmt.Printf("- Total de vendas analisadas: %v\n", stats["total_sales"])
		fmt.Printf("- Produtos únicos: %v\n", stats["unique_products"])
	}
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Testar conexão
	log.Println("Testando conexão com banco de dados...")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Ping ao banco de dados bem sucedido!")
	return db, nil
}
