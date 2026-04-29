package repository

import (
	"context"
	"time"

	"github.com/YuriLuiz1/ninja-platform-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAnimesRepository struct {
	collection *mongo.Collection
}

func AnimesRepository(collection *mongo.Collection) *MongoAnimesRepository {
	return &MongoAnimesRepository{
		collection: collection,
	}
}

func(r *MongoAnimesRepository) Save(anime models.Animes) (models.Animes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, anime)
	if err != nil {
		return models.Animes{}, err
	}

	return anime, nil
}

func(r *MongoAnimesRepository) Search() ([]models.Animes, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var animes []models.Animes

	if err = cursor.All(ctx, &animes); err != nil {
		return nil, err
	}

	return animes, nil
}

func(r *MongoAnimesRepository) Delete(id string) (error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ObjectIDFromHex transforma o id string em hexadecimal para que o mongo entenda que isso é um id
	objID, err := primitive.ObjectIDFromHex(id) 
	if err != nil{
		return err // ID Inválido
	}
	
	//Aqui ele fala Procure esse id (objID que veio da requisição) no campo id do meu anime (_id no Mongo)
	filter := bson.M{"_id": objID} 

	//Apaga o primeiro documento que encontrar com esse ID
	_, err = r.collection.DeleteOne(ctx, filter)
	return err
}

func(r *MongoAnimesRepository) SearchUnique(id string) (models.Animes, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Animes{}, err
	}	

	var anime models.Animes

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&anime)
	if err != nil {
		return models.Animes{}, err
	}

	return anime, nil
}