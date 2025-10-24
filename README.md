# Bifrost Core

A Go-based backend application with WebSocket support, PostgreSQL database, and JWT authentication.

## Features

- **RESTful API** with Gorilla Mux router
- **WebSocket support** using Socket.IO
- **PostgreSQL database** with GORM ORM
- **JWT authentication** for secure API access
- **CORS support** for cross-origin requests
- **Static file serving**
- **Environment configuration** with dotenv

## Prerequisites

- Go 1.23.0 or higher
- PostgreSQL 15+ with PostGIS extension
- Git

## Installation

### 1. Clone the repository
```bash
git clone <repository-url>
cd core
```

### 2. Install PostgreSQL and PostGIS

**Ubuntu/Debian:**
```bash
sudo apt-get install postgresql-15-postgis-3
# or for PostgreSQL 17
sudo apt install postgis postgresql-17-postgis-3
```

**macOS:**
```bash
brew install postgresql postgis
```

### 3. Set up environment variables
```bash
cp env.sample .env
# Edit .env with your database credentials and other settings
```

### 4. Install Go dependencies
```bash
go mod download
```

### 5. Run database migrations
```bash
go run main.go
```

```bash
go run main.go -migrate
```


```bash
go run main.go -seed
```

## Project Structure

```
core/
├── constants/          # Application constants and error definitions
├── models/            # Data models
├── routes/            # HTTP route handlers and middleware
├── services/          # Business logic and external services
│   ├── db/           # Database operations and repositories
│   └── socket/       # WebSocket server implementation
├── static/           # Static files served by the application
├── types/            # Custom type definitions
├── utils/            # Utility functions and helpers
├── main.go           # Application entry point
└── go.mod            # Go module dependencies
```

## API Endpoints

- `GET /` - Home endpoint
- `POST /packet` - Main packet handler for authentication and other actions
- `GET /static/*` - Static file serving

## Authentication

The application uses JWT tokens for authentication. Include the token in the Authorization header:

```
Authorization: <your-jwt-token>
```

## WebSocket

WebSocket server runs alongside the HTTP server and handles real-time communication.

## Development

To run the application in development mode:

```bash
go run main.go
```

The server will start on the port specified in your `.env` file.

## Dependencies

- **Gorilla Mux** - HTTP router and URL matcher
- **GORM** - ORM library for Go
- **PostgreSQL Driver** - Database driver for PostgreSQL
- **Socket.IO** - WebSocket library
- **JWT** - JSON Web Token implementation
- **CORS** - Cross-Origin Resource Sharing middleware


## Known Errors
- Chats.go PinnedMsg   *Message   `gorm:"foreignKey:PinnedMsgID;references:ID"`

## License

[Add your license information here]