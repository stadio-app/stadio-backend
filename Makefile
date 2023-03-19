run:
	make generate && go run server.go

build:
	make generate && go build server.go

watch:
	make generate && go run github.com/go-playground/justdoit -build="go build server.go" -run="./server"

ent-generate:
	go generate ./ent

ent-create:
	go run -mod=mod entgo.io/ent/cmd/ent new $(entity)

atlas-create:
	go run -mod=mod ent/migrate/main.go $(entity)

gql-generate:
	go run github.com/99designs/gqlgen generate

generate:
	make ent-generate && make gql-generate

docker-container:
	docker compose -f ./docker/docker-compose.yml up
