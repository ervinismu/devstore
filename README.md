# Devstore

## Prerequisites

- install docker compose

## Running Locally

1. `make environment`
2. `make migration-up`
3. create `app.env` based on `app.env.sample`
4. update `app.env`
5. `make server`
6. App running!

## Other Commands :

You can run `make help` for showing all available commands :

```bash
❯ make help
environment                    Setup environment.
help                           You are here! showing all command documenentation.
migrate-all                    Rollback migrations, all migrations
migrate-create                 Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-down                   Rollback migrations, latest migration (1)
migrate-up                     Run migrations UP
server                         Running application
shell-db                       Enter to database console
```
