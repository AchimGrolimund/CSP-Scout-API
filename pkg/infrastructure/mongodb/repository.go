package mongodb

import (
	"context"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
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

func (r *MongoRepository) CreateReport(ctx context.Context, report *domain.Report) error {
	_, err := r.client.Database(r.database).Collection(r.collection).InsertOne(ctx, report)
	return err
}

func (r *MongoRepository) GetReport(ctx context.Context, id string) (*domain.Report, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report domain.Report
	err = r.client.Database(r.database).Collection(r.collection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *MongoRepository) ListReports(ctx context.Context) ([]domain.Report, error) {
	cursor, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []domain.Report
	if err := cursor.All(ctx, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *MongoRepository) GetTopIPs(ctx context.Context) ([]application.TopIPResult, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$report.clientip"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "count", Value: -1}}}},
		{{Key: "$limit", Value: 20}},
		{{Key: "$project", Value: bson.D{
			{Key: "ip", Value: "$_id"},
			{Key: "count", Value: 1},
			{Key: "_id", Value: 0},
		}}},
	}

	cursor, err := r.client.Database(r.database).Collection(r.collection).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []application.TopIPResult
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *MongoRepository) GetTopViolatedDirectives(ctx context.Context) ([]application.TopDirectiveResult, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$report.violateddirective"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "count", Value: -1}}}},
		{{Key: "$limit", Value: 10}},
		{{Key: "$project", Value: bson.D{
			{Key: "directive", Value: "$_id"},
			{Key: "count", Value: 1},
			{Key: "_id", Value: 0},
		}}},
	}

	cursor, err := r.client.Database(r.database).Collection(r.collection).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []application.TopDirectiveResult
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *MongoRepository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
