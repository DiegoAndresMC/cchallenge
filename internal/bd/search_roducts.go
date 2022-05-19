package bd

import (
	"cchallenge/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func SearchProductsByDescriptionBrand(search string, kind int) ([]*models.Product, error) {
	// 2 for search by text
	// 1 for search by id
	var isPalindrome bool = false

	if kind == 2 {
		isPalindrome = CheckPalindrome(search)
	}
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

	fmt.Printf("Finding products with search text: %s, %T. with kind %s\n", searchText, searchText, kind)

	db := MongoCN.Database("promotions")
	col := db.Collection("products")

	if kind == 1 {
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
	} else if kind == 2 {
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
	} else {
		return products, fmt.Errorf("kind must be 1 or 2")
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

	fmt.Printf("matchStage: %v\n", matchStage)

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
