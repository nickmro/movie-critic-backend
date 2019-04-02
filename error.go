package backend

// ErrorLogger defines the operations for an error logger.
type ErrorLogger interface {
	Error(args ...interface{})
}
