package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/iamthemuffinman/logsip"
	"github.com/praelatus/backend/models"
)

// Context holds request scoped information and provides utility
// methods for accessing request data.
type Context struct {
	Val  map[string]interface{}
	vars map[string]string
	R    *http.Request
	Err  error
}

// CurrentUser will return the current user for this Context else will return
// nil
func (c *Context) CurrentUser() *models.User {
	if u, ok := c.Val["CurrentUser"].(*models.User); ok {
		return u
	}

	return nil
}

// Body returns the body of the request as a []byte and an error indicating any
// read errors. This is just a convenience function.
func (c *Context) Body() ([]byte, error) {
	return ioutil.ReadAll(c.R.Body)
}

// JSON will unmarshal the body of the request into the interface m
func (c *Context) JSON(m interface{}) error {
	b, err := c.Body()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, m)
	return err
}

// String will return the context value at key as a string if possible,
// returns "" if an error occurs.
func (c *Context) String(key string) string {
	v, ok := c.Val[key].(string)
	if ok {
		return v
	}

	log.Errorf("Failed to retrieve string value at: %s Actual value: %v\n", key, v)
	return ""
}

// Var will return the url variable stored at key
func (c *Context) Var(key string) string {
	if c.vars == nil {
		c.vars = mux.Vars(c.R)
	}

	return c.vars[key]
}
