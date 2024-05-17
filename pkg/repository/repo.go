package repository

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"github.com/vipul-rawat/gofr-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofr.dev/pkg/gofr"
	"log"
)

type ReportRepository struct {
	db *mongo.Client
}

func NewReportRepository(db *mongo.Client) *ReportRepository {
	return &ReportRepository{db: db}
}

// FindOne returns a single report
func (r *ReportRepository) FindOne(ctx *gofr.Context) (domain.Report, error) {
	var result domain.Report

	err := ctx.Mongo.FindOne(ctx, "reports", bson.M{"report.violateddirective": "connect-src"} /* valid filter */, &result)
	if err != nil {
		return domain.Report{}, err
	}

	return result, nil
}

// FindByID returns a single report by ID
func (r *ReportRepository) FindByID(ctx *gofr.Context, id string) (domain.Report, error) {
	var result domain.Report
	// Convert the string id to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	err = ctx.Mongo.FindOne(ctx, "reports", bson.M{"_id": objectId} /* valid filter */, &result)
	if err != nil {
		return domain.Report{}, err
	}

	return result, nil
}

// FindByTimeLT returns all reports with time less than the given time
func (r *ReportRepository) FindByTimeLT(ctx *gofr.Context, time int) ([]domain.Report, error) {
	var results []domain.Report

	filter := bson.M{"report.reporttime": bson.M{"$lt": time}}

	err := ctx.Mongo.Find(ctx, "reports", filter, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// FindByTimeGT returns all reports with time greater than the given time
func (r *ReportRepository) FindByTimeGT(ctx *gofr.Context, time int) ([]domain.Report, error) {
	var results []domain.Report

	filter := bson.M{"report.reporttime": bson.M{"$gt": time}}

	err := ctx.Mongo.Find(ctx, "reports", filter, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *ReportRepository) FindByUserAgent(ctx *gofr.Context, user_agent string) ([]domain.Report, error) {
	var results []domain.Report

	filter := bson.M{"report.useragent": user_agent}

	err := ctx.Mongo.Find(ctx, "reports", filter, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}	

// FindAll returns all reports
func (r *ReportRepository) FindAll(ctx *gofr.Context) ([]domain.Report, error) {
	var results []domain.Report

	err := ctx.Mongo.Find(ctx, "reports", bson.M{} /* valid filter */, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
