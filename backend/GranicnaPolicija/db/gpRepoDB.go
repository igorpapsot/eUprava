package db

import (
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GPRepo struct {
	logger *log.Logger
	client *mongo.Client
}
