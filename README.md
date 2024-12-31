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

### Get All Todos

- **URL**: `/api/todos`
- **Method**: `GET`
- **Success Response**:
    - **Code**: 200
    - **Content**: 
    ```json
    [
        {
            "_id": "60c72b2f9b1d8e6d88f8e8b7",
            "completed": false,
            "body": "Sample todo"
        }
    ]
    ```

### Create a New Todo

- **URL**: `/api/todos`
- **Method**: `POST`
- **Data Params**:
    - **Content**: 
    ```json
    {
        "body": "New todo item"
    }
    ```
- **Success Response**:
    - **Code**: 201
    - **Content**: 
    ```json
    {
        "_id": "60c72b2f9b1d8e6d88f8e8b7",
        "completed": false,
        "body": "New todo item"
    }
    ```

### Update a Todo

- **URL**: `/api/todos/:id`
- **Method**: `PATCH`
- **URL Params**: `id=[string]`
- **Success Response**:
    - **Code**: 200
    - **Content**: 
    ```json
    {
        "msg": "Todo updated",
        "success": true
    }
    ```

### Delete a Todo

- **URL**: `/api/todos/:id`
- **Method**: `DELETE`
- **URL Params**: `id=[string]`
- **Success Response**:
    - **Code**: 200
    - **Content**: 
    ```json
    {
        "msg": "Todo deleted",
        "success": true
    }
    ```

## Client Documentation

To interact with the API, you can use any HTTP client like Postman, Insomnia, or even cURL. Below are examples of how to use cURL to interact with the API:

### Get All Todos
```sh
curl -X GET http://localhost:3000/api/todos
```

### Create a New Todo
```sh
curl -X POST http://localhost:3000/api/todos -H "Content-Type: application/json" -d '{"body": "New todo item"}'
```

### Update a Todo
```sh
curl -X PATCH http://localhost:3000/api/todos/<id> -H "Content-Type: application/json"
```

### Delete a Todo
```sh
curl -X DELETE http://localhost:3000/api/todos/<id>
```

## Frontend

The frontend for this application can be built using any modern JavaScript framework like React, Vue, or Angular. Below is an example of how to set up a simple React frontend to interact with the API.

### Setting Up React Frontend

1. Create a new React application:
    ```sh
    npx create-react-app todo-app
    cd todo-app
    ```

2. Install Axios for making HTTP requests:
    ```sh
    npm install axios
    ```

3. Create a new file `src/api.js` to handle API requests:
    ```javascript
    import axios from 'axios';

    const API_URL = 'http://localhost:3000/api/todos';

    export const getTodos = () => axios.get(API_URL);
    export const createTodo = (todo) => axios.post(API_URL, todo);
    export const updateTodo = (id) => axios.patch(`${API_URL}/${id}`);
    export const deleteTodo = (id) => axios.delete(`${API_URL}/${id}`);
    ```

4. Use the API functions in your React components to interact with the backend.

5. Run the application:
    ```sh
    npm start
    ```

6. The frontend will be running on `http://localhost:3000`. 

## Contributing

If you find any bugs or have suggestions for improvement, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/).

## Acknowledgments

Special thanks to [Fiber](https://github.com/gofiber/fiber) for their excellent web framework.