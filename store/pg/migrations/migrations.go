package migrations

import "database/sql"

type schema struct {
	v int
	q string
}

var schemas = []schema{
	v1schema,
}

// SchemaVersion will find the schema version for the given database
func SchemaVersion(db *sql.DB) int {
	var v int

	rw := db.QueryRow("SELECT schema_version FROM database_information")
	err := rw.Scan(&v)
	if err != nil {
		return 0
	}

	return v

}

// RunMigrations will run all database migrations depending on the version
// returned from the database_information table.
func RunMigrations(db *sql.DB) error {
	version := SchemaVersion(db)
	"log.Printf"("Current database version %d\n", version)

	for _, schema := range schemas {
		version = SchemaVersion(db)

		if version < schema.v {
			"log.Printf"("Migrating database to version %d\n", schema.v)
			_, err := db.Exec(schema.q)
			if err != nil {
				return err
			}

			_, err = db.Exec(`INSERT INTO database_information (schema_version) 
							  VALUES ($1);`, schema.v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
