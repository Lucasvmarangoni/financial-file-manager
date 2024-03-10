package database

import (
	"context"

	"github.com/Lucasvmarangoni/logella/err"
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
			id UUID PRIMARY KEY NOT NULL,
			name VARCHAR(10) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			cpf TEXT NOT NULL,
			hash_cpf CHAR(64) UNIQUE NOT NULL,			
			email TEXT NOT NULL,	
			hash_email CHAR(64) UNIQUE NOT NULL,		
			password TEXT NOT NULL,		
			otp_secret TEXT,
			otp_auth_url TEXT,			
			otp_enabled BOOL,
			created_at TIMESTAMP NOT NULL,
			update_log JSONB
		)`)
	if err != nil {
		return errors.ErrCtx(err, "r.tx.Exec")
	}
	log.Info().Str("context", "TableRepository").Msg("Database - Created users table successfully.")
	return nil
}

func (r *TableRepositoryDb) InitTables(ctx context.Context) error {
	err := r.initUserTable(ctx)
	if err != nil {
		return errors.ErrCtx(err, "initUserTable")
	}
	return nil
}
