package handler

import (
	"encoding/json"
	"net/http"

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