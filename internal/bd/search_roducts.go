package bd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"guolmal/internal/models"
	"strconv"
	"strings"
	"time"
)

// CheckPalindrome check palindrome string
func CheckPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func ConcatStages(params ...bson.D) []bson.D {
	var concatenated []bson.D
	for _, stage := range params {
		concatenated = append(concatenated, stage)
	}
	return concatenated
}

func SearchProductsByDescriptionBrand(search, kind string) ([]*models.Product, error) {
	// if kind is not "s" or "id"

	isPalindrome := CheckPalindrome(search)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var (
		products   []*models.Product
		percent    float64
		searchText = strings.TrimSpace(search)
		matchStage bson.D
	)
	if isPalindrome {
		percent = 0.5
	} else {
		percent = 1.0
	}

	db := MongoCN.Database("promotions")
	col := db.Collection("products")

	if kind == "k" {
		// id in database is an integer, not string like search
		searchK, err := strconv.Atoi(search)
		if err != nil {
			return products, err
		}

		matchStage = bson.D{
			// $or: []
			{"$match", bson.D{
				{"id", searchK},
			}},
		}
	} else {
		matchStage = bson.D{
			// $or: []
			{"$match", bson.D{
				{"$or", []bson.D{
					{
						{"description", bson.D{
							{"$regex", searchText},
							{"$options", "i"},
						}},
					},
					{
						{"brand", bson.D{
							{"$regex", searchText},
							{"$options", "i"},
						}},
					},
				}},
			}},
		}
	}

	addFieldsStage := bson.D{
		{Key: "$addFields", Value: bson.D{
			{"isPalindrome", isPalindrome},
			{"calculatedPrice", bson.D{
				{"$multiply", []interface{}{
					"$price",
					percent,
				}},
			}},
		},
		}}

	cursor, err := col.Aggregate(ctx, mongo.Pipeline(ConcatStages(matchStage, addFieldsStage)))
	if err != nil {
		return products, err
	}

	for cursor.Next(context.TODO()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return products, err
		}

		products = append(products, &product)
	}

	fmt.Printf("prodproducts: %+v\n", products)
	return products, nil
}
