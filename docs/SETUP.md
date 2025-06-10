## Setup & Configuration

### Environment Variables

```
# Database Configuration
DB_USER=username
DB_PASS=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=ecomdb

# API Configuration
SECRET_KEY=your_jwt_secret_key_at_least_32_chars_long

# Service Configuration
GRPC_SVC_ADDR=0.0.0.0:9091
SVC_ADDR=0.0.0.0:9091
```

### Database Setup

The project uses MySQL database with migrations managed by `golang-migrate`. Database schema is defined in migration files:

1. Initial Schema: Products, Orders, Order Items
2. Users Table
3. Sessions Table
4. User-Order Relationship
5. Order Status
6. Notification System for Order Status Changes

### Running Migrations

```bash
go run db/migrations/migrate.go
```

### Starting the Services

Start the gRPC service:
```bash
go run cmd/ecom-grpc/main.go
```

Start the API service:
```bash
go run cmd/ecom-api/main.go
```

## Development

### Generating gRPC Code

The project uses Protocol Buffers for defining the service interface. To regenerate the gRPC code after modifying the `.proto` file:

```bash
# Install protoc compiler if not already installed
# https://grpc.io/docs/protoc-installation/

# Generate Go code from the proto definition
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ecom-grpc/pb/api.proto
```

### Running Tests

```bash
go test ./...
```

## Authentication & Security

The system uses JWT for authentication with separate access and refresh tokens. Access tokens have a shorter lifespan, while refresh tokens are used to obtain new access tokens.

Security features:
- Password hashing using bcrypt
- Role-based access control (user/admin)
- Token revocation
- HTTPS support (configure in production)

## Project Structure

```
.
├── cmd/
│   ├── ecom-api/          # API service entry point
│   └── ecom-grpc/         # gRPC service entry point
├── db/
│   ├── migrations/        # Database migrations
│   └── db.go              # Database connection management
├── ecom-api/
│   └── handler/           # HTTP handlers and routing
├── ecom-grpc/
│   ├── pb/                # gRPC protocol definitions
│   ├── server/            # gRPC server implementation
│   └── storer/            # Database storage implementation
├── env/                   # Environment configuration
├── token/                 # JWT token management
└── util/                  # Utility functions
```

## License

MIT
