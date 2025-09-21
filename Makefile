include .env

.PHONY: migration-down run stop full-stop test

DATABASE_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

migrate-down:
	migrate -path=migrations -database=${DATABASE_URL} down

run:
	docker compose up --build

stop:
	docker compose down

full-stop:
	docker compose down -v

args?=./...
test:
	go test $(args)

build:
	go build -o youtofy-authentication .
