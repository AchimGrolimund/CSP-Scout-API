# CSP-Scout-API

A Content Security Policy (CSP) reporting API built with Go, using Domain-Driven Design principles.

## Project Structure

The project follows a DDD (Domain-Driven Design) architecture:

```
.
├── cmd/
│   └── csp-scout-api/        # Application entry point
├── pkg/
│   ├── domain/              # Domain models and interfaces
│   ├── application/         # Application services
│   ├── infrastructure/      # Infrastructure implementations (MongoDB)
│   └── interfaces/          # HTTP handlers and routes
└── configs/                 # Configuration files
```

## Technologies Used

- Go 1.22+
- Gin Web Framework
- MongoDB Driver v1.13
- Domain-Driven Design principles

## Prerequisites

- Go 1.22 or higher
- MongoDB instance
- Environment variables configured (see Configuration section)

## Configuration

Copy the `.local.env` file in the `configs` directory and adjust the values according to your environment:

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=csp_scout
MONGODB_COLLECTION=reports
SERVER_PORT=8080
```

## Running the Application

1. Ensure MongoDB is running
2. Configure your environment variables in `configs/.local.env`
3. Run the application:

```bash
go run cmd/csp-scout-api/main.go
```

## API Endpoints

### Reports

- `POST /api/v1/reports` - Create a new CSP report
- `GET /api/v1/reports` - List all CSP reports
- `GET /api/v1/reports/:id` - Get a specific CSP report by ID

## Report Model

```go
type ReportData struct {
    DocumentUri        string
    Referrer          string
    ViolatedDirective string
    EffectiveDirective string
    OriginalPolicy    string
    Disposition       string
    BlockedUri        string
    LineNumber        int
    SourceFile        string
    StatusCode        int
    ScriptSample      string
    ClientIP          string
    UserAgent         string
    ReportTime        int
}

type Report struct {
    ID     string
    Report ReportData
}
