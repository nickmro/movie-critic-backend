package postgres

import (
	"database/sql"
)

// Tx represents a database transaction.
type Tx struct {
	*sql.Tx
}
