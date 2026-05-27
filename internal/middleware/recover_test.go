package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

func TestRecover(t *testing.T) {
	tests := []struct {
		expectedStatus int
		expectedBody   string
		handler        http.HandlerFunc
		name           string
	}{
		{
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "An internal server error occurred.",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("Could not parse JSON.")
			}),
			name: "Panic",
		},
		{
			expectedStatus: http.StatusOK,
			expectedBody:   `Ok`,
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Ok"))
			}),
			name: "Success",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/", nil)

			mid := New(*handlers.New(*jsonlog.New(os.Stdout, jsonlog.LevelInfo)), *jsonlog.New(os.Stdout, jsonlog.LevelInfo))

			mid.Recover(tt.handler).ServeHTTP(recorder, request)

			if recorder.Code != tt.expectedStatus {
				t.Fatalf("Status: %v != expected: %v", recorder.Code, tt.expectedStatus)
			}

			if !strings.Contains(recorder.Body.String(), tt.expectedBody) {
				t.Fatalf("Body: %v != expected: %v", recorder.Body.String(), tt.expectedBody)
			}
		})
	}
}
