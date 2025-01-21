package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

type RouteDescriptor struct {
	Method string
	Path   string
}

func Test_routes_exists(t *testing.T) {
	testApp := App{}

	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	routes := []RouteDescriptor{
		{
			Method: "GET",
			Path:   "/v1/posts",
		},
		{
			Method: "GET",
			Path:   "/v1/posts/{id}",
		},
		{
			Method: "POST",
			Path:   "/v1/posts",
		},
		{
			Method: "PUT",
			Path:   "/v1/posts/{id}",
		},
		{
			Method: "DELETE",
			Path:   "/v1/posts/{id}",
		},
	}

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route RouteDescriptor) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route.Path == foundRoute && route.Method == method {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find %s: %s in registered routes", route.Method, route.Path)
	}
}
