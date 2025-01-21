package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WebServer allows to serve data in JSON format
type WebServer struct {
	Port    string
	Timeout time.Duration
}

// JsonResponse represent typical webserver response
type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// NewWebServer creates a new webserver
func NewWebServer(port string) WebServer {
	return WebServer{
		Port: port,
	}
}

// ReadJSON reads JSON data from request body
func (srv *WebServer) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// decode request body
	maxBytes := 1048576 // one Mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// WriteJSON writes JSON data as a response
func (srv *WebServer) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	// encode JSON data
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// provide headers
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// ErrorJSON writes JSON error data as a response
func (srv *WebServer) ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	// process status code, by default BadRequest
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	// prepare and send error response
	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()
	_ = srv.WriteJSON(w, statusCode, payload)
}

// Serve listens and serves data for webserver
func (srv *WebServer) Serve(routes http.Handler) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", srv.Port),
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
