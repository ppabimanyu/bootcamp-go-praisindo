package exception

type Code string

const (
	InvalidArgumentCode  Code = "INVALID_ARGUMENT"  // Represents an invalid argument error.
	NotFoundCode         Code = "NOT_FOUND"         // Represents a not found error.
	AlreadyExistsCode    Code = "ALREADY_EXISTS"    // Represents an already exists error.
	PermissionDeniedCode Code = "PERMISSION_DENIED" // Represents a permission denied error.
	UnauthenticatedCode  Code = "UNAUTHENTICATED"   // Represents an unauthenticated error.
	InternalErrorCode    Code = "INTERNAL"          // Represents an internal error.
)

type Exception struct {
	Code    Code
	Message any
	Error   error
}

func (e *Exception) GetError() *string {
	if e.Error != nil {
		err := e.Error.Error()
		return &err
	}
	return nil
}

func (e *Exception) GetHttpCode() int {
	switch e.Code {
	case InvalidArgumentCode:
		return 400
	case NotFoundCode:
		return 404
	case AlreadyExistsCode:
		return 409
	case PermissionDeniedCode:
		return 403
	case UnauthenticatedCode:
		return 401
	case InternalErrorCode:
		return 500
	default:
		return 500
	}
}

func (e *Exception) GetGrpcCode() uint32 {
	switch e.Code {
	case InvalidArgumentCode:
		return 3
	case NotFoundCode:
		return 5
	case AlreadyExistsCode:
		return 6
	case PermissionDeniedCode:
		return 7
	case UnauthenticatedCode:
		return 16
	case InternalErrorCode:
		return 13
	default:
		return 13
	}
}

func InvalidArgument(message any) *Exception {
	return &Exception{
		Code:    InvalidArgumentCode,
		Message: message,
	}
}

func NotFound(message any) *Exception {
	return &Exception{
		Code:    NotFoundCode,
		Message: message,
	}
}

func AlreadyExists(message any) *Exception {
	return &Exception{
		Code:    AlreadyExistsCode,
		Message: message,
	}
}

func PermissionDenied(message any) *Exception {
	return &Exception{
		Code:    PermissionDeniedCode,
		Message: message,
	}
}

func Unauthenticated(message any) *Exception {
	return &Exception{
		Code:    UnauthenticatedCode,
		Message: message,
	}
}

func Internal(message any, err error) *Exception {
	return &Exception{
		Code:    InternalErrorCode,
		Message: message,
		Error:   err,
	}
}

func Conflict(message any) *Exception {
	return &Exception{
		Code:    AlreadyExistsCode,
		Message: message,
	}
}
