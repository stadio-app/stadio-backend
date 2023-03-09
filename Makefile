run:
	go run server.go

build:
	go build server.go

ent-generate:
	go generate ./ent

ent-create: # arguments  [entity name]
	go run -mod=mod entgo.io/ent/cmd/ent new

docker:
	docker compose -f ./docker/docker-compose.yml up