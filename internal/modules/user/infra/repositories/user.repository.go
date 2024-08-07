package repositories

import (
	"context"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Insert(user *entities.User, ctx context.Context) error
	FindByEmail(hashEmail string, ctx context.Context) (*entities.User, error)
	FindById(id pkg_entities.ID, ctx context.Context) (*entities.User, error)
	FindByCpf(hashCPF string, ctx context.Context) (*entities.User, error)
	Update(user *entities.User, ctx context.Context) error
	UpdateOTP(user *entities.User, ctx context.Context) error
	UpdateOTPVerify(user *entities.User, ctx context.Context) error
	Delete(id string, ctx context.Context) error
	CheckIfUserAlreadyExists(hashEmail, hashCPF string, ctx context.Context) (bool, error)
	GetAllEmailAndCPF() (*entities.UserEmailAndCPF, error)
}

type UserRepositoryDb struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepositoryDb {
	return &UserRepositoryDb{
		conn: conn,
	}
}

func (r *UserRepositoryDb) Insert(user *entities.User, ctx context.Context) error {
	if user.ID.String() == "" {
		user.ID = pkg_entities.NewID()
	}
	sql := `INSERT INTO users (id, name, last_name, cpf, hash_cpf, email, hash_email, password, otp_secret, otp_auth_url, otp_enabled, created_at, update_log) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {

		_, err := tx.Exec(ctx, sql,
			user.ID,
			user.Name,
			user.LastName,
			user.CPF,
			user.HashCPF,
			user.Email,
			user.HashEmail,
			user.Password,
			user.OtpSecret,
			user.OtpAuthUrl,
			user.OtpEnabled,
			user.CreatedAt,
			user.UpdateLog,
		)
		if err != nil {
			return errors.ErrCtx(err, "r.tx.Exec")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) FindById(id pkg_entities.ID, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE id = $1`
	var row pgx.Row
	user := &entities.User{}
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row = tx.QueryRow(ctx, sql, id)
		err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.CPF, &user.HashCPF, &user.Email, &user.HashEmail, &user.Password, &user.OtpSecret, &user.OtpAuthUrl, &user.OtpEnabled, &user.CreatedAt, &user.UpdateLog)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return user, nil
}

func (r *UserRepositoryDb) FindByEmail(hashEmail string, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE hash_email = $1`
	var row pgx.Row
	user := &entities.User{}
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row = tx.QueryRow(ctx, sql, hashEmail)
		err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.CPF, &user.HashCPF, &user.Email, &user.HashEmail, &user.Password, &user.OtpSecret, &user.OtpAuthUrl, &user.OtpEnabled, &user.CreatedAt, &user.UpdateLog)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return user, nil
}

func (r *UserRepositoryDb) FindByCpf(hashCPF string, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE hash_cpf = $1`
	var row pgx.Row
	user := &entities.User{}
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row = tx.QueryRow(ctx, sql, hashCPF)
		err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.CPF, &user.HashCPF, &user.Email, &user.HashEmail, &user.Password, &user.OtpEnabled, &user.CreatedAt, &user.UpdateLog)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return user, nil
}

func (r *UserRepositoryDb) Update(user *entities.User, ctx context.Context) error {
	sql := `UPDATE users SET name = $2, last_name = $3, email = $4, hash_email = $5, password = $6, update_log = $7 WHERE id = $1`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql,
			user.ID,
			user.Name,
			user.LastName,
			user.Email,
			user.HashEmail,
			user.Password,
			user.UpdateLog,
		)
		if err != nil {
			return errors.ErrCtx(err, "tx.Exec(")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) UpdateOTP(user *entities.User, ctx context.Context) error {
	sql := `UPDATE users SET otp_auth_url = $2, otp_secret = $3 WHERE id = $1`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql,
			user.ID,
			user.OtpAuthUrl,
			user.OtpSecret,
		)
		if err != nil {
			return errors.ErrCtx(err, "tx.Exec")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) UpdateOTPVerify(user *entities.User, ctx context.Context) error {
	sql := `UPDATE users SET otp_enabled = $2 WHERE id = $1`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql,
			user.ID,
			user.OtpEnabled,
		)
		if err != nil {
			return errors.ErrCtx(err, "tx.Exec")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) Delete(id string, ctx context.Context) error {
	sql := `DELETE FROM users WHERE id = $1`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql, id)
		if err != nil {
			return errors.ErrCtx(err, "tx.Exec(")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) CheckIfUserAlreadyExists(hashEmail, hashCPF string, ctx context.Context) (bool, error) {
	sql := `SELECT COUNT(*) FROM users WHERE hash_email = $1 OR hash_cpf = $2`
	var count int
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, sql, hashEmail, hashCPF)
		err := row.Scan(&count)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return false, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return count > 0, nil
}


func (r *UserRepositoryDb) GetAllEmailAndCPF() (*entities.UserEmailAndCPF, error) {
	ctx := context.Background()

	sql := `SELECT hash_cpf, hash_email FROM users;`
	var rows pgx.Rows
	emails := []string{}
	cpfs := []string{}

	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var err error
		rows, err = tx.Query(ctx, sql)
		if err != nil {
			return errors.ErrCtx(err, "tx.Query")
		}
		defer rows.Close()

		for rows.Next() {
			var cpf, email string
			err := rows.Scan(&cpf, &email)
			if err != nil {
				return errors.ErrCtx(err, "rows.Scan")
			}
			cpfs = append(cpfs, cpf)
			emails = append(emails, email)
		}

		if err := rows.Err(); err != nil {
			return errors.ErrCtx(err, "rows.Err")
		}

		return nil
	})

	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}

	return &entities.UserEmailAndCPF{
		Emails: emails,
		CPFs:   cpfs,
	}, nil
}