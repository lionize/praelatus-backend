package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chasinglogic/tessera/models"
	"github.com/gorilla/mux"
)

type Context struct {
	CurrentUser *models.User
	Vars        map[string]string
}

func (c *Context) Unauthenticated() bool {
	if c.CurrentUser == nil {
		return true
	}

	return false
}

type AppHandler func(*Context) (int, []byte)

func Authentication(fn AppHandler, userRequired, adminRequired bool) http.Handler {
	return &h{fn, userRequired, adminRequired}
}

type handler struct {
	fn       AppHandler
	reqUser  bool
	reqAdmin bool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	c := &Context{}
	c.Vars = mux.Vars(r)

	token := ""
	_ := json.Unmarshal(Srv.Cache.Get(token), c.CurrentUser)

	// Commented out for now to make development less of a headache.
	// 	if c.CurrentUser == nil && h.reqUser {
	// 		w.WriteHeader(403)
	// 		w.Write([]byte("You must be logged in to access this resource."))
	// 	}

	// 	if !c.CurrentUser.IsAdmin && h.reqAdmin {
	// 		w.WriteHeader(403)
	// 		w.Write([]byte("You must be an admin to access this resource."))
	// 	}

	statusCode, response := h.fn(c)
	Log.Infof("|%s| [%d] %s %s", r.Method, statusCode, r.URL.Path, time.Since(start).String())

	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	if err != nil {
		Log.Error(err)
	}

}
