# Budget Authen Service

Budget Authen Service is a microservice for generating JWT tokens, designed to provide authentication via gRPC.

## Features

- Generate JWT tokens for user authentication
- gRPC API for login and token generation
- Example HTTP handler for testing

## Development Setup

1.  **Clone the Repository**

    ```sh
    git clone https://github.com/siriram-hazam/budget-authen.git
    cd budget-authen
    ```

2.  **Download Go Modules**

    ```sh
    go mod download
    ```

3.  **(Optional) Run Database Migrations**

    ```sh
    make migrate
    ```

4.  **Run the gRPC Service**

    ```sh
    make run
    ```

5.  **(Optional) Use [air](https://github.com/cosmtrek/air) for hot reload**

    ```sh
    make air
    ```

6.  **Docker Usage**

    - **Build image:**
      ```sh
      make docker-build
      ```
    - **Run container:**
      ```sh
      make docker-run
      ```
    - **For hot reload in Docker (dev):**
      ```sh
      make docker-dev
      ```
      > This will mount your local code into the container and use air for hot reload.  
      > Edit your code locally and air will reload the service automatically.

## Makefile Commands

- `make build` – Build the binary
- `make run` – Run the service
- `make test` – Run tests
- `make migrate` – Run database migrations
- `make proto` – Generate gRPC code from proto
- `make air` – Run with hot reload (requires air)
- `make docker-build` – Build Docker image
- `make docker-run` – Run Docker container
- `make docker-dev` – Run Docker container with hot reload (air)

## gRPC API

- **Login**
  - Request: `username`, `password`
  - Response: `token`, `error`

See the proto definition in [`proto/auth.proto`](proto/auth.proto).

## Example Usage

```sh
grpcurl -plaintext -d '{"username":"testuser","password":"testpass"}' localhost:8080 auth.AuthService/Login
```

## Project Structure

- `cmd/auth` - gRPC server entry point
- `internal/handler` - HTTP handler (example)
- `internal/service` - Core authentication and JWT logic
- `internal/utils` - Utility functions (response, error handling)
- `grpc-auth/proto` - Generated code from proto files
- `db/migrations` - Database migration files

## Notes

- This service currently does not connect to a real user database (demo for JWT only).
- You can extend or customize features as needed.

---
