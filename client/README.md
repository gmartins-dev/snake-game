# Snake Game Frontend Documentation

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

The application will be available at [http://localhost:5173](http://localhost:5173).

## Project Structure

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
