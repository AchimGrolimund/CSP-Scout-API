package repository

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"github.com/vipul-rawat/gofr-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

type ReportRepository struct {
	db *mongo.Client
}

func NewReportRepository(db *mongo.Client) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) FindOne(ctx *gofr.Context) (domain.Report, error) {
	var result domain.Report

	err := ctx.Mongo.FindOne(ctx, "reports", bson.M{"report.violateddirective": "connect-src"} /* valid filter */, &result)
	if err != nil {
		return domain.Report{}, err
	}

	return result, nil
}

func (r *ReportRepository) FindAll(ctx *gofr.Context) ([]domain.Report, error) {
	var results []domain.Report

	err := ctx.Mongo.Find(ctx, "reports", bson.M{"report.violateddirective": "connect-src"} /* valid filter */, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
