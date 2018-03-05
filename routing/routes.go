package routing

import (
	"net/http"

	"github.com/managef/api/handlers"
)

// Route describes a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds an array of Route
type Routes struct {
	Routes []Route
}

// NewRoutes creates and returns all the API routes
func NewRoutes() (r *Routes) {
	r = new(Routes)

	r.Routes = []Route{
		{
			"Root",
			"GET",
			"/api",
			handlers.GetApi,
		},
		{
			"ServiceList",
			"GET",
			"/api/job/{id}",
			handlers.GetJob,
		},
	}

	return
}
