package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockReportsService is a mock implementation of ReportsService
type MockReportsService struct {
	mock.Mock
}

func (m *MockReportsService) CreateReport(ctx context.Context, report *domain.Report) error {
	args := m.Called(ctx, report)
	return args.Error(0)
}

func (m *MockReportsService) GetReport(ctx context.Context, id string) (*domain.Report, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Report), args.Error(1)
}

func (m *MockReportsService) ListReports(ctx context.Context) ([]domain.Report, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Report), args.Error(1)
}

func setupReportTestRouter(service *MockReportsService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")
	setupReportRoutesV1(v1, service)
	v2 := router.Group("/v2")
	setupReportRoutesV2(v2, service)
	return router
}

func TestCreateReportV1(t *testing.T) {
	testID := primitive.NewObjectID()
	
	tests := []struct {
		name           string
		setupMock      func(*MockReportsService)
		requestBody    interface{}
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			setupMock: func(m *MockReportsService) {
				m.On("CreateReport", mock.Anything, mock.AnythingOfType("*domain.Report")).Return(nil)
			},
			requestBody: domain.Report{
				ID: testID,
				Report: domain.ReportData{
					DocumentUri:        "https://example.com",
					ViolatedDirective: "script-src",
					ClientIP:          "192.168.1.1",
				},
			},
			expectedStatus: http.StatusCreated,
			expectedBody: domain.Report{
				ID: testID,
				Report: domain.ReportData{
					DocumentUri:        "https://example.com",
					ViolatedDirective: "script-src",
					ClientIP:          "192.168.1.1",
				},
			},
		},
		{
			name: "Service Error",
			setupMock: func(m *MockReportsService) {
				m.On("CreateReport", mock.Anything, mock.AnythingOfType("*domain.Report")).Return(errors.New("service error"))
			},
			requestBody: domain.Report{
				ID: testID,
				Report: domain.ReportData{
					DocumentUri: "https://example.com",
				},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"error": "service error"},
		},
		{
			name:           "Invalid Request Body",
			setupMock:      func(m *MockReportsService) {},
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"error": "json: cannot unmarshal string into Go value of type domain.Report"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockReportsService)
			tt.setupMock(mockService)
			router := setupReportTestRouter(mockService)

			body, _ := json.Marshal(tt.requestBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/reports", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if tt.expectedStatus == http.StatusCreated {
				// For successful creation, verify the structure matches
				var actualReport domain.Report
				err = json.Unmarshal(w.Body.Bytes(), &actualReport)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(domain.Report).Report.DocumentUri, actualReport.Report.DocumentUri)
				assert.Equal(t, tt.expectedBody.(domain.Report).Report.ViolatedDirective, actualReport.Report.ViolatedDirective)
				assert.Equal(t, tt.expectedBody.(domain.Report).Report.ClientIP, actualReport.Report.ClientIP)
			} else {
				expectedJSON, _ := json.Marshal(tt.expectedBody)
				actualJSON, _ := json.Marshal(response)
				assert.JSONEq(t, string(expectedJSON), string(actualJSON))
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetReportV1(t *testing.T) {
	testID := primitive.NewObjectID()
	
	tests := []struct {
		name           string
		setupMock      func(*MockReportsService)
		reportID       string
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			setupMock: func(m *MockReportsService) {
				m.On("GetReport", mock.Anything, testID.Hex()).Return(&domain.Report{
					ID: testID,
					Report: domain.ReportData{
						DocumentUri:        "https://example.com",
						ViolatedDirective: "script-src",
						ClientIP:          "192.168.1.1",
					},
				}, nil)
			},
			reportID:       testID.Hex(),
			expectedStatus: http.StatusOK,
			expectedBody: domain.Report{
				ID: testID,
				Report: domain.ReportData{
					DocumentUri:        "https://example.com",
					ViolatedDirective: "script-src",
					ClientIP:          "192.168.1.1",
				},
			},
		},
		{
			name: "Not Found",
			setupMock: func(m *MockReportsService) {
				m.On("GetReport", mock.Anything, "non-existent").Return(nil, errors.New("report not found"))
			},
			reportID:       "non-existent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   gin.H{"error": "report not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockReportsService)
			tt.setupMock(mockService)
			router := setupReportTestRouter(mockService)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/reports/"+tt.reportID, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var actualReport domain.Report
				err := json.Unmarshal(w.Body.Bytes(), &actualReport)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(domain.Report).ID, actualReport.ID)
				assert.Equal(t, tt.expectedBody.(domain.Report).Report.DocumentUri, actualReport.Report.DocumentUri)
				assert.Equal(t, tt.expectedBody.(domain.Report).Report.ViolatedDirective, actualReport.Report.ViolatedDirective)
				assert.Equal(t, tt.expectedBody.(domain.Report).Report.ClientIP, actualReport.Report.ClientIP)
			} else {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(gin.H)["error"], response["error"])
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestListReportsV1(t *testing.T) {
	testID := primitive.NewObjectID()
	
	tests := []struct {
		name           string
		setupMock      func(*MockReportsService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			setupMock: func(m *MockReportsService) {
				m.On("ListReports", mock.Anything).Return([]domain.Report{
					{
						ID: testID,
						Report: domain.ReportData{
							DocumentUri:        "https://example.com",
							ViolatedDirective: "script-src",
							ClientIP:          "192.168.1.1",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.Report{
				{
					ID: testID,
					Report: domain.ReportData{
						DocumentUri:        "https://example.com",
						ViolatedDirective: "script-src",
						ClientIP:          "192.168.1.1",
					},
				},
			},
		},
		{
			name: "Service Error",
			setupMock: func(m *MockReportsService) {
				m.On("ListReports", mock.Anything).Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"error": "service error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockReportsService)
			tt.setupMock(mockService)
			router := setupReportTestRouter(mockService)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/reports", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var actualReports []domain.Report
				err := json.Unmarshal(w.Body.Bytes(), &actualReports)
				assert.NoError(t, err)
				assert.Len(t, actualReports, len(tt.expectedBody.([]domain.Report)))
				for i, expectedReport := range tt.expectedBody.([]domain.Report) {
					assert.Equal(t, expectedReport.ID, actualReports[i].ID)
					assert.Equal(t, expectedReport.Report.DocumentUri, actualReports[i].Report.DocumentUri)
					assert.Equal(t, expectedReport.Report.ViolatedDirective, actualReports[i].Report.ViolatedDirective)
					assert.Equal(t, expectedReport.Report.ClientIP, actualReports[i].Report.ClientIP)
				}
			} else {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(gin.H)["error"], response["error"])
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestReportsV2Endpoints(t *testing.T) {
	mockService := new(MockReportsService)
	router := setupReportTestRouter(mockService)

	tests := []struct {
		name     string
		method   string
		endpoint string
	}{
		{
			name:     "CreateV2",
			method:   "POST",
			endpoint: "/v2/reports",
		},
		{
			name:     "GetV2",
			method:   "GET",
			endpoint: "/v2/reports/test-id",
		},
		{
			name:     "ListV2",
			method:   "GET",
			endpoint: "/v2/reports",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.endpoint, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusNotImplemented, w.Code)

			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, "V2 not implemented yet", response["error"])
		})
	}
}
