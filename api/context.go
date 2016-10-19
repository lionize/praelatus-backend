package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/praelatus/backend/models"
	"github.com/gorilla/mux"
)

type Context struct {
	CurrentUser *models.User
	Vars        map[string]interface{}
}

func (c *Context) Unauthenticated() bool {
	if c.CurrentUser == nil {
		return true
	}

	return false
}

type AppHandler func(*Context) (int, []byte)

func AdminRequired(fn AppHandler) http.Handler {
	return &handler{fn, true, true}
}

func AuthRequired(fn AppHandler) http.Handler {
    return &handler{fn, true, false}
}

func NoAuth(fn AppHandler) http.Handler {
    return &handler{fn, false, false}
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
	token, _ := json.Unmarshal(Srv.Cache.Get(token), c.CurrentUser)

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
