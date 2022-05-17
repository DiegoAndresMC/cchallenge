package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product _id, id, brand, description, image, price
type Product struct {
	MongoId    primitive.ObjectID `json:"_id,omitempty"bson:"_id,omitempty"`
	Id         int                `json:"id,omitempty" bson:"id"`
	Brand      string             `json:"brand,omitempty" bson:"brand"`
	Desc       string             `json:"description,omitempty" bson:"description"`
	Image      string             `json:"image,omitempty" bson:"image"`
	Price      int                `json:"price,omitempty" bson:"price"`
	CPrice     float64            `json:"calculatedPrice,omitempty" bson:"calculatedPrice"`
	Palindrome bool               `json:"isPalindrome,omitempty" bson:"isPalindrome"`
	Message    string             `json:"message"`
}
