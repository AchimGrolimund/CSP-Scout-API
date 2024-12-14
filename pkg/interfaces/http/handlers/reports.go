package handlers

import (
	"net/http"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"github.com/gin-gonic/gin"
)

type ReportsHandler struct {
	service application.ReportsService
}

func NewReportsHandler(service application.ReportsService) *ReportsHandler {
	return &ReportsHandler{
		service: service,
	}
}

// V1 Routes
func setupReportRoutesV1(router *gin.RouterGroup, service application.ReportsService) {
	handler := NewReportsHandler(service)
	reports := router.Group("/reports")
	{
		reports.POST("", handler.CreateV1)
		reports.GET("", handler.ListV1)
		reports.GET("/:id", handler.GetV1)
	}
}

// V2 Routes (for future implementation)
func setupReportRoutesV2(router *gin.RouterGroup, service application.ReportsService) {
	handler := NewReportsHandler(service)
	reports := router.Group("/reports")
	{
		reports.POST("", handler.CreateV2)
		reports.GET("", handler.ListV2)
		reports.GET("/:id", handler.GetV2)
	}
}

// V1 Handlers
func (h *ReportsHandler) CreateV1(c *gin.Context) {
	var report domain.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateReport(c.Request.Context(), &report); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, report)
}

func (h *ReportsHandler) GetV1(c *gin.Context) {
	id := c.Param("id")
	report, err := h.service.GetReport(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportsHandler) ListV1(c *gin.Context) {
	reports, err := h.service.ListReports(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

// V2 Handlers (for future implementation)
func (h *ReportsHandler) CreateV2(c *gin.Context) {
	// Implement V2 create logic when needed
	c.JSON(http.StatusNotImplemented, gin.H{"error": "V2 not implemented yet"})
}

func (h *ReportsHandler) GetV2(c *gin.Context) {
	// Implement V2 get logic when needed
	c.JSON(http.StatusNotImplemented, gin.H{"error": "V2 not implemented yet"})
}

func (h *ReportsHandler) ListV2(c *gin.Context) {
	// Implement V2 list logic when needed
	c.JSON(http.StatusNotImplemented, gin.H{"error": "V2 not implemented yet"})
}
