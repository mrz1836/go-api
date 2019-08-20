// Package models extends the schema package for model management
package models

import (
	"context"
	"database/sql"

	"github.com/mrz1836/go-api/models/schema"
)

func PersonExists(ctx context.Context, db *sql.DB, p *schema.Person) (bool, error) {
	return schema.PersonExists(ctx, db, p.ID)
}
