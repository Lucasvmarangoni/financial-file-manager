package repositories

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Insert(user *entities.User) (*entities.User, error)
	Find(id string) (*entities.User, error)
}

type UserRepositoryDb struct {
	tx pgx.Tx
}

func NewUserRepository(db pgx.Tx) *UserRepositoryDb {
	return &UserRepositoryDb{tx: db}
}

func (r *UserRepositoryDb) Insert(user *entities.User, ctx context.Context) (*entities.User, error) {
	if user.ID.String() == "" {
		user.ID = pkg_entities.NewID()
	}

	sql := `INSERT INTO users (id, name, last_name, email, cpf, password, admin, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.tx.Exec(ctx, sql,
		user.ID,
		user.Name,
		user.LastName,
		user.Email,
		user.CPF,
		user.Password,
		user.Admin,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
