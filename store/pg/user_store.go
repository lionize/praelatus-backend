package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// UserStore contains methods for storing and retrieving Users from a Postgres
// DB
type UserStore struct {
	db *sqlx.DB
}

// Get retrieves the user by row id
func (s *UserStore) Get(id int) (*models.User, error) {
	var u models.User
	err := s.db.QueryRowx("SELECT * FROM users WHERE id = $1;", id).
		StructScan(&u)
	return &u, err
}

// GetByUsername will retrieve the user by the given username.
func (s *UserStore) GetByUsername(un string) (*models.User, error) {
	var u models.User
	err := s.db.QueryRowx("SELECT * FROM users WHERE username = $1;", un).
		StructScan(&u)
	return &u, err
}

// GetAll retrieves all users from the database.
func (s *UserStore) GetAll() ([]models.User, error) {
	users := []models.User{}
	rows, err := s.db.Queryx("SELECT * FROM users;")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var u models.User

		err := rows.StructScan(&u)
		if err != nil {
			return users, err
		}

		users = append(users, u)
	}

	return users, nil
}

// Save will update the given user into the database.
func (s *UserStore) Save(u *models.User) error {
	if u.Password == "" {
		_, err := s.db.Exec(`UPDATE users SET 
		(username, email, full_name, is_admin) = (?, ?, ?, ?) WHERE id = ?;`,
			u.Username, u.Email, u.FullName, u.IsAdmin, u.ID)

		return err
	}

	_, err := s.db.Exec(`UPDATE users SET 
		(username, password, email, full_name, is_admin) = (?, ?, ?, ?) 
		WHERE id = ?;`,
		u.Username, u.Password, u.Email, u.FullName, u.IsAdmin, u.ID)

	return err
}

// New will create the user in the database
func (s *UserStore) New(u *models.User) error {
	id, err := s.db.Exec(`INSERT INTO users
		(username, password, email, full_name, is_admin) VALUES ($1, $2, $3, $4);`,
		u.Username, u.Password, u.Email, u.FullName, u.IsAdmin)
	if err != nil {
		return err
	}

	u.ID, err = id.LastInsertId()
	return err
}
