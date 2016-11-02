package pg

import (
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func testUsers(s store.Store) error {
	t1, be := models.NewUser("testuser", "test", "Test Testerson",
		"test@example.com", false)
	if be != nil {
		return be
	}

	t2, be := models.NewUser("testadmin", "test", "Test Testerson II",
		"test1@example.com", false)
	if be != nil {
		return be
	}

	users := []models.User{
		*t1,
		*t2,
	}

	for _, u := range users {
		e := s.Users().New(&u)
		if e != nil && e != store.ErrDuplicateEntry {
			return e
		}
	}

	return nil
}
