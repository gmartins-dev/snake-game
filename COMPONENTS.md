Overview of the key components and their roles:
This structure adheres to the standard Go project layout and ensures that your code is well-organized and maintainable.

## Main Application

- `api/cmd/api/main.go`: The entry point of your application. It loads the configuration and sets up the router.

## Configuration

- `api/config/config.go`: Handles loading the configuration from the .env file.

## Routing

- `api/pkg/routes/routes.go`: Sets up the HTTP routes for the application.

## Handlers

- `api/pkg/handlers/game.go`: Contains the HTTP handlers for the game-related endpoints.

## Models

- `api/pkg/models/game.go`: Defines the data structures used in the game.

## Tests

- `api/tests/game_test.go`: Contains the test cases for the game logic.

## Documentation

- `README.md`: Provides an overview and setup instructions for the project.
- `TUTORIAL.md`: Contains detailed implementation details.
- `NOTES.md`: Lists trade-offs and limitations.

## Postman Collection

- `api/postman/SnakeGameAPI.postman_collection.json`: Contains the Postman collection for testing the API endpoints.
