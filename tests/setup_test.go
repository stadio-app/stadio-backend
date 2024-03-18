package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ayaanqui/go-migration-tool/migration_tool"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stadio-app/stadio-backend/services"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"

	_ "github.com/lib/pq"
)

var db *sql.DB
var app types.ServerBase
var service services.Service
var ctx context.Context

func NewMockServer() {
	chi_router := chi.NewRouter()
	app = types.ServerBase{
		DB: db,
		Router: chi_router,
		StructValidator: validator.New(),
		MigrationDirectory: "../database/migrations",
	}

	// Run DB migrations
	db_migration := migration_tool.New(app.DB, &migration_tool.Config{
		Directory: app.MigrationDirectory,
		TableName: "migration",
	})
	db_migration.RunMigration()

	var tokens types.Tokens
	if err := utils.FileMapper("../tokens.json", &tokens); err != nil {
		log.Fatalf(err.Error())
	}

	app.Tokens = &tokens
	service = services.Service{
		DB: app.DB,
		StructValidator: app.StructValidator,
		Tokens: &tokens,
	}
}

func NewTestDb() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Build and run the given Dockerfile
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=postgres",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		hostAndPort := resource.GetHostPort("5432/tcp")
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@%s/postgres?sslmode=disable", hostAndPort))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

func TestMain(m *testing.M) {
	NewTestDb()
	NewMockServer()
	ctx = context.Background()

	// run tests...
	exitCode := m.Run()

	os.Exit(exitCode)
}
