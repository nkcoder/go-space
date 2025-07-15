## Migrate

Create migration files:
```shell
export DATABASE_URL="postgres://greenlight:Password@2025@localhost:54322/greenlight?sslmode=disable"
```

Migrate:
```shell
migrate -path=./migrations -database="$DATABASE_URL" up
migrate -path=./migrations -database="$DATABASE_URL" down
migrate -path=./migrations -database="$DATABASE_URL" goto 1

# force the database version number to 1, don't run migrations (ignore dirty state)
migrate -path=./migrations -database="$DATABASE_URL" force 1
```