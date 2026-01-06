# Error Handling Design

## Overview

This document outlines the design for centralized error handling in the Tag Management System backend. The goal is to provide consistent, informative error responses across all API endpoints while maintaining proper logging and monitoring capabilities.

## Current State

### Existing Infrastructure

- **ErrorResponse struct**: Defined in `middleware/error.go` with fields for code, message, and optional details
- **WriteError function**: Provides centralized error response writing
- **ErrorHandler middleware**: Includes panic recovery with basic internal server error response
- **LoggingMiddleware**: Includes request ID tracking and structured logging

### Issues Identified

1. **Inconsistent error responses**: Handlers return early on validation/parsing errors without sending error responses to clients
2. **Limited error categorization**: No standardized error types or codes beyond HTTP status codes
3. **Missing error context**: Errors lack sufficient context for debugging and monitoring
4. **No client-friendly messages**: Error messages are not tailored for API consumers

## Proposed Design

### Error Response Structure

```go
type ErrorResponse struct {
    Code      int         `json:"code"`
    Message   string      `json:"message"`
    Details   any         `json:"details,omitempty"`
    RequestID string      `json:"request_id,omitempty"`
    Timestamp time.Time   `json:"timestamp"`
}
```

### Error Categories and Status Codes

| Category | HTTP Status | Description | Example |
|----------|-------------|-------------|---------|
| ValidationError | 400 | Invalid request data | Missing required field, invalid format |
| NotFoundError | 404 | Resource not found | Tag/Entity ID doesn't exist |
| ConflictError | 409 | Resource conflict | Duplicate name, constraint violation |
| UnauthorizedError | 401 | Authentication required | Missing/invalid token |
| ForbiddenError | 403 | Insufficient permissions | User lacks access rights |
| InternalError | 500 | Server error | Database connection failure, unexpected panic |
| NotImplementedError | 501 | Feature not implemented | Endpoint stub |

### Custom Error Types

```go
type AppError struct {
    Type    ErrorType
    Code    int
    Message string
    Details any
    Cause   error
}

type ErrorType string

const (
    ErrorTypeValidation   ErrorType = "validation"
    ErrorTypeNotFound     ErrorType = "not_found"
    ErrorTypeConflict     ErrorType = "conflict"
    ErrorTypeUnauthorized ErrorType = "unauthorized"
    ErrorTypeForbidden    ErrorType = "forbidden"
    ErrorTypeInternal     ErrorType = "internal"
    ErrorTypeNotImplemented ErrorType = "not_implemented"
)
```

### Handler Error Handling Pattern

```go
func SomeHandler(w http.ResponseWriter, r *http.Request) {
    // Validation
    if err := DecodeAndValidate(r, &input); err != nil {
        WriteAppError(w, r, NewValidationError("Invalid input", err))
        return
    }

    // Business logic
    result, err := someService.DoSomething(input)
    if err != nil {
        WriteAppError(w, r, err) // err should be AppError
        return
    }

    WriteJSON(w, r, http.StatusOK, result)
}
```

### Middleware Integration

- **ErrorHandler**: Enhanced to catch panics and convert to AppError
- **LoggingMiddleware**: Enhanced to log errors with request context
- **Metrics**: Track error rates by type and endpoint

### Error Conversion Functions

```go
func NewValidationError(message string, cause error) *AppError {
    return &AppError{
        Type:    ErrorTypeValidation,
        Code:    http.StatusBadRequest,
        Message: message,
        Details: cause.Error(),
        Cause:   cause,
    }
}

func NewNotFoundError(resource string, id string) *AppError {
    return &AppError{
        Type:    ErrorTypeNotFound,
        Code:    http.StatusNotFound,
        Message: fmt.Sprintf("%s with id %s not found", resource, id),
    }
}
```

### Logging Strategy

- **Client errors (4xx)**: Log at WARN level with request context
- **Server errors (5xx)**: Log at ERROR level with full stack traces
- **Request ID**: Include in all error logs for traceability
- **Metrics**: Increment error counters by type and endpoint

## Implementation Steps

### Phase 1: Core Infrastructure

1. Define AppError type and ErrorType constants
2. Create error conversion functions (NewValidationError, etc.)
3. Enhance ErrorResponse struct with RequestID and Timestamp
4. Implement WriteAppError function

### Phase 2: Middleware Updates

1. Update ErrorHandler to use WriteAppError for panics
2. Enhance LoggingMiddleware to log AppErrors appropriately
3. Add error metrics collection

### Phase 3: Handler Updates

1. Update base.go DecodeAndValidate to return AppError
2. Modify all handlers to use WriteAppError instead of early returns
3. Add proper error handling for database operations
4. Implement NotImplementedError for stub endpoints

### Phase 4: Testing and Validation

1. Add unit tests for error conversion functions
2. Test error responses across all endpoints
3. Validate error logging and metrics
4. Update API documentation with error responses

## Benefits

- **Consistency**: Uniform error response format across all endpoints
- **Debugging**: Request IDs and detailed logging for troubleshooting
- **Monitoring**: Error metrics for operational visibility
- **Client Experience**: Clear, actionable error messages
- **Maintainability**: Centralized error handling logic

## Migration Considerations

- **Backward Compatibility**: Existing error responses should be maintained during transition
- **Gradual Rollout**: Update handlers incrementally to avoid breaking changes
- **Testing**: Comprehensive testing required for error scenarios
- **Documentation**: Update OpenAPI spec with error response schemas
