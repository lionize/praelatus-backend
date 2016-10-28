package migrations

import "github.com/praelatus/backend/store"

const (
	v1 = iota
)

func checkMigration(e error) {
	if e != nil {
		panic(e)
	}
}

// RunMigration will run all database migrations depending on the version
// returned from the database_information table.
func RunMigration(s store.SQLStore) error {
	version := 0

	rws, err := s.RunQuery("SELECT schema_version FROM database_information;")
	if err == nil {
		rws.Scan(&version)
	}

	if version < v1 {
		_, err = s.RunQuery(v1Schema)
		checkMigration(err)
		_, err = s.RunQuery("INSERT INTO database_information VALUES (" + string(v1) + ")")
		checkMigration(err)
	}

	return err
}
