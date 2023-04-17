# Paiwr

What it does? Nothing, but hopefully it will get you few study dates? In theory?

![graphic](/docs/graphic.png)

## Setup For Development

#### With Docker Compose

```bash
docker compose up
```

_You might have to `cd client && pnpm install` until I fix `node_modules` related problems in docker-compose._

#### Without

```bash
# Client
cd client
pnpm install # Ofc you can use npm or yarn as well
pnpm dev

# Server
cd server
go run . # DB_CONN_STR and JWT_SECRET environment variables are required

# Postgres Database with Docker
docker run --name postgres-paiwr -d -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:alpine
```

_You can use [`gow`](https://github.com/mitranim/gow) to watch `go` files and rerun `go run` when something changed, similar to `nodemon`._

## Stack

- Server written in `Go` with bunch of libraries: [`Fiber`](https://github.com/gofiber/fiber), [`pgx`](https://github.com/jackc/pgx), [`Squirrel`](https://github.com/Masterminds/squirrel) etc.
- Client written in [Solid](https://github.com/solidjs/solid) with [`Tailwind`](https://tailwindcss.com/), _you can clone clean template with eslint-prettier configured from [Here](https://github.com/wralith/solid-ts-tailwind)_

## Endpoints

```
server default      ::     http://localhost:8080
client default      ::     http://localhost:3000

<server>/metrics                   Prometheus metrics
<server>/swagger or <server>/docs  API docs (Swagger UI and Rapidoc)
<server>/monitor                   Fiber Monitor

<server>/users                     User stuff...
<server>/topics                    Topic stuff...
<client>/...                       Client App
```

## Todo

Lots of stuff...
