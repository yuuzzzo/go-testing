package service

import (
	"errors"
	"fmt"

	"github.com/YuriLuiz1/ninja-platform-go/internal/models"
)

type AnimesRepository interface {
	Save(anime models.Animes) (models.Animes, error)
	Search() ([]models.Animes, error)
	Delete(id string) error
	SearchUnique(id string) (models.Animes, error)
}

type AnimeService struct{
	repo AnimesRepository
}

func AnimesService(repo AnimesRepository) *AnimeService {
	return &AnimeService{
		repo: repo,
	}
}

func (s *AnimeService) CreateAnimes(anime models.Animes) (models.Animes, error){

	fmt.Println(anime)

	if anime.Title == "" {
		return models.Animes{}, errors.New("A idade do ninja deve ser maior que 0")
	}

	if anime.Synopsis == "" {
		return models.Animes{}, errors.New("A Aldeia do ninja não pode ser vazia")
	}

	return s.repo.Save(anime)
}

func (s *AnimeService) SearchAnimes() ([]models.Animes, error){
	return s.repo.Search()
}

func(s *AnimeService) DeleteAnimes(id string) error {
	if id == "" {
		return errors.New("O ID não pode ser vazio")
	}

	return s.repo.Delete(id)
}

func(s *AnimeService) SearchAnimeForId(id string) (models.Animes, error){
	if id == ""{
		return models.Animes{}, errors.New("ID Inválido")
	}

	return s.repo.SearchUnique(id)
}