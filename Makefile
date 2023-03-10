run:
	make ent-generate && go run server.go

build:
	make ent-generate && go build server.go

watch:
	make ent-generate && go run github.com/go-playground/justdoit -build="go build server.go" -run="./server"

ent-generate:
	go generate ./ent

ent-create: # arguments  [entity name]
	go run -mod=mod entgo.io/ent/cmd/ent new

docker-container:
	docker compose -f ./docker/docker-compose.yml up