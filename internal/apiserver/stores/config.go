package stores

import (
	"context"

	"github.com/omekov/exclusive/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

type Store struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *Store) ConfigMongo(URI string) error {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	s.collection = client.Database("exclusive").Collection("customers")
	return nil
}

func (s *Store) RegisterCustomer(customer *models.Customer) error {
	_, err := s.collection.InsertOne(ctx, customer)
	return err
}

func (s *Store) FindCustomer(username string) {
	s.collection.FindOne(ctx, username)
}
