package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/handlers"
	usermodel "github.com/dositadi/groupie-tracker/internal/models/user_model"
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
				panic("An internal server error occurred.")
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
			h := handlers.New(*logger, &usermodel.UserModel{}, nil, nil, artistapi.ArtistInfo{}, *groupietracker.New())
			mid := New(*h, *logger)

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
