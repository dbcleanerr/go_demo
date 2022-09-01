postgres:
	docker run -d --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=welcome postgres:14.4

createdb:
	docker exec -it postgres14 psql -U postgres -d postgres  -c "create database simple_bank"

dropdb:
	docker exec -it postgres14 psql -U postgres -d postgres  -c "drop database simple_bank"

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

.PHONY: createdb postgres dropdb