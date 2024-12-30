# To-do App

This is a simple To-do application built with Go and Fiber. It allows users to create, read, update, and delete to-do items. The application uses MongoDB as the database to store the to-do items.

## Features

- Create a new to-do item
- Get all to-do items
- Update a to-do item
- Delete a to-do item

## Tech Stack

- **Go**: The programming language used for the backend.
- **Fiber**: An Express-inspired web framework for Go.
- **MongoDB**: A NoSQL database used to store the to-do items.
- **Go MongoDB Driver**: The official MongoDB driver for Go.
- **GoDotEnv**: A package to load environment variables from a `.env` file.

## Getting Started

To run this project locally, follow these steps:

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd To_do_App
    ```

2. Create a `.env` file in the root directory and add your MongoDB URI and port:
    ```env
    MONGODB_URI=<your-mongodb-uri>
    PORT=3000
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

4. The server will be running on `http://localhost:3000`.

## API Endpoints

- `GET /api/todos`: Get all to-do items.
- `POST /api/todos`: Create a new to-do item.
- `PATCH /api/todos/:id`: Update a to-do item.
- `DELETE /api/todos/:id`: Delete a to-do item.
