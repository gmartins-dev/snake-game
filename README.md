# Snake Game 🐍

A simplified version of the nostalgic snake game from old Nokia phones, built using React and Golang.

## Backend (API)

### Overview

The Snake Game API allows you to play the classic Snake game through HTTP requests. You can start a game, make moves, and validate them.

### Setup

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

### Running the Server

To start the server, run the following command from the `api/cmd/api` directory:

```
go run main.go
```

The server will start on the address specified in the .env file (default is :3001).

### API Endpoints

#### 1. Start Game

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

#### 2. Validate Moves

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
  - 400 Bad Request: Game over if the snake hits a wall or an invalid move is made. Invalid JSON returns an error.

#### 3. End Game

- **POST /end**
- Description: End the current game and reset the game state.
- Request Body: None
- Response: None

#### Example Postman Collection

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
      "name": "Validate Moves",
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

## Frontend (Client)

### Overview

The Snake Game frontend application is built using React. It communicates with the backend API to manage the game state.

### Setup Instructions

#### Prerequisites

- Node.js (version 14.x or higher)
- npm (version 6.x or higher) or yarn (version 1.x or higher)

#### Installation

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

### Project Structure

```
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
```

#### Main Components

##### App.tsx

The root component of the application. It sets up the main layout and includes the GameBoard component.

##### GameBoard.tsx

Manages the game state and renders the game board. Communicates with the backend API to start the game, validate moves, and end the game.

##### Snake.tsx

Renders the snake on the game board based on its current position.

##### Apple.tsx

Renders the apple on the game board based on its current position.

### API Integration

The frontend communicates with the backend API to manage the game state. The main API endpoints are:

- `POST /api/start`: Starts a new game and initializes the game state.
- `POST /api/validate`: Validates the moves and updates the game state.
- `POST /api/end`: Ends the game and resets the game state.

### Conclusion

This documentation provides an overview of the Snake Game backend (API) and frontend (Client) applications. It includes setup instructions, project structure, main components, and API integration details. For more information, refer to the source code and comments within the code.
