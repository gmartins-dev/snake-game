package handlers

import (
	"api/pkg/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var gameState models.GameState
var rng *rand.Rand

func init() {
    rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// StartGame initializes the game state and starts a new game
func StartGame(w http.ResponseWriter, r *http.Request) {
    gameState = models.GameState{
        Snake: models.Position{X: 0, Y: 0},
        Apple: models.Position{X: rng.Intn(10), Y: rng.Intn(10)},
        Score: 0,
    }
    json.NewEncoder(w).Encode(gameState)
}

// ValidateMoves processes the moves and updates the game state
func ValidateMoves(w http.ResponseWriter, r *http.Request) {
    var moves []models.Position
    if err := json.NewDecoder(r.Body).Decode(&moves); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    for _, move := range moves {
        gameState.Snake.X += move.X
        gameState.Snake.Y += move.Y

        // Check if the snake has moved out of bounds
        if gameState.Snake.X < 0 || gameState.Snake.X >= 10 || gameState.Snake.Y < 0 || gameState.Snake.Y >= 10 {
            http.Error(w, "Game Over", http.StatusBadRequest)
            return
        }

        // Check if the snake has eaten the apple
        if gameState.Snake == gameState.Apple {
            gameState.Score++
            gameState.Apple = models.Position{X: rng.Intn(10), Y: rng.Intn(10)}
        }
    }

    json.NewEncoder(w).Encode(gameState)
}

// EndGame resets the game state and ends the game
func EndGame(w http.ResponseWriter, r *http.Request) {
    gameState = models.GameState{} // Reset the game state
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Game ended"))
}
