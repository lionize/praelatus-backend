package migrations

import "github.com/praelatus/backend/store"

const (
	v1 = iota
	v2
)

func checkMigration(e error) {
	if e != nil {
		panic(e)
	}
}

// RunMigrations will run all database migrations depending on the version
// returned from the database_information table.
func RunMigrations(s store.SQLStore) error {
	version := 0

	rws, err := s.RunQuery("SELECT schema_version FROM database_information;")
	if err == nil {
		rws.Scan(&version)
	}

	// TODO revisit this to make it a little cleaner (i.e. not infinite if
	// statements)
	if version < v1 {
		_, err = s.RunQuery(v1Schema)
		checkMigration(err)
		_, err = s.RunQuery("INSERT INTO database_information VALUES (" + string(v1) + ")")
		checkMigration(err)
	}

	if version < v2 {
		_, err = s.RunQuery(v2Schema)
		checkMigration(err)
		_, err = s.RunQuery("INSERT INTO database_information VALUES (" + string(v2) + ")")
		checkMigration(err)
	}

	return err
}
