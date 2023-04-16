package main

import (
	"context"
	"log"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sethvargo/go-envconfig"
	"github.com/wralith/paiwr/server/topic"
	"github.com/wralith/paiwr/server/user"
)

func createPool(connStr string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}

// TODO: Default should be replaced with required!
type Config struct {
	DbConnStr string `env:"TOPIC_DB_URI,default=postgresql://root:secret@localhost:5432/paiwr?sslmode=disable"`
}

func main() {
	var config Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max: 20,
	}))
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Get("/monitor", monitor.New())

	pool := createPool(config.DbConnStr)

	userRepo := user.NewPgRepo(pool)
	userRoutes := user.NewRoutes(userRepo)

	topicRepo := topic.NewPgRepo(pool)
	topicRoutes := topic.NewRoutes(topicRepo)

	if err := topicRepo.MigrateWeirdly(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := userRepo.MigrateWeirdly(context.Background()); err != nil {
		log.Fatal(err)
	}

	app.Post("/user/login", userRoutes.Login)
	app.Post("/user/register", userRoutes.Register)
	app.Get("/user/:id", userRoutes.FindByID)
	app.Patch("/user/update-password", userRoutes.UpdatePassword)

	app.Post("/topics", topicRoutes.Create)
	app.Get("/topics/:id", topicRoutes.FindByID)
	app.Get("/topics/owner/:id", topicRoutes.FindByOwnerID)
	app.Get("/topics/pair/:id", topicRoutes.FindInvolved)
	app.Delete("/topics/:id", topicRoutes.Delete)

	err := app.Listen(":8080")
	log.Fatal(err)
}