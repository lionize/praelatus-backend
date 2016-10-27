package pg

import (
	log "github.com/iamthemuffinman/logsip"
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg/migrations"
)

// Store implements the store.Store interface for a postgres DB.
type Store struct {
	db       *sqlx.DB
	replicas *[]sqlx.DB
	users    *pgUserStore
	projects *pgProjectStore
	tickets  *pgTicketStore
}

// TODO implement interfaces
type pgProjectStore struct{}
type pgUserStore struct{}

// New connects to the postgres database provided and returns a store
// that's connected.
func New(conn string, replicas ...string) store.Store {
	// TODO: replica support
	s := &Store{}

	d, err := sqlx.Open("postgres", conn)
	if err != nil {
		log.Panicln(err)
	}

	s.db = d

	migrations.RunMigration(s)

	return s
}

// Users returns the underlying UserStore for a postgres DB
// TODO fix structs to implement interface
func (pg *Store) Users() store.UserStore {
	// return pg.users
	return nil
}

// Projects returns the underlying ProjectStore for a postgres DB
func (pg *Store) Projects() store.ProjectStore {
	// return pg.projects
	return nil
}

// Tickets returns the underlying TicketStore for a postgres DB
func (pg *Store) Tickets() store.TicketStore {
	// return pg.tickets
	return nil
}

// RunQuery runs the query specified by q and returns the rows and any errors
// encountered.
func (pg *Store) RunQuery(q string) (*sqlx.Rows, error) {
	return pg.db.Queryx(q)
}

// SchemaVersion returns the current schema version of the database.
func (pg *Store) SchemaVersion() int {
	var v int

	rw, err := pg.RunQuery("SELECT schema_version FROM database_information")
	if err != nil {
		return 0
	}

	rw.Scan(&v)
	return v
}
