package mongodb

import (
	"context"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTopIPs implements StatisticsRepository.GetTopIPs
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

	cursor, err := r.getCollection().Aggregate(ctx, pipeline)
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

// GetTopViolatedDirectives implements StatisticsRepository.GetTopViolatedDirectives
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

	cursor, err := r.getCollection().Aggregate(ctx, pipeline)
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
