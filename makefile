# Docker
docker-container:
	docker compose -f ./docker-compose.yml up

# GraphQL
gql-generate:
	go run github.com/99designs/gqlgen generate

# DB migrations
create-migration:
	go run github.com/ayaanqui/go-migration-tool --directory "./database/migrations" create-migration $(fileName)

# Server
run:
	go run server.go
build:
	go build server.go
build-dev:
	make gql-generate && make build
watch:
	go run github.com/cosmtrek/air -c .air.toml -- -h
