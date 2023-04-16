package user

import (
	"context"
	"fmt"
	"testing"
	"time"

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

func TestMain(m *testing.M) {
	p := postgres.Preset(
		postgres.WithUser("test", "secret"),
		postgres.WithDatabase("users"),
	)
	c, _ := gnomock.Start(p)
	defer func() { _ = gnomock.Stop(c) }()

	pool, _ := pgxpool.New(context.Background(), fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.DefaultPort(), "test", "secret", "users",
	))

	repo = NewPgRepo(pool)
	m.Run()
}

func setup(t *testing.T) {
	err := repo.MigrateWeirdly(context.Background())
	require.NoError(t, err)
}

func teardown(t *testing.T) {
	err := repo.DropWeirdly(context.Background())
	require.NoError(t, err)
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
