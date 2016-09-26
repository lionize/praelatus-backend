package api

import "github.com/gorilla/mux"

type Routes struct {
	Root     *mux.Router
	ApiRoot  *mux.Router
	Users    *mux.Router
	Projects *mux.Router
	Tickets  *mux.Router
}

var BaseRoutes *Routes

func BuildRoutes() {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = Srv.Router
	BaseRoutes.ApiRoot = Srv.Router.PathPrefix("/api/v1").Subrouter()
	BaseRoutes.Users = BaseRoutes.ApiRoot.PathPrefix("/users").Subrouter()
	BaseRoutes.Projects = BaseRoutes.ApiRoot.PathPrefix("/projects").Subrouter()

	InitUserRoutes()
	InitProjectRoutes()
	InitTicketRoutes()

			projects.GET("", s.ListProjects)
			// 				projects.POST("/projects/create", s.CreateProject)

			projects.GET("/:key", s.GetProject)
			// projects.PUT("/projects/:key", s.UpdateProject)
			// 				projects.GET("/projects/{key}/members/add", s.AddProjectMember)
			// 				projects.GET("/projects/{key}/members", s.ListProjectsMembers)
			// 				projects.GET("/projects/{key}/tickets", s.ListProjectsTickets)

		tickets := v1.Group("/tickets")
		{
			tickets.POST("", s.ListTickets)
			tickets.POST("/create", s.CreateTicket)

			tickets.GET("/:key", s.GetTicket)
			tickets.PUT("/:key", s.UpdateTicket)

			tickets.POST("/:key/advance", s.AdvanceTicket)
			tickets.POST("/:key/retract", s.RetractTicket)
		}
	}
}

