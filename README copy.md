# Snake Game

Simply version using React and Golang of the nostalgic snake game from the old Nokia phones

# Snake Game (API Documentation)

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
    },
    {
      "name": "End Game",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": ""
        },
        "url": {
          "raw": "http://localhost:3001/end",
          "protocol": "http",
          "host": ["localhost"],
          "port": "3001",
          "path": ["end"]
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

==================================================================

# Snake Game (Frontend Documentation)

Simply version using React and Golang of the nostalgic snake game from the old Nokia phones

## Overview

This is the frontend application for the Snake Game, built using React. It communicates with the backend API to manage the game state.

## Setup Instructions

### Prerequisites

- Node.js (version 14.x or higher)
- npm (version 6.x or higher) or yarn (version 1.x or higher)

### Installation

1. Clone the repository:

```shell
git clone <repository-url>
```

2. Install the dependencies:

```shell
npm install
```

or

```shell
yarn install
```

3. Start the development server:

```shell
npm start
```

or

```shell
yarn start
```

The application will be available at [http://localhost:3000](http://localhost:3000).

## Project Structure

snake-game/client
├── public/
│ ├── index.html
│ └── ...
├── src/
│ ├── components/
│ │ ├── Apple.tsx
│ │ ├── GameBoard.tsx
│ │ ├── Snake.tsx
│ │ └── ...
│ ├── App.tsx
│ ├── index.tsx
│ └── ...
├── package.json
├── tsconfig.json
└── ...

### Main Components

#### App.tsx

The root component of the application. It sets up the main layout and includes the GameBoard component.

#### GameBoard.tsx

Manages the game state and renders the game board. Communicates with the backend API to start the game, validate moves, and end the game.

#### Snake.tsx

Renders the snake on the game board based on its current position.

#### Apple.tsx

Renders the apple on the game board based on its current position.

### API Integration

The frontend communicates with the backend API to manage the game state. The main API endpoints are:

- `POST /api/start`: Starts a new game and initializes the game state.
- `POST /api/validate`: Validates the moves and updates the game state.
- `POST /api/end`: Ends the game and resets the game state.

## Conclusion

This documentation provides an overview of the Snake Game frontend application, including setup instructions, project structure, main components, and API integration. For more details, refer to the source code and comments within the code.
