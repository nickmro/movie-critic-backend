package postgres

import backend "github.com/nickmro/movie-critic-backend"

// Error represents a postgres error.
type Error string

// The possible database errors.
const (
	ErrNotFound = Error("not_found")
	ErrInternal = Error("internal")
)

// LogError logs errors.
func LogError(logger backend.ErrorLogger, err error) {
	if logger != nil {
		logger.Error(err)
	}
}

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
