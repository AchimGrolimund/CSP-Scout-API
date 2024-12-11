package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoRepository(uri, database, collection string) (*MongoRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return &MongoRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *MongoRepository) Create(ctx context.Context, report domain.Report) error {
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(ctx, report)
	if err != nil {
		return fmt.Errorf("failed to insert report: %v", err)
	}

	return nil
}

func (r *MongoRepository) GetByID(ctx context.Context, id string) (*domain.Report, error) {
	collection := r.client.Database(r.database).Collection(r.collection)

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	var report domain.Report
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&report)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get report: %v", err)
	}

	return &report, nil
}

func (r *MongoRepository) List(ctx context.Context) ([]domain.Report, error) {
	collection := r.client.Database(r.database).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to list reports: %v", err)
	}
	defer cursor.Close(ctx)

	var reports []domain.Report
	if err = cursor.All(ctx, &reports); err != nil {
		return nil, fmt.Errorf("failed to decode reports: %v", err)
	}

	return reports, nil
}

func (r *MongoRepository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
