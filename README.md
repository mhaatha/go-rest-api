# Go REST API

A simple RESTful API built with Go, created as a learning project. This application demonstrates a clean, layered architecture for building web services in Go. It provides basic CRUD (Create, Read, Update, Delete) operations for a `Category` resource.

## Project Structure

The codebase follows a layered architecture to separate concerns, making it modular and easier to maintain.

-   **/app**: Handles application concerns like database connections (`database.go`) and routing (`router.go`).
-   **/controller**: Contains the HTTP handlers that receive requests and pass them to the service layer.
-   **/service**: Implements the core business logic of the application.
-   **/repository**: The data access layer, responsible for all communication with the database.
-   **/model**: Defines the data structures, including domain models and web request/response objects.
-   **/exception**: Manages custom error handling.
-   **/middleware**: Contains HTTP middleware, such as for authentication.
-   **/helper**: Utility functions used across the application.

## API Endpoints

The API provides the following endpoints for managing categories.

| Method | Endpoint                      | Description                  |
| :----- | :---------------------------- | :--------------------------- |
| `GET`    | `/api/categories`             | Get all categories           |
| `GET`    | `/api/categories/{categoryId}`| Get a single category by ID  |
| `POST`   | `/api/categories`             | Create a new category        |
| `PUT`    | `/api/categories/{categoryId}`| Update an existing category  |
| `DELETE` | `/api/categories/{categoryId}`| Delete a category by ID      |

## Prerequisites

Before you begin, ensure you have the following installed:

-   [Go](https://go.dev/doc/install) (version 1.24 or later)
-   [MySQL](https://dev.mysql.com/downloads/installer/)

## Getting Started

Follow these steps to get the application up and running.

**1. Clone the repository:**

```sh
git clone https://github.com/your-username/go-rest-api.git
cd go-rest-api
```

**2. Set up the database:**

-   Connect to your MySQL instance.
-   Create the database for the project:
    ```sql
    CREATE DATABASE go_rest_api;
    ```
-   Create the `category` table:
    ```sql
    USE go_rest_api;
    CREATE TABLE category (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
    );
    ```

**3. Configure environment variables:**

-   Create a `.env` file in the root of the project.
-   Add your database password to the `.env` file:
    ```
    DB_PASSWORD="your_mysql_password"
    ```

**4. Install dependencies:**

The project uses Go Modules to manage dependencies. They will be downloaded automatically when you build or run the application.

```sh
go mod tidy
```

**5. Run the server:**

```sh
go run main.go
```

The server will start on `http://localhost:3000`.

## Running Tests

To run the unit tests for this project, execute the following command from the root directory:

```sh
go test ./...
```
