package application

import (
	"context"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
)

// ReportsRepository defines reports-specific repository methods
type ReportsRepository interface {
	CreateReport(ctx context.Context, report *domain.Report) error
	GetReport(ctx context.Context, id string) (*domain.Report, error)
	ListReports(ctx context.Context) ([]domain.Report, error)
}

// ReportsService defines reports-specific service methods
type ReportsService interface {
	CreateReport(ctx context.Context, report *domain.Report) error
	GetReport(ctx context.Context, id string) (*domain.Report, error)
	ListReports(ctx context.Context) ([]domain.Report, error)
}

type reportsService struct {
	repo ReportsRepository
}

func NewReportsService(repo ReportsRepository) ReportsService {
	return &reportsService{
		repo: repo,
	}
}

func (s *reportsService) CreateReport(ctx context.Context, report *domain.Report) error {
	return s.repo.CreateReport(ctx, report)
}

func (s *reportsService) GetReport(ctx context.Context, id string) (*domain.Report, error) {
	return s.repo.GetReport(ctx, id)
}

func (s *reportsService) ListReports(ctx context.Context) ([]domain.Report, error) {
	return s.repo.ListReports(ctx)
}
