package application

import (
	"context"
)

// TopIPResult represents a client IP with its occurrence count
type TopIPResult struct {
	IP    string `json:"ip"`
	Count int    `json:"count"`
}

// TopDirectiveResult represents a violated directive with its occurrence count
type TopDirectiveResult struct {
	Directive string `json:"directive"`
	Count     int    `json:"count"`
}

// StatisticsRepository defines statistics-specific repository methods
type StatisticsRepository interface {
	GetTopIPs(ctx context.Context) ([]TopIPResult, error)
	GetTopViolatedDirectives(ctx context.Context) ([]TopDirectiveResult, error)
}

// StatisticsService defines statistics-specific service methods
type StatisticsService interface {
	GetTopIPs(ctx context.Context) ([]TopIPResult, error)
	GetTopViolatedDirectives(ctx context.Context) ([]TopDirectiveResult, error)
}

type statisticsService struct {
	repo StatisticsRepository
}

func NewStatisticsService(repo StatisticsRepository) StatisticsService {
	return &statisticsService{
		repo: repo,
	}
}

func (s *statisticsService) GetTopIPs(ctx context.Context) ([]TopIPResult, error) {
	return s.repo.GetTopIPs(ctx)
}

func (s *statisticsService) GetTopViolatedDirectives(ctx context.Context) ([]TopDirectiveResult, error) {
	return s.repo.GetTopViolatedDirectives(ctx)
}
