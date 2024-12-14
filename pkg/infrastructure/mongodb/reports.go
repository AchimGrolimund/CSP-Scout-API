package mongodb

import (
	"context"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateReport implements ReportsRepository.CreateReport
func (r *MongoRepository) CreateReport(ctx context.Context, report *domain.Report) error {
	_, err := r.getCollection().InsertOne(ctx, report)
	return err
}

// GetReport implements ReportsRepository.GetReport
func (r *MongoRepository) GetReport(ctx context.Context, id string) (*domain.Report, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report domain.Report
	err = r.getCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// ListReports implements ReportsRepository.ListReports
func (r *MongoRepository) ListReports(ctx context.Context) ([]domain.Report, error) {
	cursor, err := r.getCollection().Find(ctx, bson.M{})
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
