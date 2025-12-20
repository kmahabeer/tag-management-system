# Architectural Plan for Refactoring Tag Management System Backend to Go

## Overview

This plan outlines a comprehensive refactoring of the Python FastAPI backend to Go, maintaining full API compatibility with the existing OpenAPI specification while leveraging Go's performance and concurrency advantages. The design prioritizes minimal third-party dependencies, using Go's standard library where possible.

## 1. Project Structure

```txt
tag-management-system-go/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   ├── connection.go        # DB connection and migration
│   │   ├── queries/
│   │   │   ├── tags.go          # Tag-related queries
│   │   │   ├── entities.go      # Entity-related queries
│   │   │   └── ...              # Other domain queries
│   │   └── migrations/
│   │       └── 000001_initial.sql
│   ├── handlers/
│   │   ├── tags.go              # Tag HTTP handlers
│   │   ├── entities.go          # Entity HTTP handlers
│   │   └── ...                  # Other handlers
│   ├── middleware/
│   │   ├── cors.go              # CORS middleware
│   │   ├── logging.go           # Request logging
│   │   └── error.go             # Error handling
│   ├── models/
│   │   ├── tag.go               # Tag structs
│   │   ├── entity.go            # Entity structs
│   │   └── ...                  # Other model structs
│   └── services/
│       ├── tags.go              # Tag business logic
│       ├── entities.go          # Entity business logic
│       └── ...                  # Other services
├── pkg/
│   └── utils/
│       └── validation.go        # Shared utilities
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## 2. Dependencies

### Minimal Third-Party Libraries

- **github.com/lib/pq**: PostgreSQL driver for database/sql (essential for PostgreSQL support)
- **github.com/gorilla/mux**: Lightweight HTTP router (minimal overhead, ~1.5k lines)
- **github.com/google/uuid**: UUID generation and parsing (standard for UUID handling)

### Justification

- **database/sql + lib/pq**: Standard library interface with pure Go PostgreSQL driver. No ORM to keep dependencies minimal and maintain control over SQL.
- **gorilla/mux**: Chosen over chi for slightly better performance and feature set. Both are minimal; gorilla/mux provides path variables and middleware support.
- **google/uuid**: Widely used, minimal dependency for UUID operations.
- **No additional libraries**: Using Go's standard library for JSON (encoding/json), logging (log), HTTP (net/http), etc.

## 3. Database Layer

### Connection Management

```go
type DB struct {
    *sql.DB
}

func NewDB(dsn string) (*DB, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    // Connection pool settings
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    return &DB{db}, nil
}
```

### Query Execution

- Use prepared statements for performance and security
- Implement repository pattern with interfaces for testability
- Transaction management for complex operations (composite tagging, bulk updates)

### Example Repository Interface

```go
type TagRepository interface {
    GetByID(ctx context.Context, id uuid.UUID) (*models.Tag, error)
    List(ctx context.Context, limit, offset int) ([]*models.Tag, error)
    Create(ctx context.Context, tag *models.Tag) error
    Update(ctx context.Context, tag *models.Tag) error
    Delete(ctx context.Context, id uuid.UUID) error
}
```

## 4. Models/Structs

### Database Models

```go
type Tag struct {
    ID             uuid.UUID       `db:"id" json:"id"`
    Name           string          `db:"name" json:"name"`
    DisplayName    sql.NullString  `db:"display_name" json:"display_name"`
    Metadata       json.RawMessage `db:"metadata" json:"metadata"`
    PartOfSpeechID uuid.UUID       `db:"part_of_speech_id" json:"part_of_speech_id"`
    CreatedAt      time.Time       `db:"created_at" json:"created_at"`
    UpdatedAt      time.Time       `db:"updated_at" json:"updated_at"`
}
```

### API Models

Separate structs for API requests/responses to handle validation and transformation:

```go
type CreateTagRequest struct {
    Name           string          `json:"name" validate:"required"`
    DisplayName    *string         `json:"display_name"`
    Metadata       json.RawMessage `json:"metadata"`
    PartOfSpeechID *uuid.UUID      `json:"part_of_speech_id"`
}

