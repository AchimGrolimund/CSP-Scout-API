package api

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/repository"
	"gofr.dev/pkg/gofr"
	"strconv"
)

type Handler struct {
	reportRepo *repository.ReportRepository
}

func NewHandler(reportRepo *repository.ReportRepository) *Handler {
	return &Handler{reportRepo: reportRepo}
}

func (h *Handler) FindOne(ctx *gofr.Context) (interface{}, error) {
	// Extract query parameters from the HTTP request
	return h.reportRepo.FindOne(ctx)
}

func (h *Handler) FindByID(ctx *gofr.Context) (interface{}, error) {
	// Extract the 'time' query parameter from the HTTP request
	id := ctx.Request.Param("id")

	return h.reportRepo.FindByID(ctx, id)
}

func (h *Handler) FindByTimeLT(ctx *gofr.Context) (interface{}, error) {
	// Extract the 'time' query parameter from the HTTP request
	timeStr := ctx.Request.Param("time")

	// Convert the 'time' query parameter to an integer
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return nil, err
	}

	return h.reportRepo.FindByTimeLT(ctx, time)
}

func (h *Handler) FindByTimeGT(ctx *gofr.Context) (interface{}, error) {
	// Extract the 'time' query parameter from the HTTP request
	timeStr := ctx.Request.Param("time")

	// Convert the 'time' query parameter to an integer
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return nil, err
	}

	return h.reportRepo.FindByTimeGT(ctx, time)
}

func (h *Handler) FindAll(ctx *gofr.Context) (interface{}, error) {
	return h.reportRepo.FindAll(ctx)
}
