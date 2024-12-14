package application

import (
	"context"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
)

type ReportRepository interface {
	CreateReport(ctx context.Context, report *domain.Report) error
	GetReport(ctx context.Context, id string) (*domain.Report, error)
	ListReports(ctx context.Context) ([]domain.Report, error)
	GetTopIPs(ctx context.Context) ([]TopIPResult, error)
	GetTopViolatedDirectives(ctx context.Context) ([]TopDirectiveResult, error)
	Close(ctx context.Context) error
}

type TopIPResult struct {
	IP    string `json:"ip"`
	Count int    `json:"count"`
}

type TopDirectiveResult struct {
	Directive string `json:"directive"`
	Count     int    `json:"count"`
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{
		repo: repo,
	}
}

func (s *ReportService) CreateReport(ctx context.Context, report *domain.Report) error {
	return s.repo.CreateReport(ctx, report)
}

func (s *ReportService) GetReport(ctx context.Context, id string) (*domain.Report, error) {
	return s.repo.GetReport(ctx, id)
}

func (s *ReportService) ListReports(ctx context.Context) ([]domain.Report, error) {
	return s.repo.ListReports(ctx)
}

func (s *ReportService) GetTopIPs(ctx context.Context) ([]TopIPResult, error) {
	return s.repo.GetTopIPs(ctx)
}

func (s *ReportService) GetTopViolatedDirectives(ctx context.Context) ([]TopDirectiveResult, error) {
	return s.repo.GetTopViolatedDirectives(ctx)
}
