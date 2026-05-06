package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/YuriLuiz1/ninja-platform-go/internal/handler"
	"github.com/YuriLuiz1/ninja-platform-go/internal/repository"
	"github.com/YuriLuiz1/ninja-platform-go/internal/service"
	"github.com/YuriLuiz1/ninja-platform-go/pkg/security"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Carrega as variáveis do arquivo .env para o sistema
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não encontrei o arquivo .env, usando variáveis de ambiente do sistema")
	}

	// Pega a URL do banco diretamente no .env
	mongoURI := os.Getenv("DATABASE_URI")
	if mongoURI == "" {
		log.Fatal("DATABASE_URI não foi configurada!")
	}

	// Cria um limite de tempo de 10 segundos para conectar e atribui a variavel ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ctx armazena o tempo de conexão definido acima por isso ele esta aqui em baixo
	// Tenta conectar no banco
	uri := mongoURI // variavel uri armazena a uri do banco buscada la no .env no codigo mais acima
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Falha ao conectar no Banco de Dados: ", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	db := client.Database("ninja-animes")
	animesCollection := db.Collection("animes")
	geninsCollection := db.Collection("genins")

	// Instancia interna dos animes
	repo := repository.AnimesRepository(animesCollection)
	svc := service.AnimesService(repo)
	hdl := handler.AnimesHandler(svc)

	// Rotas dos animes
	mux.HandleFunc("POST /cadastrar", hdl.Create)
	mux.HandleFunc("GET /buscar", hdl.Search)
	mux.HandleFunc("GET /buscarId/{id}", hdl.SearchById)
	mux.HandleFunc("DELETE /deletar/{id}", hdl.Delete)
	mux.HandleFunc("PATCH /atualizar/{id}", hdl.UpdateById)

	// Instancia interna dos usuários
	repoGenins := repository.GeninsRepository(geninsCollection)
	svcGenins := service.GeninsService(repoGenins)
	hdlGenins := handler.GeninsHandler(svcGenins)

	// Rotas dos usuários
	mux.HandleFunc("GET /buscarGenins", hdlGenins.Search)
	mux.HandleFunc("POST /cadastrarGenin", hdlGenins.Save)


	fmt.Println("Servidor rodando na porta 8000")
	http.ListenAndServe(":8000", security.EnableCORS(mux))
}