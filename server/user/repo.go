package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Repo interface {
	MigrateWeirdly(ctx context.Context) error
	DropWeirdly(ctx context.Context) error
	Save(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	VerifyPassword(ctx context.Context, username, password string) error // IDK
	Update(ctx context.Context, u *User) error
}

type PgRepo struct {
	pool *pgxpool.Pool
}

func NewPgRepo(pool *pgxpool.Pool) *PgRepo {
	return &PgRepo{pool: pool}
}

var _ Repo = &PgRepo{}

func (r *PgRepo) MigrateWeirdly(ctx context.Context) error {
	_, err := r.pool.Exec(ctx,
		`create table if not exists users (
		id uuid primary key,
		username varchar,
		email varchar,
		hashed_password bytea,
		bio varchar,
		created_at timestamptz,
		updated_at timestamptz
	)`)
	return err
}

func (r *PgRepo) DropWeirdly(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, "drop table if exists users")
	return err
}

func (r *PgRepo) Save(ctx context.Context, u *User) error {
	sql, args, err := sq.Insert("users").
		Columns("id", "username", "email", "hashed_password", "bio", "created_at", "updated_at").
		Values(u.ID, u.Username, u.Email, u.HashedPassword, u.Bio, u.CreatedAt, u.UpdatedAt).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	return err
}

func (r *PgRepo) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	sql, args, err := sq.Select("id", "username", "email", "hashed_password", "bio", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PgRepo) FindByUsername(ctx context.Context, username string) (*User, error) {
	sql, args, err := sq.Select("id", "username", "email", "hashed_password", "bio", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PgRepo) Update(ctx context.Context, u *User) error {
	sql, args, err := sq.Update("users").
		Set("email", u.Email).
		Set("hashed_password", u.HashedPassword).
		Set("bio", u.Bio).
		Set("updated_at", u.UpdatedAt).
		Where(sq.Eq{"id": u.ID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	_, err = r.pool.Exec(ctx, sql, args...)
	return err
}

// VerifyPassword returns nil on success
// TODO: This looks unneccessary or should return user too
func (r *PgRepo) VerifyPassword(ctx context.Context, username, password string) error {
	sql, args, err := sq.Select("hashed_password").
		From("users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	var currentPassword []byte
	err = r.pool.QueryRow(ctx, sql, args...).Scan(&currentPassword)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(currentPassword, []byte(password))
}
