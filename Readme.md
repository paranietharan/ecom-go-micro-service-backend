# E-Commerce Microservice Backend

A robust Go-based backend for e-commerce applications built using a microservices architecture. This project implements a clean separation between the API layer and business logic, using gRPC for service communication.

## Architecture

The system consists of two main microservices:

1. **ecom-api**: HTTP REST API layer for client applications
2. **ecom-grpc**: Core business logic service

```
┌────────────┐      gRPC      ┌────────────┐       ┌────────────┐
│  ecom-api  │ ◄────────────► │ ecom-grpc  │ ◄───► │   MySQL    │
│  (REST)    │                │ (Service)  │       │ (Database) │
└────────────┘                └────────────┘       └────────────┘
```

## Core Features

- **Product Management**: CRUD operations for products
- **Order Management**: Create, read, delete orders with order items
- **User Management**: Authentication, authorization with JWT
- **Session Management**: Token-based session handling
- **Notification System**: Order status change notifications

## Project Structure

```
.
├── cmd/
│   ├── ecom-api/          # API service entry point
│   └── ecom-grpc/         # gRPC service entry point
├── db/
│   ├── migrations/        # Database migrations
│   │   ├── *.up.sql       # Migration scripts for applying changes
│   │   ├── *.down.sql     # Migration scripts for rolling back
│   │   └── migrate.go     # Migration tool
│   └── db.go              # Database connection management
├── ecom-api/
│   └── handler/           # HTTP handlers and routing
├── ecom-grpc/
│   ├── pb/                # gRPC protocol definitions
│   │   ├── api.proto      # Proto definition file
│   │   ├── api.pb.go      # Generated protocol message code
│   │   └── api_grpc.pb.go # Generated gRPC service code
│   ├── server/            # gRPC server implementation
│   └── storer/            # Database storage implementation
├── env/                   # Environment configuration
├── token/                 # JWT token management
└── util/                  # Utility functions
```

## Key Components

### API Layer (ecom-api)

- REST endpoints for client applications
- JWT-based authentication and authorization
- Request/response mapping between HTTP and gRPC

### Service Layer (ecom-grpc)

- Core business logic implementation
- Database interactions
- Data validation and transformation

### Database

- MySQL database with structured data model
- Database schema migrations using golang-migrate

## Data Models

### Product

```go
type Product struct {
    ID           int64
    Name         string
    Image        string
    Category     string
    Description  string
    Rating       int64
    NumReviews   int64
    Price        float32
    CountInStock int64
    CreatedAt    time.Time
    UpdatedAt    *time.Time
}
```

### Order

```go
type Order struct {
    ID            int64
    PaymentMethod string
    TaxPrice      float32
    ShippingPrice float32
    TotalPrice    float32
    UserID        int64
    Status        string  // pending, shipped, delivered
    CreatedAt     time.Time
    UpdatedAt     *time.Time
    Items         []OrderItem
}
```

### User

```go
type User struct {
    ID        int64
    Name      string
    Email     string
    Password  string  // Hashed
    IsAdmin   bool
    CreatedAt time.Time
    UpdatedAt *time.Time
}
```

### Session

```go
type Session struct {
    ID           string
    UserEmail    string
    RefreshToken string
    IsRevoked    bool
    CreatedAt    time.Time
    ExpiresAt    time.Time
}
```

## API Endpoints

### Products

- `GET /products`: List all products
- `GET /products/{id}`: Get a specific product
- `POST /products`: Create a product (admin only)
- `PATCH /products/{id}`: Update a product (admin only)
- `DELETE /products/{id}`: Delete a product (admin only)

### Orders

- `GET /myorder`: Get authenticated user's order
- `GET /orders`: List all orders (admin only)
- `POST /orders`: Create a new order
- `DELETE /orders/{id}`: Delete an order

### Users

- `POST /users`: Register a new user
- `POST /users/login`: Login as a user
- `GET /users`: List all users (admin only)
- `PATCH /users`: Update user profile
- `DELETE /users/{id}`: Delete a user (admin only)
- `POST /users/logout`: Logout (invalidate session)

### Tokens

- `POST /tokens/renew`: Renew access token using refresh token
- `POST /tokens/revoke`: Revoke a session

## Setup & Development

### Environment Variables

Create a `.env` file in the project root with the following values:

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

### Running The Project

1. **Initialize Database**

   Run the database migrations:
   
   ```bash
   go run db/migrations/migrate.go
   ```
   
   This will create all necessary tables in the MySQL database:
   - products
   - orders
   - order_items
   - users
   - sessions
   - notification_states
   - notification_events_queue
   
   The migration system uses [golang-migrate](https://github.com/golang-migrate/migrate) to manage database schema evolution through versioned migrations.

2. **Start the Services**

   Start the gRPC service:
   
   ```bash
   go run cmd/ecom-grpc/main.go
   ```
   
   Start the API service:
   
   ```bash
   go run cmd/ecom-api/main.go
   ```

### Database Migrations

The project maintains a series of SQL migration files in `db/migrations/`:

- `20240615023552_init_schema`: Creates products, orders, order_items tables
- `20240716220538_add_users_table`: Adds user management
- `20240716233848_add_sessions_table`: Adds session handling
- `20240724232703_add_user_id_fk`: Links users to orders
- `20240817135742_add_order_status`: Adds status tracking to orders
- `20240818130930_add_notification_events_tables`: Adds notification system

Each migration has an `.up.sql` file for applying changes and a `.down.sql` file for rolling back changes.

### Working with gRPC

The service interface is defined using Protocol Buffers in `ecom-grpc/pb/api.proto`. This defines:
- Message types (ProductReq, ProductRes, OrderReq, etc.)
- Service methods (CreateProduct, GetProduct, ListProducts, etc.)

To modify the service interface:

1. Edit the `api.proto` file
2. Regenerate Go code with:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ecom-grpc/pb/api.proto
```

This will update the generated code in:
- `api.pb.go`: Protocol message code
- `api_grpc.pb.go`: gRPC service code

## Testing

Run all tests:

```bash
go test ./...
```

Run specific tests:

```bash
go test ./ecom-grpc/storer -v
```

## Security

- **Password Storage**: Passwords are hashed using bcrypt
- **Authentication**: JWT tokens with separate access and refresh tokens
- **Authorization**: Role-based (admin/user) protection of endpoints
- **Sessions**: Revocable refresh tokens

## Documentation

Detailed documentation is available in the `/docs` folder:

- [DATABASE.md](docs/DATABASE.md): Database schema details
- [MIGRATIONS.md](docs/MIGRATIONS.md): Database migration information
- [SETUP.md](docs/SETUP.md): Setup and configuration instructions
- [GRPC.md](docs/GRPC.md): gRPC service implementation details
- [POSTMAN.md](docs/POSTMAN.md): Comprehensive Postman testing guide with sample requests
