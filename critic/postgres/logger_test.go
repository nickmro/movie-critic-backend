package postgres_test

import "fmt"

// Logger returns a test logger.
type Logger struct{}

func (l *Logger) Error(args ...interface{}) {
	fmt.Println(args...)
}
