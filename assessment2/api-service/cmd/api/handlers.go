package main

import (
	"api-service/internal/domain"
	"api-service/internal/server"
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type JsonPostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// DefaultPage and DefaultLimit are default pagination parameters
const DefaultPage = 1
const DefaultLimit = 5

// PostsGetHandler is an endpoint handler for posts list
func (app *App) PostsGetHandler(w http.ResponseWriter, r *http.Request) {
	// read and parse query parameters
	titleParam := r.URL.Query().Get("title")
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	page, err := strconv.ParseInt(pageParam, 10, 32)
	if err != nil || page < 1 {
		page = DefaultPage
	}
	limit, err := strconv.ParseInt(limitParam, 10, 32)
	if err != nil || limit < 1 {
		limit = DefaultLimit
	}

	// fetch posts from store
	ctx, cancel := context.WithTimeout(context.Background(), app.WebServer.Timeout*time.Second)
	defer cancel()
	posts, err := app.PostStore.Get(ctx, titleParam, int(page), int(limit))
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// return successful json response with list of posts
	response := server.JsonResponse{
		Error:   false,
		Message: "",
		Data:    posts,
	}
	_ = app.WebServer.WriteJSON(w, http.StatusOK, response)
}

// PostsGetOneHandler is an endpoint handler for specific post
func (app *App) PostsGetOneHandler(w http.ResponseWriter, r *http.Request) {
	// get post id from URL params
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// fetch posts from store
	ctx, cancel := context.WithTimeout(context.Background(), app.WebServer.Timeout*time.Second)
	defer cancel()
	post, err := app.PostStore.GetOne(ctx, int(id))
	if err != nil {
		if errors.Is(err, domain.ErrorPostNotFound) {
			app.WebServer.ErrorJSON(w, err, http.StatusNotFound)
			return
		}
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// return successful json response with list of posts
	response := server.JsonResponse{
		Error:   false,
		Message: "",
		Data:    post,
	}
	_ = app.WebServer.WriteJSON(w, http.StatusOK, response)
}

// PostsAddHandler is an endpoint handler for add new post
func (app *App) PostsAddHandler(w http.ResponseWriter, r *http.Request) {
	// read json input
	var jsonPayload JsonPostPayload
	err := app.WebServer.ReadJSON(w, r, &jsonPayload)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// construct domain object from input data
	post := domain.Post{
		Title:   jsonPayload.Title,
		Content: jsonPayload.Content,
		Author:  jsonPayload.Author,
	}

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), app.WebServer.Timeout*time.Second)
	defer cancel()

	// save post document to store
	id, err := app.PostStore.Insert(ctx, post)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// return successful json response
	response := server.JsonResponse{
		Error:   false,
		Message: "post added",
		Data:    id,
	}
	_ = app.WebServer.WriteJSON(w, http.StatusOK, response)
}

// PostsUpdateHandler is an endpoint handler for update existing post
func (app *App) PostsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// get post id from URL params
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// read json input
	var jsonPayload JsonPostPayload
	err = app.WebServer.ReadJSON(w, r, &jsonPayload)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// construct domain object from input data
	post := domain.Post{
		Title:   jsonPayload.Title,
		Content: jsonPayload.Content,
		Author:  jsonPayload.Author,
	}

	// create context with deadline
	ctx, cancel := context.WithTimeout(context.Background(), app.WebServer.Timeout*time.Second)
	defer cancel()

	// update post document in store
	err = app.PostStore.Update(ctx, int(id), post)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// return successful json response
	response := server.JsonResponse{
		Error:   false,
		Message: "post updated",
		Data:    nil,
	}
	_ = app.WebServer.WriteJSON(w, http.StatusOK, response)
}

// PostsDeleteHandler is an endpoint handler for delete existing post
func (app *App) PostsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// get post id from URL params
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), app.WebServer.Timeout*time.Second)
	defer cancel()

	// delete post document from store
	err = app.PostStore.Delete(ctx, int(id))
	if err != nil {
		app.WebServer.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// return successful json response
	response := server.JsonResponse{
		Error:   false,
		Message: "post deleted",
		Data:    nil,
	}
	_ = app.WebServer.WriteJSON(w, http.StatusOK, response)
}
