package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRepository implements the application.Repository interface
type MongoRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

// NewMongoRepository creates a new MongoDB repository instance
func NewMongoRepository(uri, database, collection string) (*MongoRepository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

// Close implements the Close method required by the Repository interface
func (r *MongoRepository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

// getCollection returns the MongoDB collection
func (r *MongoRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection(r.collection)
}
