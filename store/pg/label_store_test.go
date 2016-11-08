package pg

import "testing"

// TODO
func TestGet(t *testing.T) {
	l, e := s.Labels().Get(1)
	failIfErr(t, e)
}
