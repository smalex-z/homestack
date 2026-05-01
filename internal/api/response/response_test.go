package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"homestack/internal/api/response"
)

func TestSuccessAndCreatedSetEnvelope(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		write      func(http.ResponseWriter)
		wantStatus int
	}{
		{
			name:       "Success",
			write:      func(w http.ResponseWriter) { response.Success(w, map[string]int{"n": 1}) },
			wantStatus: http.StatusOK,
		},
		{
			name:       "Created",
			write:      func(w http.ResponseWriter) { response.Created(w, map[string]string{"id": "abc"}) },
			wantStatus: http.StatusCreated,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			tc.write(rec)
			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("Content-Type = %q, want application/json", ct)
			}
			var env response.Response
			if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			if !env.Success {
				t.Errorf("env.Success = false, want true (body: %s)", rec.Body.String())
			}
			if env.Error != "" {
				t.Errorf("env.Error = %q, want empty on success", env.Error)
			}
			if env.Data == nil {
				t.Errorf("env.Data is nil, want populated")
			}
		})
	}
}

func TestErrorHelpersMapToStatuses(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		write      func(http.ResponseWriter)
		wantStatus int
		wantMsg    string
	}{
		{"BadRequest", func(w http.ResponseWriter) { response.BadRequest(w, "bad") }, http.StatusBadRequest, "bad"},
		{"NotFound", func(w http.ResponseWriter) { response.NotFound(w, "missing") }, http.StatusNotFound, "missing"},
		{"Conflict", func(w http.ResponseWriter) { response.Conflict(w, "dup") }, http.StatusConflict, "dup"},
		{"InternalError", func(w http.ResponseWriter) { response.InternalError(w, "boom") }, http.StatusInternalServerError, "boom"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			tc.write(rec)
			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			var env response.Response
			if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if env.Success {
				t.Errorf("env.Success = true, want false on error")
			}
			if env.Error != tc.wantMsg {
				t.Errorf("env.Error = %q, want %q", env.Error, tc.wantMsg)
			}
		})
	}
}

func TestNoContentSetsStatusOnly(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	response.NoContent(rec)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
	if rec.Body.Len() != 0 {
		t.Errorf("body = %q, want empty", rec.Body.String())
	}
}
