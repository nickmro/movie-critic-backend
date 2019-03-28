package postgres

// Error represents a postgres error.
type Error string

// The possible database errors.
const (
	ErrNotFound = Error("not_found")
)

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
