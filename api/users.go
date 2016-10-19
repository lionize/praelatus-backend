package api

import (
	"net/http"
	"strconv"

	"github.com/chasinglogic/tessera/models"
	"github.com/labstack/echo"
)

func InitUserRoutes() {
	BaseRoutes.Users.Handle("/", AdminRequired(ListUsers, true, true)).Methods("GET")
	BaseRoutes.Users.Handle("/create", (CreateUser, true, true)).Methods("POST")
	BaseRoutes.Users.Handle("/{name}", Authentication(GetUser, true, true)).Methods("GET")
	BaseRoutes.Users.Handle("/{name}", Authentication(UpdateUser, true, true)).Methods("POST", "PUT")
}

func (as *ApiServer) ListUsers(c echo.Context) error {
	users, err := as.Store.Users().GetAll()
	if err != nil {
		// TODO: Don't return raw terrible stuff.
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (as *ApiServer) CreateUser(c echo.Context) error {
	var u models.User

	c.Bind(&u)

	err := as.Store.Users().Save(&u)
	// TODO: Handle errors properly (i.e. username taken etc.)
	if err != nil {
		return c.String(http.StatusInternalServerError, dberr.Error())
	}

	return c.JSON(http.StatusOK, &user)
}

func (as *ApiServer) LoginUser(c echo.Context) error {
	var l models.LoginRequest

	c.Bind(&l)

	user, dberr := l.Login(gc.db)
	// TODO: Handle errors properly (i.e. bad password)
	if dberr != nil {
		gc.log.Error("Database error logging in")
		return c.String(http.StatusInternalServerError,
			"Username or password was invalid.")
	}

	return c.JSON(http.StatusOK, &user)
}

func (as *ApiServer) GetUser(c echo.Context) error {
	var gerr error
	var u models.User

	// TODO: Sanitize this somehow? Not sure how this could really be used
	// maliciously....
	username := c.Param("name")

	// If we get an integer search based on ID instead of username
	if id, err := strconv.Atoi(username); err == nil {
		u, gerr = as.Store.Users().Get(id)
	} else {
		u, gerr = as.Store.Users().GetByName(username)
	}

	if gerr != nil {
		// TODO: Don't return raw terrible stuff.
		return c.String(http.StatusInternalServerError, gerr.Error())
	}

	if u.Username == "" {
		return c.String(http.StatusNotFound, "User does not exist.")
	}

	// Don't return the password
	u.Password = ""

	return c.JSON(http.StatusOK, &u)
}

// TODO: implement this, updating will suck if it needs to be performant
func (as *ApiServer) UpdateUser(c echo.Context) error {
	var u models.User
	c.Bind(&u)
	return nil
}