type TagResponse struct {
    ID             uuid.UUID       `json:"id"`
    Name           string          `json:"name"`
    DisplayName    *string         `json:"display_name"`
    Metadata       json.RawMessage `json:"metadata"`
    PartOfSpeechID *uuid.UUID      `json:"part_of_speech_id"`
    CreatedAt      time.Time       `json:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at"`
}
```

## 5. HTTP Framework and Routing

### Router Setup

```go
func NewRouter() *mux.Router {
    r := mux.NewRouter()
    
    // API versioning
    api := r.PathPrefix("/api/v1").Subrouter()
    
    // Tag routes
    tags := api.PathPrefix("/tags").Subrouter()
    tags.HandleFunc("", handlers.ListTags).Methods("GET")
    tags.HandleFunc("", handlers.CreateTag).Methods("POST")
    tags.HandleFunc("/{id}", handlers.GetTag).Methods("GET")
    tags.HandleFunc("/{id}", handlers.UpdateTag).Methods("PATCH")
    tags.HandleFunc("/{id}", handlers.DeleteTag).Methods("DELETE")
    
    return r
}
```

### Handler Structure

```go
func CreateTag(w http.ResponseWriter, r *http.Request) {
    var req models.CreateTagRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        // Error handling
    }
    
    // Validation
    if err := validate.Struct(req); err != nil {
        // Validation error
    }
    
    // Call service
    tag, err := services.CreateTag(r.Context(), req)
    if err != nil {
        // Service error
    }
    
    // Response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tag)
}
```

## 6. Handlers/Controllers

### Structure

- Each domain (tags, entities, etc.) has its own handler file
- Handlers focus on HTTP concerns: parsing requests, validation, response formatting
- Business logic delegated to services
- Use context for request-scoped values and cancellation

### Error Handling

```go
type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details any    `json:"details,omitempty"`
}

func writeError(w http.ResponseWriter, code int, message string, details any) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(ErrorResponse{
        Code:    code,
        Message: message,
        Details: details,
    })
}
```

## 7. Services/Business Logic

### Structure

- Mirror Python service layer but adapted to Go patterns
- Use interfaces for dependency injection and testing
- Handle transactions for complex operations

### Example Service

```go
type TagService struct {
    repo repositories.TagRepository
    db   *database.DB
}

func (s *TagService) CreateTag(ctx context.Context, req *models.CreateTagRequest) (*models.Tag, error) {
    // Validation logic
    if req.Name == "" {
        return nil, errors.New("name is required")
    }
    
    // Check for duplicates
    existing, err := s.repo.GetByName(ctx, req.Name)
    if err != nil && !errors.Is(err, sql.ErrNoRows) {
        return nil, err
    }
    if existing != nil {
        return nil, errors.New("tag already exists")
    }
    
    // Create tag
    tag := &models.Tag{
        ID:             uuid.New(),
        Name:           req.Name,
        DisplayName:    req.DisplayName,
        Metadata:       req.Metadata,
        PartOfSpeechID: req.PartOfSpeechID,
        CreatedAt:      time.Now(),
        UpdatedAt:      time.Now(),
    }
    
    return s.repo.Create(ctx, tag)
}
```

### Complex Features Implementation

- **Composite Tagging**: Service methods for phrase construction, validation, and persistence
- **Versioning**: Relationship-based versioning through entity_relationships
- **Ratings**: Contextual rating services for tags, entities, and relationships

## 8. Middleware

### CORS Middleware

```go
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

### Logging Middleware

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap ResponseWriter to capture status code
        rw := &responseWriter{w, http.StatusOK}
        
        next.ServeHTTP(rw, r)
        
        log.Printf("%s %s %d %v", r.Method, r.URL.Path, rw.status, time.Since(start))
    })
}
```

### Error Handling Middleware

Centralized error handling with recovery from panics.

## 9. Configuration

### Environment-Based Configuration

```go
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
}

