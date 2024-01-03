package database

import (
	"context"

	logger "github.com/Lucasvmarangoni/financial-file-manager/pkg/log"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type TableRepository interface {
	InitTables(ctx context.Context) error
	initUserTable(ctx context.Context) error
}

type TableRepositoryDb struct {
	tx pgx.Tx
}

func NewTableRepository(db pgx.Tx) *TableRepositoryDb {
	return &TableRepositoryDb{tx: db}
}

func (r *TableRepositoryDb) initUserTable(ctx context.Context) error {

	_, err := r.tx.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name TEXT,
			last_name TEXT,
			email TEXT UNIQUE,
			cpf TEXT,
			password TEXT,
			admin BOOLEAN,
			created_at TIMESTAMP,
			updated_at TIMESTAMP[]
		)`)
	if err != nil {
		return err
	}
	log.Info().Str("Context", "TableRepository").Msg("Created users table successfully.")
	return nil
}

func (r *TableRepositoryDb) InitTables(ctx context.Context) error {
	err := r.initUserTable(ctx)
	if err != nil {
		return logger.NewError(err, "InitTables")
	}
	return nil
}
