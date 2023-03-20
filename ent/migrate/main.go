//go:build ignore

// Source: https://entgo.io/docs/versioned-migrations/#option-2-create-a-migration-generation-script

package main

import (
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/migrate/entv2/migrate"
	_ "github.com/lib/pq"
	"github.com/stadio-app/stadio-backend/utils"
)

func main() {
    ctx := context.Background()
    // Create a local migration directory able to understand Atlas migration file format for replay.
    dir, err := atlas.NewLocalDir("ent/migrate/migrations")
    if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }
    // Migrate diff options.
    opts := []schema.MigrateOption{
        schema.WithDir(dir),
        schema.WithMigrationMode(schema.ModeReplay),
        schema.WithDialect(dialect.Postgres),
        schema.WithFormatter(atlas.DefaultFormatter),
    }
    if len(os.Args) != 2 {
        log.Fatalln("migration name is required. Use: 'make atlas-create entity=<name>'")
    }
    // Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
    err = migrate.NamedDiff(ctx, utils.PostgresDNS(), os.Args[1], opts...)
    if err != nil {
        log.Fatalf("failed generating migration file: %v", err)
    }
}
