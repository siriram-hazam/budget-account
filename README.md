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
    If your service uses a database, run migrations:

```sh
migrate -path db/migrations -database "postgres://postgres:123456@localhost:5432/auth?sslmode=disable" up
```

4.  **Run the gRPC Service**

```sh
go run ./cmd/auth
```

Or use [air](https://github.com/cosmtrek/air) for hot reload:

```sh
air
```

## gRPC API

- **Login**
- Request: `username`, `password`
- Response: `token`, `error`

See the proto definition in [`proto/auth.proto`](proto/auth.proto).

## Example Usage

```sh
grpcurl  -plaintext  -d  '{"username":"testuser","password":"testpass"}'  localhost:8080  auth.AuthService/Login
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

**Maintainer:** [siriramhazam](https://github.com/siriramhazam)
