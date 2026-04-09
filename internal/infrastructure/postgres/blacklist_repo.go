package postgres

import (
	"context"
	"database/sql"

	domain "github.com/tarvarrs/transaction-blacklist-guard/internal/domain/walletoperation"
)

type BlacklistReporitory struct {
	db *sql.DB
}

func NewBlacklistRepository(db *sql.DB) *BlacklistReporitory {
	return &BlacklistReporitory{db: db}
}

func (r *BlacklistReporitory) HasAny(ctx context.Context, ids ...domain.ParticipantID) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM blacklist WHERE id IN ($1, $2))`

	var exists bool
	if err := r.db.QueryRowContext(ctx, query, ids[0], ids[1]).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
