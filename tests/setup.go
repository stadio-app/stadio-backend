package tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ayaanqui/go-migration-tool/migration_tool"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/stadio-app/stadio-backend/database"
	"github.com/stadio-app/stadio-backend/services"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"

	_ "github.com/lib/pq"
)

func NewMockDB(t *testing.T) *sql.DB {
	db, err := database.CreateDbConnection(database.DbConnection{
		Host: "localhost", 
		Username: "postgres", 
		Password: "postgres", 
		DbName: "postgres", 
		Port: 54312,
		SslMode: false,
	})
	if err != nil {
		fmt.Println("could not establish connection with test db", err)
		t.FailNow() 
		return nil
	}
	return db
}

func NewMockServer(t *testing.T) (types.ServerBase, services.Service) {
	db_conn := NewMockDB(t)
	chi_router := chi.NewRouter()
	server := types.ServerBase{
		DB: db_conn,
		Router: chi_router,
		StructValidator: validator.New(),
		MigrationDirectory: "../database/migrations",
	}

	// Run DB migrations
	db_migration := migration_tool.New(server.DB, &migration_tool.Config{
		Directory: server.MigrationDirectory,
		TableName: "migration",
	})
	db_migration.RunMigration()

	var tokens types.Tokens
	if err := utils.FileMapper("../tokens.json", &tokens); err != nil {
		panic(err)
	}

	server.Tokens = &tokens
	service := services.Service{
		DB: server.DB,
		StructValidator: server.StructValidator,
		Tokens: &tokens,
	}
	return server, service
}
