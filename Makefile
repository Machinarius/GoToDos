.PHONY: up down run

up:
	docker compose -f deploy/docker-compose.yml up

down:
	docker compose -f deploy/docker-compose.yml down

run:
	- go run ./cmd/server.go