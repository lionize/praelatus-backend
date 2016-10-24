package api

import (
	"encoding/json"
	"net/http"
)

func InitProjectRoutes() {
	BaseRoutes.Projects.Handle("/", Authentication(ListProjects, true, true)).Methods("GET")
	BaseRoutes.Projects.Handle("/{key}", Authentication(GetProject, true, true)).Methods("GET")
}

func ListProjects(c *Context) (int, []byte) {
	projects, err := s.Store.FindAllProjects()
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	pjson, err := json.Marshal(projects)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	return http.StatusOK, pjson
}

func GetProject(c *Context) (int, []byte) {
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
