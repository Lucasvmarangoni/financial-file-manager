package repositories

import (
	"context"
	"log"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/jackc/pgx/v4"
)

type ContractRepository interface {
	InitTable() error
	Insert(contract *entities.Contract) (*entities.Contract, error)
	Find(id string) (*entities.Contract, error)
}

type ContractRepositoryDb struct {
	tx pgx.Tx
}

func NewContractRepository(db pgx.Tx) *ContractRepositoryDb {
	return &ContractRepositoryDb{tx: db}
}

func (r *ContractRepositoryDb) Insert(contract *entities.Contract, ctx context.Context, tx pgx.Tx) (*entities.Contract, error) {
	if contract.File.ID.String() == "" {
		contract.File.ID = pkg_entities.NewID()
	}

	sql := `INSERT INTO contracts (id, type, created_at, customer, title, parties, object, extract, invoice) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.tx.Exec(ctx, sql,
		contract.File.ID,
		contract.File.Type,
		contract.File.CreatedAt,
		contract.File.Customer,
		contract.Title,
		contract.Parties,
		contract.Object,
		contract.Extract,
		contract.Invoice,
	)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (r *ContractRepositoryDb) InitTable(ctx context.Context) error {

	log.Println("Creating accounts table.")
	_, err := r.tx.Exec(ctx, `CREATE TABLE IF NOT EXISTS contracts (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			created_at TIMESTAMP,
			type TEXT,
			customer TEXT,
			title TEXT,
			parties TEXT[],
			object TEXT,
			extract TEXT[],
			invoice TEXT[]
		`)
	if err != nil {
		return err
	}
	return nil
}
