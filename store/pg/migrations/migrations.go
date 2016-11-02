package migrations

import "github.com/praelatus/backend/store"

type schema struct {
	v int
	q string
}

var schemas = []schema{
	v1schema,
}

// RunMigrations will run all database migrations depending on the version
// returned from the database_information table.
func RunMigrations(s store.SQLStore) error {
	for _, schema := range schemas {
		version := s.SchemaVersion()

		if version < schema.v {
			_, err := s.RunExec(schema.q)
			if err != nil {
				return err
			}

			_, err = s.RunExec(`INSERT INTO database_information 
								VALUES (schema_version) = 
								(` + string(schema.v) + `);`)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
