package repository

import (
	"context"
	"time"

	"github.com/YuriLuiz1/ninja-platform-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoGeninsRepository struct {
	collection *mongo.Collection
}

func GeninsRepository(collection *mongo.Collection) *MongoGeninsRepository{
	return &MongoGeninsRepository{
		collection: collection,
	}
}

func (r *MongoGeninsRepository) Search() ([]models.Genins, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var genins []models.Genins

	if err = cursor.All(ctx, &genins); err != nil{
		return nil, err
	}

	return genins, nil
}

func (r *MongoGeninsRepository) Save(genin models.Genins) (models.Genins, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, genin)
	if err != nil {
		return models.Genins{}, err
	}

	return genin, nil
}