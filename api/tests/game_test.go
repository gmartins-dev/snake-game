package tests

import (
	"api/pkg/handlers"
	"api/pkg/routes"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestStartGame verifies that the /start endpoint returns a 200 OK status.
func TestStartGame(t *testing.T) {
    req, err := http.NewRequest("POST", "/start", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.StartGame)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}

// TestValidateMoves verifies that the /validate endpoint returns a 200 OK status
// when provided with a valid JSON payload.
func TestValidateMoves(t *testing.T) {
    var jsonStr = []byte(`[
        {"x": 1, "y": 0},
        {"x": 0, "y": 1}
    ]`)
    req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router := routes.SetupRouter()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}

// TestValidateMovesInvalidJSON verifies that the /validate endpoint returns a 400 Bad Request status
// when provided with an invalid JSON payload.
func TestValidateMovesInvalidJSON(t *testing.T) {
    var jsonStr = []byte(`invalid json`)
    req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router := routes.SetupRouter()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }
}

// TestValidateMovesMissingFields verifies that the /validate endpoint returns a 400 Bad Request status
// when provided with a JSON payload missing required fields.
func TestValidateMovesMissingFields(t *testing.T) {
    var jsonStr = []byte(`[
        {"x": 1}
    ]`)
    req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router := routes.SetupRouter()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }
}

// TestEndGame verifies that the /end endpoint returns a 200 OK status.
func TestEndGame(t *testing.T) {
    req, err := http.NewRequest("POST", "/end", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.EndGame)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
