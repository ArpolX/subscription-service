package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// unique violation
func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" // number error
	}
	return false
}
