package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ParseUUID(id string) (pgtype.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid user ID format: %w", err)
	}
	pgUUID := pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}

	return pgUUID, nil
}
