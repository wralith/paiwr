package topic

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo interface {
	MigrateWeirdly(ctx context.Context) error
	DropWeirdly(ctx context.Context) error
	Save(ctx context.Context, t *Topic) error
	FindByID(ctx context.Context, id uuid.UUID) (*Topic, error)
	FindByOwner(ctx context.Context, id uuid.UUID) ([]Topic, error)
	FindInvolved(ctx context.Context, id uuid.UUID) ([]Topic, error)
	Update(ctx context.Context, t *Topic) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PgRepo struct {
	pool *pgxpool.Pool
}

func NewPgRepo(pool *pgxpool.Pool) *PgRepo {
	return &PgRepo{pool: pool}
}

var _ Repo = &PgRepo{}

// well it is a poc, definitely a some migration tool should be used!
func (r *PgRepo) MigrateWeirdly(ctx context.Context) error {
	_, err := r.pool.Exec(ctx,
		`create table if not exists topics (
		id uuid primary key,
		category varchar,
		title varchar,
		capacity int,
		owner uuid,
		parties uuid[],
		created_at timestamptz,
		updated_at timestamptz,
		finished_at timestamptz
	)`)
	return err
}

func (r *PgRepo) DropWeirdly(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, "drop table if exists topics")
	return err
}

func (r *PgRepo) Save(ctx context.Context, t *Topic) error {
	sql, args, err := sq.Insert("topics").
		Columns("id", "category", "title", "capacity", "owner", "parties", "created_at", "updated_at", "finished_at").
		Values(t.ID, t.Cateogry, t.Title, t.Capacity, t.Owner, t.Parties, t.CreatedAt, t.UpdatedAt, t.FinishedAt).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return err
	}
	_, err = r.pool.Exec(ctx, sql, args...)
	return err
}

func (r *PgRepo) FindByID(ctx context.Context, id uuid.UUID) (*Topic, error) {
	// var topic Topic
	sql, args, err := sq.Select("id", "category", "title", "capacity", "owner", "parties", "created_at", "updated_at", "finished_at").
		From("topics").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return r.collectOneTopic(ctx, rows)
}

func (r *PgRepo) FindByOwner(ctx context.Context, id uuid.UUID) ([]Topic, error) {
	sql, args, err := sq.Select("id", "category", "title", "capacity", "owner", "parties", "created_at", "updated_at", "finished_at").
		From("topics").
		Where(sq.Eq{"owner": id}).
		OrderBy("created_at").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return r.collectTopics(ctx, rows)
}

func (r *PgRepo) FindInvolved(ctx context.Context, id uuid.UUID) ([]Topic, error) {
	sql, args, err := sq.Select("id", "category", "title", "capacity", "owner", "parties", "created_at", "updated_at", "finished_at").
		From("topics").
		Where("? = any(parties)", id).
		OrderBy("created_at").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return r.collectTopics(ctx, rows)
}

func (r *PgRepo) Update(ctx context.Context, t *Topic) error {
	sql, args, err := sq.Update("topics").
		Set("category", t.Cateogry).
		Set("title", t.Title).
		Set("capacity", t.Capacity).
		Set("owner", t.Owner).
		Set("parties", t.Parties).
		Set("updated_at", t.UpdatedAt).
		Set("finished_at", t.FinishedAt).
		Where(sq.Eq{"id": t.ID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	return err
}

func (r *PgRepo) Delete(ctx context.Context, id uuid.UUID) error {
	sql, args, err := sq.Delete("topics").Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	return err
}

// collectOneTopic wraps pgx.CollectOneRow to collect first found row as Topic struct
func (r *PgRepo) collectOneTopic(ctx context.Context, rows pgx.Rows) (*Topic, error) {
	topic, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Topic])
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

// collectTopics wraps pgx.CollectRows
func (r *PgRepo) collectTopics(ctx context.Context, rows pgx.Rows) ([]Topic, error) {
	topic, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Topic])
	if err != nil {
		return nil, err
	}

	return topic, nil
}
