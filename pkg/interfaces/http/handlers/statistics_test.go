package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStatisticsService is a mock implementation of StatisticsService
type MockStatisticsService struct {
	mock.Mock
}

func (m *MockStatisticsService) GetTopIPs(ctx context.Context) ([]application.TopIPResult, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]application.TopIPResult), args.Error(1)
}

func (m *MockStatisticsService) GetTopViolatedDirectives(ctx context.Context) ([]application.TopDirectiveResult, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]application.TopDirectiveResult), args.Error(1)
}

func setupTestRouter(service application.StatisticsService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")
	setupStatisticsRoutesV1(v1, service)
	v2 := router.Group("/v2")
	setupStatisticsRoutesV2(v2, service)
	return router
}

func TestGetTopIPsV1(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockStatisticsService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			setupMock: func(m *MockStatisticsService) {
				m.On("GetTopIPs", mock.Anything).Return([]application.TopIPResult{
					{IP: "192.168.1.1", Count: 10},
					{IP: "192.168.1.2", Count: 5},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []application.TopIPResult{
				{IP: "192.168.1.1", Count: 10},
				{IP: "192.168.1.2", Count: 5},
			},
		},
		{
			name: "Service Error",
			setupMock: func(m *MockStatisticsService) {
				m.On("GetTopIPs", mock.Anything).Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"error": "service error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockStatisticsService)
			tt.setupMock(mockService)
			router := setupTestRouter(mockService)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/statistics/top-ips", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			expectedJSON, _ := json.Marshal(tt.expectedBody)
			actualJSON, _ := json.Marshal(response)
			assert.JSONEq(t, string(expectedJSON), string(actualJSON))

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetTopViolatedDirectivesV1(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockStatisticsService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			setupMock: func(m *MockStatisticsService) {
				m.On("GetTopViolatedDirectives", mock.Anything).Return([]application.TopDirectiveResult{
					{Directive: "script-src", Count: 15},
					{Directive: "style-src", Count: 8},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []application.TopDirectiveResult{
				{Directive: "script-src", Count: 15},
				{Directive: "style-src", Count: 8},
			},
		},
		{
			name: "Service Error",
			setupMock: func(m *MockStatisticsService) {
				m.On("GetTopViolatedDirectives", mock.Anything).Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"error": "service error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockStatisticsService)
			tt.setupMock(mockService)
			router := setupTestRouter(mockService)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/statistics/top-directives", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			expectedJSON, _ := json.Marshal(tt.expectedBody)
			actualJSON, _ := json.Marshal(response)
			assert.JSONEq(t, string(expectedJSON), string(actualJSON))

			mockService.AssertExpectations(t)
		})
	}
}

func TestV2Endpoints(t *testing.T) {
	mockService := new(MockStatisticsService)
	router := setupTestRouter(mockService)

	tests := []struct {
		name     string
		endpoint string
	}{
		{
			name:     "GetTopIPsV2",
			endpoint: "/v2/statistics/top-ips",
		},
		{
			name:     "GetTopViolatedDirectivesV2",
			endpoint: "/v2/statistics/top-directives",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.endpoint, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusNotImplemented, w.Code)

			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, "V2 not implemented yet", response["error"])
		})
	}
}
