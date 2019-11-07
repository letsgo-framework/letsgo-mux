package database

import "go.mongodb.org/mongo-driver/mongo"

// UserCollection return s mongodb collection
func UserCollection() *mongo.Collection {
	return DB.Collection("users")
}
