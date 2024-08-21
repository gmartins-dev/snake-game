# Snake Game API Tutorial

## Overview

This Snake game API allows you to start a new game and validate moves. The game involves a snake that moves around a grid to eat an apple. The game state is managed by the server, ensuring fair play and preventing cheating.

## Game Rules

1. The snake starts at position (0, 0) and moves to the right initially.
2. The snake will never grow longer than length 1.
3. There is only one apple on the game board at a time.
4. If the snake hits the edge of the game bounds, the game is over.
5. The game state is validated by the API to ensure no cheating.

## Game State

The game state includes:
- The position of the snake.
- The position of the apple.
- The current score.

## API Endpoints

### 1. Start a New Game

**Endpoint**: `POST /start`

**Request Body** (optional):
{
  "width": 10,
  "height": 10
}

**Response**:
{
  "snake": { "x": 0, "y": 0 },
  "apple": { "x": <random_x>, "y": <random_y> },
  "score": 0
}

### 2. Validate Moves

**Endpoint**: `POST /validate`

**Request Body**:
[
  { "x": 1, "y": 0 },
  { "x": 0, "y": 1 }
]

**Responses**:
- **200 OK**:
  {
    "snake": { "x": <new_x>, "y": <new_y> },
    "apple": { "x": <random_x>, "y": <random_y> },
    "score": <new_score>
  }
- **400 Bad Request**:
  {
    "error": "Game over. Snake hit the wall or invalid move."
  }

## Example Workflow

### 1. Start a New Game

**Request**:
curl -X POST http://localhost:3001/start -H "Content-Type: application/json" -d '{"width": 10, "height": 10}'

**Response**:
{
  "snake": { "x": 0, "y": 0 },
  "apple": { "x": 5, "y": 5 },
  "score": 0
}

### 2. Validate Moves

**Request**:
curl -X POST http://localhost:3001/validate -H "Content-Type: application/json" -d '[{"x": 1, "y": 0}, {"x": 1, "y": 0}, {"x": 0, "y": 1}]'

**Response**:
{
  "snake": { "x": 2, "y": 1 },
  "apple": { "x": 3, "y": 3 },
  "score": 1
}

## Implementation Details

### main.go

The main.go file initializes the server and sets up the routes.

package main

import (
    "api/config"
    "api/pkg/routes"
    "log"
    "net/http"
)

func main() {
    config.LoadConfig()
    r := routes.SetupRouter()
    log.Fatal(http.ListenAndServe(config.ServerAddress, r))
}

### Game State

The game state is managed by the `GameState` struct, which includes the snake's position, the apple's position, and the current score.

type GameState struct {
    Snake Position `json:"snake"`
    Apple Position `json:"apple"`
    Score int      `json:"score"`
}

### Position Struct

Represents the coordinates (x, y) on the game board.

type Position struct {
    X int `json:"x"`
    Y int `json:"y"`
}

This tutorial provides a basic understanding of how the Snake game API works, including the game rules, API endpoints, and example workflows.
