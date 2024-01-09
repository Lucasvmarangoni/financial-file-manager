package repositories

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Insert(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindById(id pkg_entities.ID) (*entities.User, error)
	FindByCpf(cpf string) (*entities.User, error)
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

func (r *UserRepositoryDb) FindByEmail(email string, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE email = $1`
	row := r.tx.QueryRow(ctx, sql, email)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.CPF, &user.Password, &user.Admin, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryDb) FindById(id pkg_entities.ID, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE id = $1`
	row := r.tx.QueryRow(ctx, sql, id)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.CPF, &user.Password, &user.Admin, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryDb) FindByCpf(cpf string, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE cpf = $1`
	row := r.tx.QueryRow(ctx, sql, cpf)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.CPF, &user.Password, &user.Admin, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryDb) Update(user *entities.User, ctx context.Context) (*entities.User, error) {
	sql := `UPDATE users SET name = $2, last_name = $3, email = $4, cpf = $5, password = $6, admin = $7, updated_at = $8 WHERE id = $1`
	_, err := r.tx.Exec(ctx, sql,
		user.ID,
		user.Name,
		user.LastName,
		user.Email,
		user.CPF,
		user.Password,
		user.Admin,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryDb) ToggleAdmin(id string, ctx context.Context) error {
	sql := `UPDATE users SET admin = NOT admin WHERE id = $1`
	_, err := r.tx.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryDb) Delete(id string, ctx context.Context) error {
	sql := `DELETE FROM users WHERE id = $1`
	_, err := r.tx.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}
