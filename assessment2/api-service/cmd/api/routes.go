package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// ApiVersion defines current version of API
const ApiVersion = "/v1"

func (app *App) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Could be separated from public endpoints to internal http server in future (+metrics)
	mux.Use(middleware.Heartbeat(ApiVersion + "/healthcheck"))

	// Get paginated list of posts endpoint
	mux.Get(ApiVersion+"/posts", app.PostsGetHandler)
	// Get post endpoint
	mux.Get(ApiVersion+"/posts/{id}", app.PostsGetOneHandler)
	// Add a new post endpoint
	mux.Post(ApiVersion+"/posts", app.PostsAddHandler)
	// Update post endpoint
	mux.Put(ApiVersion+"/posts/{id}", app.PostsUpdateHandler)
	// Delete post endpoint
	mux.Delete(ApiVersion+"/posts/{id}", app.PostsDeleteHandler)

	return mux
}
