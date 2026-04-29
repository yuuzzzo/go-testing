package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Animes struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string `bson:"title" json:"title"`
	Average string `bson:"average" json:"average"`
	Synopsis string	`bson:"synopsis" json:"synopsis"`
	CapeImage string `bson:"capeImage" json:"capeImage"`
	OpinionNinja string `bson:"opinionNinja" json:"opinionNinja"`
	CategoryId int `bson:"categoryId" json:"categoryId"`
	DurationEp int	`bson:"durationEp" json:"durationEp"`
	StatusFinished bool `bson:"statusFinished" json:"statusFinished"`
	StreamingPlatforms []string `bson:"streamingPlatforms" json:"streamingPlatforms"`
	Studios []string `bson:"studios" json:"studios"`
	Temp int `bson:"temp" json:"temp"`
	Episodes int `bson:"episodes" json:"episodes"`
	ReleaseDate int `bson:"releaseDate" json:"releaseDate"`
	SubCategories []string `bson:"subCategories" json:"subCategories"`
}