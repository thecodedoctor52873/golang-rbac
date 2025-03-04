# Golang RBAC Project - Setup & Run Instructions

## 1. Clone the repository
```sh
git clone https://github.com/yourusername/golang-rbac.git
cd golang-rbac
```

## 2. Install dependencies
```sh
go mod tidy
```

## 3. Run the application
```sh
go run cmd/main.go
```

## 4. Build and Run using Docker
```sh
docker build -t golang-rbac .
docker run -p 8080:8080 golang-rbac
```

## 5. Test API Endpoints

### Register User
```sh
curl -X POST http://localhost:8080/register -d '{"username":"admin","password":"secret","role":"admin"}' -H "Content-Type: application/json"
```

### Login
```sh
curl -X POST http://localhost:8080/login -d '{"username":"admin","password":"secret"}' -H "Content-Type: application/json"
```

### Access Protected Endpoint
```sh
curl -H "Authorization: Bearer <your_token>" http://localhost:8080/admin/data
```

