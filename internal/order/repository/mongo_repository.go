package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type mongoRepository struct {
	db *mongo.Client
}

func NewMongoRepository() {}