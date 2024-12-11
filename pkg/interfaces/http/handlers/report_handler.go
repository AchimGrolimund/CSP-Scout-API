package handlers

import (
	"net/http"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportHandler struct {
	service *application.ReportService
}

func NewReportHandler(service *application.ReportService) *ReportHandler {
	return &ReportHandler{
		service: service,
	}
}

func (h *ReportHandler) CreateReport(c *gin.Context) {
	var report domain.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new ObjectID if not provided
	if report.ID.IsZero() {
		report.ID = primitive.NewObjectID()
	}

	err := h.service.CreateReport(c.Request.Context(), report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, report)
}

func (h *ReportHandler) GetReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// Validate ID format
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	report, err := h.service.GetReport(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if report == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "report not found"})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) ListReports(c *gin.Context) {
	reports, err := h.service.ListReports(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

func SetupRoutes(router *gin.Engine, handler *ReportHandler) {
	api := router.Group("/api/v1")
	{
		api.POST("/reports", handler.CreateReport)
		api.GET("/reports/:id", handler.GetReport)
		api.GET("/reports", handler.ListReports)
	}
}
