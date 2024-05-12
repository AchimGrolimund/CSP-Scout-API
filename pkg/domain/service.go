package domain

import (
	"gofr.dev/pkg/gofr"
)

type ReportRepository interface {
	FindOne(ctx *gofr.Context) (Report, error)
	FindAll(ctx *gofr.Context) ([]Report, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(ctx *gofr.Context) (Report, error) {
	return s.repo.FindOne(ctx)
}

func (s *ReportService) GetAllReports(ctx *gofr.Context) ([]Report, error) {
	return s.repo.FindAll(ctx)
}
