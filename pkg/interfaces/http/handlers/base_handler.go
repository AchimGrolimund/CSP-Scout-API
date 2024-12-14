package handlers

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/gin-gonic/gin"
)

// APIVersion represents the API version
type APIVersion string

const (
	V1 APIVersion = "v1"
	V2 APIVersion = "v2"
)

// RegisterRoutes configures all API routes with versioning
func RegisterRoutes(router *gin.Engine, service *application.Service) {
	// Register V1 routes
	apiV1 := router.Group("/api/v1")
	setupV1Routes(apiV1, service)

	// Register V2 routes when needed
	// apiV2 := router.Group("/api/v2")
	// setupV2Routes(apiV2, service)
}

// setupV1Routes configures all V1 API routes
func setupV1Routes(router *gin.RouterGroup, service *application.Service) {
	// Reports CRUD routes
	setupReportRoutesV1(router, service.Reports)

	// Statistics routes
	setupStatisticsRoutesV1(router, service.Statistics)
}

// setupV2Routes configures all V2 API routes
// Uncomment and implement when V2 is needed
/*
func setupV2Routes(router *gin.RouterGroup, service *application.Service) {
	// Reports CRUD routes
	setupReportRoutesV2(router, service.Reports)

	// Statistics routes
	setupStatisticsRoutesV2(router, service.Statistics)
}
*/
