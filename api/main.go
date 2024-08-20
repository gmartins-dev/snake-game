package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Position struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type GameState struct {
    Snake Position `json:"snake"`
    Apple Position `json:"apple"`
    Score int      `json:"score"`
}

var gameState GameState
var rng *rand.Rand

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/start", startGame).Methods("POST")
    r.HandleFunc("/validate", validateMoves).Methods("POST")
    log.Fatal(http.ListenAndServe(":3001", r))
}

func startGame(w http.ResponseWriter, r *http.Request) {
    gameState = GameState{
        Snake: Position{X: 0, Y: 0},
        Apple: Position{X: rng.Intn(10), Y: rng.Intn(10)},
        Score: 0,
    }
    json.NewEncoder(w).Encode(gameState)
}

func validateMoves(w http.ResponseWriter, r *http.Request) {
    var moves []Position
    json.NewDecoder(r.Body).Decode(&moves)

    for _, move := range moves {
        gameState.Snake.X += move.X
        gameState.Snake.Y += move.Y

        if gameState.Snake.X < 0 || gameState.Snake.X >= 10 || gameState.Snake.Y < 0 || gameState.Snake.Y >= 10 {
            http.Error(w, "Game Over", http.StatusBadRequest)
            return
        }

        if gameState.Snake == gameState.Apple {
            gameState.Score++
            gameState.Apple = Position{X: rng.Intn(10), Y: rng.Intn(10)}
        }
    }

    json.NewEncoder(w).Encode(gameState)
}

func init() {
    rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}
