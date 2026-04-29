package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YuriLuiz1/ninja-platform-go/internal/models"
	"github.com/YuriLuiz1/ninja-platform-go/internal/service"
)

type AnimeHandler struct {
	service *service.AnimeService
}

func AnimesHandler(animeService *service.AnimeService) *AnimeHandler {
	return &AnimeHandler{
		service: animeService,
	}
}

func (anime *AnimeHandler) Create(w http.ResponseWriter, r*http.Request){
	var animes models.Animes
	json.NewDecoder(r.Body).Decode(&animes)

	createdAnime, err := anime.service.CreateAnimes(animes)


	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAnime)
}

func (anime *AnimeHandler) Search(w http.ResponseWriter, r*http.Request){
	
	searchAnimes, err := anime.service.SearchAnimes()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(searchAnimes)
}

func (anime *AnimeHandler) Delete(w http.ResponseWriter, r*http.Request){

	// r.PathValue("id"): No Go 1.22+, isso pega o valor da url definida no main
	id := r.PathValue("id")

	err := anime.service.DeleteAnimes(id)
	if err != nil {
		http.Error(w, "Erro ao deletar: "+err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (animeh *AnimeHandler) SearchById(w http.ResponseWriter, r*http.Request){
	id := r.PathValue("id")

	anime, err := animeh.service.SearchAnimeForId(id)

	if err != nil {
		http.Error(w, "Anime não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anime)
}

