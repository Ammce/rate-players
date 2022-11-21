### Player Ratings App

## DB Schema URL

- https://drawsql.app/teams/player-ratings/diagrams/public-schema

## Migrations

- Migrations up

```shell
migrate -path src/adapters/database/postgres/migrations -database "postgres://postgres:postgres@localhost:5432/rate_players?sslmode=disable" up
```

- Migrations create

```shell
migrate create -ext sql -dir src/adapters/database/postgres/migrations player_image_url
```

- Migrations remove dirtyness ( Run the previous timestamp migration before the dirty one with fore)

```shell
migrate -path src/adapters/database/postgres/migrations -database "postgres://postgres:postgres@localhost:5432/rate_players?sslmode=disable" force 20221109210337
```