type ServerConfig struct {
    Port string
    Host string
}

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func LoadConfig() (*Config, error) {
    return &Config{
        Server: ServerConfig{
            Port:   getEnv("PORT", "8080"),
            Host:   getEnv("HOST", "localhost"),
        },
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnv("DB_PORT", "5432"),
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            DBName:   getEnv("DB_NAME", "app"),
            SSLMode:  getEnv("DB_SSLMODE", "disable"),
        },
    }, nil
}
```

## 10. Migration Strategy

### Phase 1: Database Preservation

- Keep existing PostgreSQL database unchanged
- No schema modifications required for API compatibility

### Phase 2: Parallel Development

- Develop Go backend alongside Python backend
- Use same database for both applications during transition
- Implement comprehensive logging and monitoring

### Phase 3: Gradual Migration

1. Deploy Go backend to staging environment
2. Run both backends with load balancer
3. Gradually shift traffic from Python to Go
4. Monitor for performance improvements and errors

### Phase 4: Full Cutover

- Complete traffic migration to Go backend
- Decommission Python backend
- Update deployment pipelines

### Data Preservation

- All existing data remains in PostgreSQL
- UUIDs and relationships preserved
- No data migration required

### API Compatibility

- Maintain exact OpenAPI spec compliance
- Same endpoint paths, request/response formats
- Same error codes and messages

## 11. Testing Strategy

### Unit Tests

- Test individual functions and methods
- Mock database interactions using interfaces
- Use testify for assertions

### Integration Tests

- Test with real PostgreSQL database
- Use testcontainers for isolated DB instances
- Test complete request/response cycles

### API Tests

- Use httptest for HTTP handler testing
- Test all endpoints against OpenAPI spec
- Load testing with realistic data volumes

### Example Test Structure

```go
func TestCreateTag(t *testing.T) {
    // Setup mock repository
    mockRepo := &mocks.TagRepository{}
    service := services.NewTagService(mockRepo)
    
    // Test cases
    tests := []struct {
        name    string
        req     *models.CreateTagRequest
        wantErr bool
    }{
        {"valid tag", &models.CreateTagRequest{Name: "test"}, false},
        {"empty name", &models.CreateTagRequest{Name: ""}, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := service.CreateTag(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateTag() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## 12. Deployment Considerations

### Containerization

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

### Production Configuration

- Use environment variables for configuration
- Implement health checks (/meta/health)
- Configure connection pooling appropriately
- Enable gzip compression for responses

### CI/CD Pipeline

- GitHub Actions for automated testing and deployment
- Multi-stage Docker builds for optimized images
- Database migrations in CI pipeline
- Integration tests against staging database

### Performance Optimizations

- Connection pooling tuned for production load
- Prepared statements for frequently executed queries
- Efficient JSON handling with encoding/json
- Goroutine pooling for concurrent requests

### Monitoring

- Structured logging with request IDs
- Metrics collection (response times, error rates)
- Health endpoints for load balancer checks
- Database connection monitoring

## Benefits of Go Implementation

1. **Performance**: Go's compiled nature and efficient garbage collection provide better performance than Python
2. **Concurrency**: Goroutines enable efficient handling of concurrent requests
3. **Memory Usage**: Lower memory footprint compared to Python
4. **Deployment**: Single binary deployment simplifies containerization
5. **Type Safety**: Compile-time type checking reduces runtime errors
6. **Standard Library**: Rich standard library reduces dependency management

## Risk Mitigation

1. **API Compatibility**: Comprehensive testing against OpenAPI spec
2. **Data Integrity**: Use transactions for complex operations
3. **Performance Regression**: Load testing before production deployment
4. **Error Handling**: Centralized error handling with proper HTTP status codes
5. **Monitoring**: Extensive logging and metrics during transition period

This plan provides a solid foundation for refactoring to Go while maintaining system reliability and API compatibility.
