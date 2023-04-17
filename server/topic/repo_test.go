package topic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ownerID = uuid.New()
var secondUserID = uuid.New()
var t1 = CreateTopic(CreateTopicOpts{
	Category: Software,
	Title:    "test",
	Capacity: 4,
	Owner:    ownerID,
})
var t2 = CreateTopic(CreateTopicOpts{
	Category: SocialSciences,
	Title:    "test marx",
	Capacity: 7,
	Owner:    ownerID,
})

var repo *PgRepo
var connStr = ""

func TestMain(m *testing.M) {
	p := postgres.Preset(
		postgres.WithUser("test", "secret"),
		postgres.WithDatabase("topics"),
	)
	c, _ := gnomock.Start(p)
	defer func() { _ = gnomock.Stop(c) }()

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "test", "secret", c.Host, c.DefaultPort(), "topics")
	pool, _ := pgxpool.New(context.Background(), connStr)

	repo = NewPgRepo(pool)
	m.Run()
}

func setup(t *testing.T) {
	fmt.Println(connStr)
	m, err := migrate.New("file://../migrations", connStr)
	require.NoError(t, err)
	err = m.Up()
	require.NoError(t, err)
}

func teardown(t *testing.T) {
	m, err := migrate.New("file://../migrations", connStr)
	require.NoError(t, err)
	m.Down()
}

func TestSave(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &t1)
	assert.NoError(t, err)
	err = repo.Save(context.Background(), &t1)
	assert.Error(t, err) // Topic with id already exists
}

func TestFindByID(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &t1)
	assert.NoError(t, err)
	got, err := repo.FindByID(context.Background(), t1.ID)
	assert.NoError(t, err)
	assert.Equal(t, t1.Title, got.Title)
	assert.WithinDuration(t, t1.CreatedAt, got.CreatedAt, time.Second)
}

func TestFindByOwner(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &t1)
	assert.NoError(t, err)
	err = repo.Save(context.Background(), &t2)
	assert.NoError(t, err)
	got, err := repo.FindByOwner(context.Background(), ownerID)
	assert.Equal(t, t1.Title, got[0].Title)
	assert.WithinDuration(t, t1.CreatedAt, got[0].CreatedAt, time.Second)
	assert.Equal(t, t2.Title, got[1].Title)
	assert.WithinDuration(t, t2.CreatedAt, got[1].CreatedAt, time.Second)
}

func TestFindWhereParties(t *testing.T) {
	setup(t)
	defer teardown(t)

	t3 := t1
	err := t3.AddParties(secondUserID)
	assert.NoError(t, err)
	err = repo.Save(context.Background(), &t3)
	got, err := repo.FindInvolved(context.Background(), secondUserID)
	assert.Equal(t, t1.Title, got[0].Title)
	assert.WithinDuration(t, t1.CreatedAt, got[0].CreatedAt, time.Second)
}

func TestUpdate(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &t1)
	assert.NoError(t, err)
	t2 := t1
	t2.Finish()
	err = repo.Update(context.Background(), &t2)
	assert.NoError(t, err)
	got, err := repo.FindByID(context.Background(), t1.ID)
	assert.WithinDuration(t, time.Now(), got.FinishedAt, time.Second)
}

func TestDelete(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &t1)
	assert.NoError(t, err)
	err = repo.Delete(context.Background(), t1.ID)
	assert.NoError(t, err)
	_, err = repo.FindByID(context.Background(), t1.ID)
	assert.Error(t, err)
}
