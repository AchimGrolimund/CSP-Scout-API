package handlers

import (
	"net/http"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	service application.StatisticsService
}

func NewStatisticsHandler(service application.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{
		service: service,
	}
}

// V1 Routes
func setupStatisticsRoutesV1(router *gin.RouterGroup, service application.StatisticsService) {
	handler := NewStatisticsHandler(service)
	stats := router.Group("/statistics")
	{
		stats.GET("/top-ips", handler.GetTopIPsV1)
		stats.GET("/top-directives", handler.GetTopViolatedDirectivesV1)
	}
}

// V2 Routes (for future implementation)
func setupStatisticsRoutesV2(router *gin.RouterGroup, service application.StatisticsService) {
	handler := NewStatisticsHandler(service)
	stats := router.Group("/statistics")
	{
		stats.GET("/top-ips", handler.GetTopIPsV2)
		stats.GET("/top-directives", handler.GetTopViolatedDirectivesV2)
	}
}

// V1 Handlers
func (h *StatisticsHandler) GetTopIPsV1(c *gin.Context) {
	topIPs, err := h.service.GetTopIPs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topIPs)
}

func (h *StatisticsHandler) GetTopViolatedDirectivesV1(c *gin.Context) {
	topDirectives, err := h.service.GetTopViolatedDirectives(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topDirectives)
}

// V2 Handlers (for future implementation)
func (h *StatisticsHandler) GetTopIPsV2(c *gin.Context) {
	// Implement V2 top IPs logic when needed
	c.JSON(http.StatusNotImplemented, gin.H{"error": "V2 not implemented yet"})
}

func (h *StatisticsHandler) GetTopViolatedDirectivesV2(c *gin.Context) {
	// Implement V2 top directives logic when needed
	c.JSON(http.StatusNotImplemented, gin.H{"error": "V2 not implemented yet"})
}
