# Golang RBAC API

## Setup

1. Install dependencies:
   ```sh
   go mod tidy
   ```

2. Run the application:
   ```sh
   go run cmd/main.go
   ```

3. Build and run using Docker:
   ```sh
   docker build -t golang-rbac .
   docker run -p 8080:8080 golang-rbac
   ```
