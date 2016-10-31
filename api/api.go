package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

type Routes struct {
	Root     *mux.Router
	Users    *mux.Router
	Projects *mux.Router
	Tickets  *mux.Router
}

var BaseRoutes *Routes
var Store store.Store
var Cache *store.Cache

func Run(port string) {
	Store = pg.New(os.Getenv("PRAELATUS_DB"))

	BaseRoutes = &Routes{}
	BaseRoutes.Root = mux.NewRouter()
	BaseRoutes.Users = BaseRoutes.Root.PathPrefix("/users").Subrouter()
	BaseRoutes.Projects = BaseRoutes.Root.PathPrefix("/projects").Subrouter()
	BaseRoutes.Tickets = BaseRoutes.Root.PathPrefix("/tickets").Subrouter()

	InitUserRoutes()
	InitProjectRoutes()
	InitTicketRoutes()

	http.ListenAndServe(port, BaseRoutes.Root)
}
