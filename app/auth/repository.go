package auth

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateAuth(ctx context.Context, model AuthEntity) (err error) {
	query := "INSERT INTO auth (email, password, role, created_at, updated_at, public_id) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = r.db.ExecContext(ctx, query, model.Email, model.Password, model.Role, model.CreatedAt, model.UpdatedAt, model.PublicId)
	if err != nil {
		return
	}

	return
}

func (r repository) GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error) {
	query := "SELECT id, email, password, role, created_at, updated_at, public_id FROM auth WHERE email = ?"

	row := r.db.QueryRowContext(ctx, query, email)
	if row.Err() != nil {
		return
	}

	err = row.Scan(
		&model.Id,
		&model.Email,
		&model.Password,
		&model.Role,
		&model.CreatedAt,
		&model.UpdatedAt,
		&model.PublicId,
	)
	if err != nil {
		return
	}

	return
}
