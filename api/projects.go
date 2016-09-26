package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitProjectRoutes() {
	BaseRoutes.Projects.Handle("/", Authentication(ListProjects, true, true)).Methods("GET")
	BaseRoutes.Projects.Handle("/{key}", Authentication(GetProject, true, true)).Methods("GET")
}

func ListProjects(c echo.Context) error {
	projects, err := s.Store.FindAllProjects()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, projects)
}

func (gc *GlobalContext) GetProject(c echo.Context) error {
	// TODO: Sanitize this somehow? Not sure how this could really be used
	// maliciously....
	key := c.Param("key")
	p, err := s.Store.Project().Get(key)
	if err != nil {
		// TODO: Don't return raw terrible stuff.
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if p.Key == "" {
		return c.String(http.StatusNotFound, "Project does not exist.")
	}

	return c.JSON(http.StatusOK, &p)
}
