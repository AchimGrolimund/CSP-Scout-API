package application

import (
	"context"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
)

type ReportRepository interface {
	Create(ctx context.Context, report domain.Report) error
	GetByID(ctx context.Context, id string) (*domain.Report, error)
	List(ctx context.Context) ([]domain.Report, error)
	Close(ctx context.Context) error
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{
		repo: repo,
	}
}

func (s *ReportService) CreateReport(ctx context.Context, report domain.Report) error {
	return s.repo.Create(ctx, report)
}

func (s *ReportService) GetReport(ctx context.Context, id string) (*domain.Report, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ReportService) ListReports(ctx context.Context) ([]domain.Report, error) {
	return s.repo.List(ctx)
}

func (s *ReportService) Close(ctx context.Context) error {
	return s.repo.Close(ctx)
}
