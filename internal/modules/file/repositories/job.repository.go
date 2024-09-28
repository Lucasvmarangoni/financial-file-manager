package repositories

import (
	"context"
	"log"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type JobRepository interface {
	InitTable(ctx context.Context) error
	Insert(job *entities.Job, ctx context.Context, tx pgx.Tx) (*entities.Job, error)
	// Find(id string) (*entities.Job, error)
}

type JobRepositoryDb struct {
	Tx pgx.Tx
}

func NewJobRepository(db pgx.Tx) *JobRepositoryDb {
	return &JobRepositoryDb{Tx: db}
}

func (r *JobRepositoryDb) Insert(job *entities.Job, ctx context.Context, tx pgx.Tx) (*entities.Job, error) {
	if job.ID == uuid.Nil || job.ID.String() == "" {
		job.ID = pkg_entities.NewID()
	}

	sql := `INSERT INTO jobs (id, output_bucket_path, status, file_id, error, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.Tx.Exec(ctx, sql,
		job.ID,
		job.OutputBucketPath,
		job.Status,
		job.FileID,
		job.Error,
		job.CreatedAt,
		job.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (r *JobRepositoryDb) InitTable(ctx context.Context) error {

	log.Println("Creating accounts table.")
	_, err := r.Tx.Exec(ctx, `CREATE TABLE IF NOT EXISTS jobs (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			output_bucket_path TEXT, 
			status TEXT, 
			file_id TEXT, 
			error TEXT, 
			created_at TIMESTAMP, 
			updated_at TIMESTAMP
		`)
	if err != nil {
		return err
	}
	return nil
}
