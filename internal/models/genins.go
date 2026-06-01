package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Genins struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password,omitempty"`
	ResetToken string `bson:"resetToken" json:"resetToken"`
	ResetTokenExpire time.Time `bson:"resetTokenExpire" json:"resetTokenExpire"`
}