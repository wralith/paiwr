package topic

import (
	"context"

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
	_, err := r.pool.Exec(ctx,
		`insert into topics
		(id, category, title, capacity, owner, parties, created_at, updated_at, finished_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		t.ID, t.Cateogry, t.Title, t.Capacity, t.Owner, t.Parties, t.CreatedAt, t.UpdatedAt, t.FinishedAt)

	return err
}

func (r *PgRepo) FindByID(ctx context.Context, id uuid.UUID) (*Topic, error) {
	// var topic Topic
	rows, err := r.pool.Query(ctx,
		`select id, category, title, capacity, owner, parties, created_at, updated_at, finished_at
		from topics where id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return r.collectOneTopic(ctx, rows)
}

func (r *PgRepo) FindByOwner(ctx context.Context, id uuid.UUID) ([]Topic, error) {
	// var topic Topic
	rows, err := r.pool.Query(ctx,
		`select id, category, title, capacity, owner, parties, created_at, updated_at, finished_at
		from topics where owner=$1 order by created_at`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return r.collectTopics(ctx, rows)
}

func (r *PgRepo) FindInvolved(ctx context.Context, id uuid.UUID) ([]Topic, error) {
	rows, err := r.pool.Query(ctx,
		`select id, category, title, capacity, owner, parties, created_at, updated_at, finished_at
		from topics where $1=any(parties) order by created_at`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return r.collectTopics(ctx, rows)
}

func (r *PgRepo) Update(ctx context.Context, t *Topic) error {
	_, err := r.pool.Exec(ctx,
		`update topics
		set category=$1, title=$2, capacity=$3, owner=$4, parties=$5, created_at=$6, updated_at=$7, finished_at=$8
		where id=$9`,
		t.Cateogry, t.Title, t.Capacity, t.Owner, t.Parties, t.CreatedAt, t.UpdatedAt, t.FinishedAt, t.ID)
	return err
}

func (r *PgRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `delete from topics where id=$1`, id)
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
