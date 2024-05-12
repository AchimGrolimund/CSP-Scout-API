package api

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/repository"
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	reportRepo *repository.ReportRepository
}

func NewHandler(reportRepo *repository.ReportRepository) *Handler {
	return &Handler{reportRepo: reportRepo}
}

func (h *Handler) FindOne(ctx *gofr.Context) (interface{}, error) {
	return h.reportRepo.FindOne(ctx)
}

func (h *Handler) FindAll(ctx *gofr.Context) (interface{}, error) {
	return h.reportRepo.FindAll(ctx)
}
