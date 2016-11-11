package pg_test

import "testing"

func TestUserGet(t *testing.T) {
	u, e := s.Users().Get(1)
	failIfErr("User Get", t, e)

	if u == nil {
		t.Error("Expected a user got nil")
	}
}
