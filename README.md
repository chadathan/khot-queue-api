# README.md

# Go Fiber Backend

This project is a backend application built using the Go programming language and the Fiber web framework. It provides a simple API for managing user data.

## Project Structure

```
go-fiber-backend
├── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go  # HTTP request handlers
│   ├── middleware
│   │   └── middleware.go # Middleware functions
│   └── models
│       └── models.go    # Data structures
├── pkg
│   └── utils
│       └── utils.go     # Utility functions
├── configs
│   └── config.go        # Configuration settings
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd jubili-form-backend
   ```

2. Install the dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run main.go
   ```

## Usage

Once the application is running, you can access the API at `http://localhost:8888`. Use tools like Postman or curl to interact with the endpoints.

## Next Step
1. Install postgres library
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```