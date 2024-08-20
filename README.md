# Snake Game

Simply version using React and Golang of the nostalgic snake game from the old Nokia phones

# Snake Game API

A nostalgic implementation of the classic Snake game, designed to be played via an API. This project allows you to start a game, make moves, and validate them through HTTP requests.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Setup](#setup)
- [Running the Server](#running-the-server)
- [API Endpoints](#api-endpoints)
  - [Start Game](#start-game)
  - [Validate Moves](#validate-moves)
- [Example Postman Collection](#example-postman-collection)
- [Notes](#notes)

## Overview

This API allows you to start a game of Snake and validate moves. The game runs on a customizable grid, where the snake can eat apples, grow, and score points.

## Project Structure

```
.gitignore
api/
.env
cmd/
  api/
    main.go
config/
  config.go
docs/
go.mod
go.sum
pkg/
  handlers/
    game.go
  models/
    game.go
routes/
  routes.go
postman/
  api-doc.md
  SnakeGameAPI.postman_collection.json
scripts/
tests/
  game_test.go
NOTES.md
README.md
```

## Setup

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/snake-game.git
   cd snake-game/api
   ```

2. Install dependencies:

   ```
   go mod download
   ```

3. Create a .env file:

   ```
   cp .env.example .env
   ```

4. Update the .env file with your configuration:
   ```
   SERVER_ADDRESS=:3001
   ```

## Running the Server

To start the server, run the following command from the `api/cmd/api` directory:

```
go run main.go
```

The server will start on the address specified in the .env file (default is :3001).

## API Endpoints

### 1. Start Game

- **POST /start**
- Description: Initialize a new game.
- Request Body (optional):
  ```json
  {
    "width": 10,
    "height": 10
  }
  ```
- Response:
  - Returns the initial game state, including grid size, snake position, apple position, and score.

### 2. Validate Moves

- **POST /validate**
- Description: Validate one or more moves for the snake.
- Request Header: Content-Type: application/json
- Request Body:
  ```json
  [
    { "x": 1, "y": 0 },
    { "x": 0, "y": 1 }
  ]
  ```
- Responses:
  - 200 OK: Updated game state with snake position, apple position, and score.
  - 400 Bad Request: Game over if the snake hits a wall or invalid move is made. Invalid JSON returns an error.

## Example Postman Collection

Import the following JSON into Postman to test the API:

```json
{
  "info": {
    "name": "Snake Game API",
    "description": "API for the Snake Game",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Start Game",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": ""
        },
        "url": {
          "raw": "http://localhost:3001/start",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["start"]
        }
      },
      "response": []
    },
    {
      "name": "Validate Moves - Valid Moves",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "[\n    {\"x\": 1, \"y\": 0},\n    {\"x\": 0, \"y\": 1}\n]"
        },
        "url": {
          "raw": "http://localhost:3001/validate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["validate"]
        }
      },
      "response": []
    },
    {
      "name": "Validate Moves - Snake Eats Apple",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "[\n    {\"x\": 1, \"y\": 0},\n    {\"x\": 0, \"y\": 1}\n]"
        },
        "url": {
          "raw": "http://localhost:3001/validate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["validate"]
        }
      },
      "response": []
    },
    {
      "name": "Validate Moves - Snake Hits Wall",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "[\n    {\"x\": -1, \"y\": 0}\n]"
        },
        "url": {
          "raw": "http://localhost:3001/validate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["validate"]
        }
      },
      "response": []
    },
    {
      "name": "Validate Moves - Multiple Moves",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "[\n    {\"x\": 1, \"y\": 0},\n    {\"x\": 1, \"y\": 0},\n    {\"x\": 0, \"y\": 1},\n    {\"x\": 0, \"y\": 1}\n]"
        },
        "url": {
          "raw": "http://localhost:3001/validate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["validate"]
        }
      },
      "response": []
    },
    {
      "name": "Validate Moves - Invalid JSON",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "invalid_json"
        },
        "url": {
          "raw": "http://localhost:3001/validate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["validate"]
        }
      },
      "response": []
    }
  ]
}
```

## Notes

- The game starts with the snake at (0, 0) moving right.
- The grid defaults to 10x10 unless specified.
- The game ends if the snake hits a wall or makes an invalid move.

Enjoy playing the Snake game via API!
