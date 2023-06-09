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
	jwtmw "github.com/gofiber/jwt/v3"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sethvargo/go-envconfig"
	_ "github.com/wralith/paiwr/server/docs"
	"github.com/wralith/paiwr/server/pkg/validate"
	"github.com/wralith/paiwr/server/topic"
	"github.com/wralith/paiwr/server/user"
)

type Config struct {
	DbConnStr string `env:"DB_CONN_STR,required"`
	JWTSecret string `env:"JWT_SECRET,required"`
	Port      string `env:"PORT,default=8080"`
}

//	@title			Paiwr Server
//	@version		1.0
//	@description	Paiwr Server

//	@host		localhost:8080
//	@BasePath	/
//	@schemes	http

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

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
		Max: 100,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Static("/docs/spec", "docs/swagger.yaml")
	app.Static("/docs", "docs/docs.html")
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Get("/monitor", monitor.New())

	pool, err := pgxpool.New(context.Background(), config.DbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	validator := validate.NewValidate()

	userRepo := user.NewPgRepo(pool)
	userRoutes := user.NewRoutes(userRepo, config.JWTSecret, validator)

	topicRepo := topic.NewPgRepo(pool)
	topicRoutes := topic.NewRoutes(topicRepo, validator)

	migrateDb(config)

	app.Post("/users/login", userRoutes.Login)
	app.Post("/users/register", userRoutes.Register)
	app.Get("/users/:id", userRoutes.FindByID)

	app.Get("/topics/:id", topicRoutes.FindByID)
	app.Get("/topics/owner/:id", topicRoutes.FindByOwnerID)
	app.Get("/topics/pair/:id", topicRoutes.FindInvolved)

	// Auth Required
	app.Use(jwtmw.New(jwtmw.Config{SigningKey: []byte("secret")}))

	app.Patch("/users/update-password", userRoutes.UpdatePassword)

	app.Post("/topics", topicRoutes.Create)
	app.Delete("/topics/:id", topicRoutes.Delete)

	err = app.Listen(":" + config.Port)
	log.Fatal(err)
}
