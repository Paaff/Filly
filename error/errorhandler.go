package errorhandler

import "net/http"

// Custom error created for handling HTTP errors more fluently.
type AppError struct {
	Error   error
	Message string
	Code    int
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}
