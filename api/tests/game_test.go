package tests

import (
	"api/pkg/handlers"
	"api/pkg/routes"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
