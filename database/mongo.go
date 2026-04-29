package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não encontrei o arquivo .env, carregando variáveis do sistema!")
	}

	mongoUri := os.Getenv("DATABASE_URL")
	if mongoUri == ""{
		log.Fatal("DATABASE_URL não encontrada no .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := mongoUri

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	return client, nil
}