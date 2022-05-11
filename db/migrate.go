package main

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/lib/pq"
)

func main() {
	// We need a name for the new migration file.
	if len(os.Args) < 2 {
		log.Fatalln("no name given")
	}
	migrationName := os.Args[1]

	// Load the graph.
	graph, err := entc.LoadGraph("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// get tables
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}

	// Create a local migration directory.
	d, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalln(err)
	}

	// Open connection to the database.
	connectionString := "postgresql://postgres:postgres@127.0.0.1/test_db?sslmode=disable"
	// "root:pass@tcp(localhost:3306)/test"
	dlct, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	// Inspect it and compare it with the graph.
	m, err := schema.NewMigrate(dlct, schema.WithDir(d))
	if err != nil {
		log.Fatalln(err)
	}

	if err := m.NamedDiff(context.Background(), migrationName, tbls...); err != nil {
		log.Fatalln(err)
	}
}

// sudo -i
// curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
// echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
// apt-get update
// apt-get install -y migrate
// exit
