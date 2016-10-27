package middleware

import (
	"net/http"
	"strings"

	"github.com/praelatus/backend/models"
)

// GetToken will parse the token out of the headers for the given http.Request
func GetToken(r *http.Request) string {
	var tokenStr string

	// Attempt to parse token out of the header
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == "BEARER" {
		// Default session token
		tokenStr = authHeader[7:]
	} else if len(authHeader) > 5 && strings.ToLower(authHeader[0:5]) == "token" {
		// OAuth token
		tokenStr = authHeader[6:]
	}

	return tokenStr
}

// ValidateToken will validate the token with our jwt library and return the
// corresponding user.
// TODO
func ValidateToken(token string) *models.User {
	if token == "" {
		return nil
	}

	return nil
}

// AuthMw verifies that the token is present on the request and is for a valid
// user.
func AuthMw(c *Context) *Context {
	if c.Val["CurrentUser"] != nil {
		return c
	}

	u := ValidateToken(GetToken(c.R))
	if u == nil {
		c.Err = ErrUnauthorized
		return c
	}

	c.Val["CurrentUser"] = u
	return c
}

// AdminMw verifies that the token is present on the request and is for a valid
// user.
func AdminMw(c *Context) *Context {
	c = AuthMw(c)

	if c.CurrentUser() != nil && c.CurrentUser().IsAdmin {
		return c
	}

	c.Err = ErrForbidden
	return c
}
