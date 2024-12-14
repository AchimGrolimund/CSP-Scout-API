package handlers

import (
	"net/http"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service *application.ReportService
}

func NewReportHandler(service *application.ReportService) *ReportHandler {
	return &ReportHandler{
		service: service,
	}
}

func SetupRoutes(router *gin.Engine, handler *ReportHandler) {
	api := router.Group("/api")
	{
		api.POST("/reports", handler.CreateReport)
		api.GET("/reports", handler.ListReports)
		api.GET("/reports/:id", handler.GetReport)
		api.GET("/reports/top-ips", handler.GetTopIPs)
		api.GET("/reports/top-directives", handler.GetTopViolatedDirectives)
	}
}

func (h *ReportHandler) CreateReport(c *gin.Context) {
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

func (h *ReportHandler) GetReport(c *gin.Context) {
	id := c.Param("id")
	report, err := h.service.GetReport(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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

func (h *ReportHandler) GetTopIPs(c *gin.Context) {
	topIPs, err := h.service.GetTopIPs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topIPs)
}

func (h *ReportHandler) GetTopViolatedDirectives(c *gin.Context) {
	topDirectives, err := h.service.GetTopViolatedDirectives(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topDirectives)
}
