run:
	go run server.go

build:
	go build server.go

ent-generate-schema:
	go generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema

ent-create: # arguments  [entity name]
	go run -mod=mod entgo.io/ent/cmd/ent new