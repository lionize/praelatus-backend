package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/praelatus/backend/models"
)

func InitUserRoutes() {
	BaseRoutes.Users.Handle("/", AdminRequired(ListUsers, true, true)).Methods("GET")
	BaseRoutes.Users.Handle("/", NoAuth(CreateUser, true, true)).Methods("POST")
	BaseRoutes.Users.Handle("/{name}", Authentication(GetUser, true, true)).Methods("GET")
	BaseRoutes.Users.Handle("/{name}", Authentication(UpdateUser, true, true)).Methods("POST", "PUT")
}

func ListUsers(c *Context) error {
	users, err := as.Store.Users().GetAll()
	if err != nil {
		// TODO: Don't return raw terrible stuff.
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c *Context) error {
	var u models.User

	json.Unmarshal(c.Body)
	err := as.Store.Users().Save(&u)
	// TODO: Handle errors properly (i.e. username taken etc.)
	if err != nil {
		return c.String(http.StatusInternalServerError, dberr.Error())
	}

	return http.StatusOK, &user
}

func LoginUser(c *Context) (int, []byte) {
	return http.StatusNotImplemented, []byte("Not implemented")
}

func GetUser(c *Context) (int, []byte) {
	var u models.User
	var err error

	username := c.Vars.String("username")

	// If we get an integer search based on ID instead of username
	if id, err := strconv.Atoi(username); err == nil {
		u, err = Store.Users().Get(id)
	} else {
		u, err = Store.Users().GetByUsername(username)
	}

	// TODO: Properly return not found when appropriate
	if err != nil {
		// TODO: Don't return raw terrible stuff.
		return http.StatusInternalServerError, []byte(err.Error())
	}

	ujson, err := json.Marshal(&u)

	return http.StatusOK, ujson
}

// TODO: implement this
func UpdateUser(c *Context) (int, []byte) {
	return http.StatusNotImplemented, []byte("Not implemented")
}
