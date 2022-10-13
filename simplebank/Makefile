DB_URL=postgresql://postgres:welcome@localhost:5432/simple_bank?sslmode=disable
BINARY=simplebank

postgres:
	docker run -d --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=welcome postgres:14.4

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

createdb:
	docker exec -it postgres14 createdb -U postgres simple_bank

dropdb:
	docker exec -it postgres14 dropdb -U postgres simple_bank

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

psql:
	docker exec -it postgres14 psql -U postgres -d simple_bank

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

.PHONY: postgres migrateup migrateup1 migratedown migratedown1 createdb dropdb sqlc psql build