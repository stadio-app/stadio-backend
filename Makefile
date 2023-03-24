# Docker
docker-container:
	docker compose -f ./docker/docker-compose.yml up

# Ent. Framework
ent-generate:
	go generate ./ent
ent-create:
	go run -mod=mod entgo.io/ent/cmd/ent new $(entity)

# Atlas (DB migration manager)
atlas-create:
	go run -mod=mod ent/migrate/migratec.go $(entity)
atlas-validate:
	atlas migrate validate --dir file://ent/migrate/migrations

# GraphQL
gql-generate:
	go run github.com/99designs/gqlgen generate

# Server
generate:
	make ent-generate && make gql-generate
run:
	make generate && go run server.go
build:
	make generate && go build server.go
watch:
	make generate && go run github.com/go-playground/justdoit -build="go build server.go" -run="./server"
