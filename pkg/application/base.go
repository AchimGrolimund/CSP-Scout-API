package application

import (
	"context"
)

// Repository defines the complete repository interface combining all sub-repositories
type Repository interface {
	ReportsRepository
	StatisticsRepository
	Close(ctx context.Context) error
}

// Service defines the complete service interface combining all sub-services
type Service struct {
	Reports    ReportsService
	Statistics StatisticsService
}

// NewService creates a new complete service instance
func NewService(repo Repository) *Service {
	return &Service{
		Reports:    NewReportsService(repo),
		Statistics: NewStatisticsService(repo),
	}
}
