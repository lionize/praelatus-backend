package api

import (
	"github.com/gorilla/mux"
	"github.com/praelatus/backend/store"
)

var Srv *ApiServer

// ApiServer holds items which need to be globally accessible.
// such as DB connections and loggers.
type ApiServer struct {
	Store  store.Store
	Cache  store.Cache
	Routes *mux.Router
}

func New() {
	Srv = &ApiServer{}
	Srv.Store = store.NewSqlStore()
	Srv.Cache = store.NewRedisCache()
	Srv.Router = mux.NewRouter()

	BuildRoutes()
}
