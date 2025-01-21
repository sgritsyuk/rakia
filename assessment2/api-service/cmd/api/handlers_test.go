package main

import (
	"api-service/internal/domain"
	"api-service/internal/store"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
)

const testId int = 42

var testPost = domain.Post{
	Title:   "Title 123",
	Content: "Ipsum non tempora magnam neque tempora",
	Author:  "Author 456",
}

type handlersFixture struct {
	ctx   context.Context
	store *store.MockPostStore
}

func newHandlersFixture(t *testing.T) *handlersFixture {
	t.Helper()

	const timeout = 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)

	ctrl := gomock.NewController(t)
	store := store.NewMockPostStore(ctrl)

	return &handlersFixture{
		ctx:   ctx,
		store: store,
	}
}

func newTestApp(fixture *handlersFixture) *App {
	return &App{
		PostStore: fixture.store,
	}
}

// TestHandlers_PostsGet tests PostsGet endpoint
func TestHandlers_PostsGet(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fixture := newHandlersFixture(t)
		app := newTestApp(fixture)

		posts := []domain.Post{testPost}

		fixture.store.EXPECT().
			Get(gomock.Any(), "", 1, 2).
			Return(&posts, nil)

		req, _ := http.NewRequest("GET", "/v1/posts/?page=1&limit=2", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.PostsGetHandler)
		handler.ServeHTTP(rr, req)

		jsonPosts, _ := json.Marshal(posts)
		expectedBody := fmt.Sprintf("{\"error\":false,\"message\":\"\",\"data\":%s}", string(jsonPosts))
		if rr.Code != http.StatusOK {
			t.Errorf("expected http.StatusOK, but got %d", rr.Code)
		}
		if rr.Body.String() != expectedBody {
			t.Errorf("incorrect response body, got %s", rr.Body.String())
		}
	})
}

// TestHandlers_PostsGetOne tests PostsGetOne endpoint
func TestHandlers_PostsGetOne(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fixture := newHandlersFixture(t)
		app := newTestApp(fixture)

		fixture.store.EXPECT().
			GetOne(gomock.Any(), testId).
			Return(&testPost, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", fmt.Sprintf("%d", testId))
		req, _ := http.NewRequest("GET", "/v1/posts/{id}", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.PostsGetOneHandler)
		handler.ServeHTTP(rr, req)

		jsonPost, _ := json.Marshal(testPost)
		expectedBody := fmt.Sprintf("{\"error\":false,\"message\":\"\",\"data\":%s}", string(jsonPost))
		if rr.Code != http.StatusOK {
			t.Errorf("expected http.StatusOK, but got %d", rr.Code)
		}
		if rr.Body.String() != expectedBody {
			t.Errorf("incorrect response body, got %s", rr.Body.String())
		}
	})
}

// TestHandlers_PostsAdd tests PostsAdd endpoint
func TestHandlers_PostsAdd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fixture := newHandlersFixture(t)
		app := newTestApp(fixture)

		postBody := map[string]interface{}{
			"title":   testPost.Title,
			"content": testPost.Content,
			"author":  testPost.Author,
		}

		fixture.store.EXPECT().
			Insert(gomock.Any(), testPost).
			Return(testId, nil)

		body, _ := json.Marshal(postBody)
		req, _ := http.NewRequest("POST", "/v1/posts", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.PostsAddHandler)
		handler.ServeHTTP(rr, req)

		expectedBody := fmt.Sprintf("{\"error\":false,\"message\":\"post added\",\"data\":%d}", testId)
		if rr.Code != http.StatusOK {
			t.Errorf("expected http.StatusOK, but got %d", rr.Code)
		}
		if rr.Body.String() != expectedBody {
			t.Errorf("incorrect response body, got %s", rr.Body.String())
		}
	})
}

// TestHandlers_PostsUpdate tests PostsUpdate endpoint
func TestHandlers_PostsUpdate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fixture := newHandlersFixture(t)
		app := newTestApp(fixture)

		postBody := map[string]interface{}{
			"title":   testPost.Title,
			"content": testPost.Content,
			"author":  testPost.Author,
		}

		fixture.store.EXPECT().
			Update(gomock.Any(), testId, testPost).
			Return(nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", fmt.Sprintf("%d", testId))
		body, _ := json.Marshal(postBody)
		req, _ := http.NewRequest("PUT", "/v1/posts/{id}", bytes.NewReader(body))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.PostsUpdateHandler)
		handler.ServeHTTP(rr, req)

		expectedBody := "{\"error\":false,\"message\":\"post updated\"}"
		if rr.Code != http.StatusOK {
			t.Errorf("expected http.StatusOK, but got %d", rr.Code)
		}
		if rr.Body.String() != expectedBody {
			t.Errorf("incorrect response body, got %s", rr.Body.String())
		}
	})
}

// TestHandlers_PostsDelete tests PostsDelete endpoint
func TestHandlers_PostsDelete(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fixture := newHandlersFixture(t)
		app := newTestApp(fixture)

		fixture.store.EXPECT().
			Delete(gomock.Any(), testId).
			Return(nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", fmt.Sprintf("%d", testId))
		req, _ := http.NewRequest("DELETE", "/v1/posts/{id}", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.PostsDeleteHandler)
		handler.ServeHTTP(rr, req)

		expectedBody := "{\"error\":false,\"message\":\"post deleted\"}"
		if rr.Code != http.StatusOK {
			t.Errorf("expected http.StatusOK, but got %d", rr.Code)
		}
		if rr.Body.String() != expectedBody {
			t.Errorf("incorrect response body, got %s", rr.Body.String())
		}
	})
}
