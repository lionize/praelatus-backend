package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/chasinglogic/tessera/store"
	"github.com/gorilla/mux"
)

var Log *logrus.Logger
var Srv *ApiServer

// ApiServer holds items which need to be globally accessible.
// such as DB connections and loggers.
type ApiServer struct {
	Store  store.Store
	Cache  store.Cache
	Routes *mux.Router
}

func New(log *logrus.Logger) {
	Srv = &ApiServer{}
	Srv.Store = store.NewSqlStore()
	Srv.Cache = store.NewRedisCache()
	Srv.Router = mux.NewRouter()

	BuildRoutes()

	Log = log
}
