package errs

import "net/http"

var (
	ErrorInternalServer = NewAppError("internal server error", http.StatusInternalServerError)
	ErrorNotfound       = NewAppError("error not found ", http.StatusNotFound)
	ErrorAuthorization  = NewAppError("error authorization", http.StatusUnauthorized)
	ErrorBadRequest     = NewAppError("error bad reqest", http.StatusBadRequest)
)

type AppError struct {
	Msg      string
	HttpCode int
}

func (e AppError) Error() string {
	return e.Msg
}

func NewAppError(msg string, code int) error {
	return &AppError{msg, code}
}
