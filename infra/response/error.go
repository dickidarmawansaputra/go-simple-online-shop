package response

import (
	"errors"
	"net/http"
)

// error general

var (
	ErrNotFound = errors.New("not found")
)

var (
	// products
	ErrProductRequired = errors.New("Product must be required")
	ErrProductInvalid  = errors.New("Product must have minimum 4 characters")
	ErrStockInvalid    = errors.New("Stock must be greater than 0")
	ErrPriceInvalid    = errors.New("Price must be greater than 0")

	// auth
	ErrEmailRequired    = errors.New("Email must be required")
	ErrEmailInvalid     = errors.New("Email is invalid")
	ErrPasswordRequired = errors.New("Password must be required")
	ErrPasswordInvalid  = errors.New("Password must have minimum 6 characters")
	ErrAuthIsNotExists  = errors.New("Auth is not exists")
	ErrAuthIsExists     = errors.New("Auth is already exists")
	ErrEmailAlreadyUsed = errors.New("Email is already used")
	ErrPasswordNotMatch = errors.New("Password is not match")

	// transactions
	ErrAmountInvalid          = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(message string, code string, httpCode int) Error {
	return Error{
		Message:  message,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
)

var (
	// 400 http code, 01 sequence sequential
	ErrorProductRequired  = NewError(ErrProductRequired.Error(), "40001", http.StatusBadRequest)
	ErrorProductInvalid   = NewError(ErrProductInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorStockInvalid     = NewError(ErrStockInvalid.Error(), "40003", http.StatusBadRequest)
	ErrorPriceInvalid     = NewError(ErrPriceInvalid.Error(), "40004", http.StatusBadRequest)
	ErrorEmailRequired    = NewError(ErrEmailRequired.Error(), "40005", http.StatusBadRequest)
	ErrorEmailInvalid     = NewError(ErrEmailInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40007", http.StatusBadRequest)
	ErrorPasswordInvalid  = NewError(ErrPasswordInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorAuthIsNotExists  = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
	ErrorAuthIsExists     = NewError(ErrAuthIsExists.Error(), "40101", http.StatusUnauthorized)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40102", http.StatusUnauthorized)
	ErrorNotFound         = NewError(ErrNotFound.Error(), "40402", http.StatusNotFound)

	ErrorInvalidAmount = NewError(ErrAmountInvalid.Error(), "40009", http.StatusBadRequest)
)

var (
	ErrorMapping = map[string]Error{
		ErrProductRequired.Error():  ErrorProductRequired,
		ErrProductInvalid.Error():   ErrorProductInvalid,
		ErrStockInvalid.Error():     ErrorStockInvalid,
		ErrPriceInvalid.Error():     ErrorPriceInvalid,
		ErrEmailRequired.Error():    ErrorEmailRequired,
		ErrEmailInvalid.Error():     ErrorEmailInvalid,
		ErrPasswordRequired.Error(): ErrorPasswordRequired,
		ErrPasswordInvalid.Error():  ErrorPasswordInvalid,
		ErrAuthIsNotExists.Error():  ErrorAuthIsNotExists,
		ErrAuthIsExists.Error():     ErrorAuthIsExists,
		ErrEmailAlreadyUsed.Error(): ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error(): ErrorPasswordNotMatch,
		ErrNotFound.Error():         ErrorNotFound,
	}
)
