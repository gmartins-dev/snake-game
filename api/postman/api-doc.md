# Snake Game API Documentation

## Overview

This API allows you to start a game of Snake and validate moves. The game runs on a customizable grid, where the snake can eat apples, grow, and score points.

## Base URL

`http://localhost:3001`

## Endpoints

### 1. Start Game

- **POST** `/start`
- **Description**: Initialize a new game.
- **Request Body** (optional):

  ```json
  {
    "width": 10,
    "height": 10
  }
  ```

- **Response**:
  - Returns the initial game state, including grid size, snake position, apple position, and score.

### 2. Validate Moves

- **POST** `/validate`
- **Description**: Validate one or more moves for the snake.
- **Request Header**: `Content-Type: application/json`
- **Request Body**:
  ```json
  [
    { "x": 1, "y": 0 },
    { "x": 0, "y": 1 }
  ]
  ```
- **Responses**:
  - **200 OK**: Updated game state with snake position, apple position, and score.
  - **400 Bad Request**: Game over if the snake hits a wall or invalid move is made. Invalid JSON returns an error.

## Notes

- The game starts with the snake at `(0, 0)` moving right.
- The grid defaults to `10x10` unless specified.
- The game ends if the snake hits a wall or makes an invalid move.

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
