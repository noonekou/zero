package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorCodeType int

const (
	ErrCodeSuccess = 0
	ErrCodeFail    = -1

	ErrCodeNotAuthorized         = 10000
	ErrCodeTokenInvalid          = 10001
	ErrCodeTokenExpired          = 10002
	ErrCodeNoPermission          = 10003
	ErrCodeRoleNotFound          = 10004
	ErrCodeRoleAlreadyExist      = 10005
	ErrCodePermissionNotFound    = 10006
	ErrCodeRoleNameCannotBeEmpty = 10007

	ErrCodeUsernameNotExist          = 20000
	ErrCodeUsernameOrPasswordIsEmpty = 20001
	ErrCodePasswordMissMatch         = 20002
	ErrCodeUsernameAlreadyExist      = 20003
	ErrCodeUserNotFound              = 20004
)

type MyError struct {
	Code    ErrorCodeType `json:"code"`
	Message string        `json:"message"`
}

func (e *MyError) Error() string {
	return e.Message
}

func (e *MyError) GRPCStatus() *status.Status {
	return status.New(codes.Code(e.Code), e.Message)
}

func NewMyError(code ErrorCodeType, message string) *MyError {
	return &MyError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrSuccess = NewMyError(ErrCodeSuccess, "success")
	ErrFail    = NewMyError(ErrCodeFail, "fail")

	ErrNotAuthorized         = NewMyError(ErrCodeNotAuthorized, "not authorized")
	ErrTokenInvalid          = NewMyError(ErrCodeTokenInvalid, "token invalid")
	ErrTokenExpired          = NewMyError(ErrCodeTokenExpired, "token expired")
	ErrNoPermission          = NewMyError(ErrCodeNoPermission, "no permission")
	ErrRoleNotFound          = NewMyError(ErrCodeRoleNotFound, "role not found")
	ErrRoleAlreadyExist      = NewMyError(ErrCodeRoleAlreadyExist, "role already exist")
	ErrPermissionNotFound    = NewMyError(ErrCodePermissionNotFound, "permission not found")
	ErrRoleNameNotFound      = NewMyError(ErrCodeRoleNameCannotBeEmpty, "role name cannot be empty")
	ErrRoleNameCannotBeEmpty = NewMyError(ErrCodeRoleNameCannotBeEmpty, "role name cannot be empty")

	ErrUsernameNotExist          = NewMyError(ErrCodeUsernameNotExist, "username not exist")
	ErrUsernameOrPasswordIsEmpty = NewMyError(ErrCodeUsernameOrPasswordIsEmpty, "username or password is empty")
	ErrPasswordMissMatch         = NewMyError(ErrCodePasswordMissMatch, "password miss match")
	ErrUsernameAlreadyExist      = NewMyError(ErrCodeUsernameAlreadyExist, "username already exist")
	ErrUserNotFound              = NewMyError(ErrCodeUserNotFound, "user not found")
)

func IsMyError(err error, code ErrorCodeType) bool {
	if myErr, ok := err.(*MyError); ok {
		return myErr.Code == code
	}
	return false
}

func GetMyError(err error) *MyError {
	if myErr, ok := err.(*MyError); ok {
		return myErr
	}
	return nil
}
