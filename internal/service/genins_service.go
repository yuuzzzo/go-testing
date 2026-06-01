package service

import (
	"errors"

	"github.com/YuriLuiz1/ninja-platform-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type GeninsRepository interface {
	Search() ([]models.Genins, error)
	Save(genin models.Genins) (models.Genins, error)
	SearchOne(email string) (*models.Genins, error)
}

type GeninService struct{
	repo GeninsRepository
}

func GeninsService(repo GeninsRepository) *GeninService{
	return &GeninService{
		repo: repo,
	}
}

func (s *GeninService) SearchGenins() ([]models.Genins, error){
	return s.repo.Search()
}

func (s *GeninService) CreateGenins(genin models.Genins) (models.Genins, error) {
		if genin.Name == "" {
			return models.Genins{}, errors.New("O nome do Genin não pode ser vazio")
		}

		if genin.Email == ""{
			return models.Genins{}, errors.New("O email do Genin não pode ser vazio")
		}

		if genin.Password == "" {
			return models.Genins{}, errors.New("A senha do Genin não pode ser vazia")
		}

		return s.repo.Save(genin)
}

func (s *GeninService) SearchOneGenin(email string, password string) (*models.Genins, error){

	geninData, err := s.repo.SearchOne(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(geninData.Password), []byte(password))

	if err != nil {
		return nil, errors.New("Password is invalid")
	}

	return geninData, nil
}