package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	username = "test"
	password = "test"
	email    = "test@test.com"
)
var user, _ = CreateUser(CreateUserOpts{Username: username, Password: password, Email: email})

var repo *PgRepo
var connStr = ""

func TestMain(m *testing.M) {
	p := postgres.Preset(
		postgres.WithUser("test", "secret"),
		postgres.WithDatabase("users"),
	)
	c, _ := gnomock.Start(p)
	defer func() { _ = gnomock.Stop(c) }()

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "test", "secret", c.Host, c.DefaultPort(), "users")
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

	err := repo.Save(context.Background(), &user)
	assert.NoError(t, err)
	err = repo.Save(context.Background(), &user)
	assert.Error(t, err)
}

func TestFindByID(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &user)
	assert.NoError(t, err)
	got, err := repo.FindByID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, got.Username)
	assert.WithinDuration(t, user.CreatedAt, got.CreatedAt, time.Second)
}

func TestFindByUsername(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &user)
	assert.NoError(t, err)
	got, err := repo.FindByUsername(context.Background(), user.Username)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, got.Email)
	assert.WithinDuration(t, user.CreatedAt, got.CreatedAt, time.Second)
}

func TestUpdate(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &user)
	assert.NoError(t, err)
	got, err := repo.FindByID(context.Background(), user.ID)
	assert.NoError(t, err)

	updatedUser := user
	updatedUser.UpdateBio("Hello")
	updatedUser.UpdateEmail("tester@test.com")
	err = repo.Update(context.Background(), &updatedUser)
	assert.NoError(t, err)
	got, err = repo.FindByID(context.Background(), user.ID)
	assert.NoError(t, err)

	assert.Equal(t, got.Bio, "Hello")
	assert.Equal(t, got.Email, "tester@test.com")
}

func TestVerifyPassword(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := repo.Save(context.Background(), &user)
	assert.NoError(t, err)

	err = repo.VerifyPassword(context.Background(), username, password)
	require.NoError(t, err)
	err = repo.VerifyPassword(context.Background(), username, "wrong")
	require.Error(t, err)
}
