package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/YuriLuiz1/ninja-platform-go/internal/models"
	"github.com/YuriLuiz1/ninja-platform-go/internal/service"
	"github.com/YuriLuiz1/ninja-platform-go/pkg/security"
)

type GeninHandler struct{
	service *service.GeninService
}

func GeninsHandler(geninService *service.GeninService) *GeninHandler{
	return &GeninHandler{
		service: geninService,
	}
}

func(genin *GeninHandler) Search(w http.ResponseWriter, r*http.Request){
	pass := r.Header.Get("Authorization")

	adminPassEnv := os.Getenv("ADMIN_PASS")
		if adminPassEnv == "" {
			log.Println("Erro critico: Senha admin não configurada no ambiente")
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
			return
	}
	
	if pass != adminPassEnv {
		http.Error(w, "Senha admin incorreta ou vazia, por favor verificar", http.StatusUnauthorized)
		return
	}

	searchGenins, err := genin.service.SearchGenins()

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(searchGenins)

}

func(g *GeninHandler) Save(w http.ResponseWriter, r*http.Request){
	var genin models.Genins
	json.NewDecoder(r.Body).Decode(&genin)
	
	hashedPassword, err := security.HashPassword(genin.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	genin.Password = hashedPassword

	createdGenin, err := g.service.CreateGenins(genin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdGenin)
}

func (g *GeninHandler) SearchOne(w http.ResponseWriter, r*http.Request){
	var genin models.Genins	
	err := json.NewDecoder(r.Body).Decode(&genin)

	if err != nil{
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	searchGenin, err := g.service.SearchOneGenin(genin.Email, genin.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	searchGenin.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(searchGenin)
}