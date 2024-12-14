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
├── configs/                 # Configuration files
└── docker/                  # Docker configuration
```

## Technologies Used

- Go 1.22+
- Gin Web Framework
- MongoDB Driver
- Domain-Driven Design principles
- Testify (testing framework)

## Prerequisites

- Go 1.22 or higher
- MongoDB instance
- Environment variables configured (see Configuration section)

## Configuration

Copy the `.local.env` file in the `configs` directory and adjust the values according to your environment:

```env
MONGODB_URI=mongodb://localhost:27017/
MONGODB_DATABASE=csp-report
MONGODB_COLLECTION=reports
SERVER_PORT=8081
```

## Running the Application

1. Ensure MongoDB is running
2. Configure your environment variables in `configs/.local.env`
3. Run the application:

```bash
go run cmd/csp-scout-api/main.go
```

The server will start on the configured port (default: 8081).

## API Endpoints

### Reports

- `POST /api/v1/reports` - Create a new CSP report
- `GET /api/v1/reports` - List all CSP reports
- `GET /api/v1/reports/:id` - Get a specific CSP report by ID

### Statistics

- `GET /api/v1/statistics/top-ips` - Get the most frequent client IPs
- `GET /api/v1/statistics/top-directives` - Get the most violated CSP directives

## Data Models

### Report Model

```go
type ReportData struct {
    DocumentUri        string `json:"documenturi"`
    Referrer          string `json:"referrer"`
    ViolatedDirective string `json:"violateddirective"`
    EffectiveDirective string `json:"effectivedirective"`
    OriginalPolicy    string `json:"originalpolicy"`
    Disposition       string `json:"disposition"`
    BlockedUri        string `json:"blockeduri"`
    LineNumber        int    `json:"linenumber"`
    SourceFile        string `json:"sourcefile"`
    StatusCode        int    `json:"statuscode"`
    ScriptSample      string `json:"scriptsample"`
    ClientIP          string `json:"clientip"`
    UserAgent         string `json:"useragent"`
    ReportTime        int    `json:"reporttime"`
}

type Report struct {
    ID     primitive.ObjectID `json:"_id"`
    Report ReportData         `json:"report"`
}
```

### Statistics Models

```go
type TopIPResult struct {
    IP    string `json:"ip"`
    Count int    `json:"count"`
}

type TopDirectiveResult struct {
    Directive string `json:"directive"`
    Count     int    `json:"count"`
}
```

## Testing

The project includes comprehensive test suites for the API handlers. Run the tests using:

```bash
go test ./... -v
```

Test coverage includes:
- Report creation, retrieval, and listing
- Statistics endpoints
- Error handling
- Input validation
- V2 endpoint placeholders

## API Versioning

The API supports versioning through URL prefixes:
- V1: Currently implemented (`/api/v1/...`)
- V2: Placeholder endpoints for future implementation (`/api/v2/...`)

## Error Handling

The API returns appropriate HTTP status codes and JSON error messages:
- 200: Successful operation
- 201: Resource created
- 400: Bad request / Invalid input
- 404: Resource not found
- 500: Internal server error

Example error response:
```json
{
    "error": "error message here"
}
```
