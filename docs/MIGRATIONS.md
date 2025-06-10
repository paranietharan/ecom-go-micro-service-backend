# Database Migration Guide

This document provides detailed information about the database migrations in the E-Commerce Microservice Backend.

## Migration Architecture

The project uses [golang-migrate](https://github.com/golang-migrate/migrate) to manage database schema migrations. Migrations are stored as versioned SQL files in the `db/migrations/` directory.

## Migration Files

Each migration consists of two files:
- `.up.sql`: Applied when migrating forward
- `.down.sql`: Applied when rolling back a migration

## Running Migrations

You can run the migrations using the migration tool:

```bash
go run db/migrations/migrate.go
```

### Implementation Details

The migration tool connects to the MySQL database using environment variables and applies all pending migrations in sequence:

```go
func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println("Warning: .env file not found, using environment variables")
    }

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
    migrationsPath := "file://db/migrations"

    m, err := migrate.New(migrationsPath, dbURL)
    if err != nil {
        log.Fatalf("Failed to create migrate instance: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Migration failed: %v", err)
    }

    log.Println("Migration completed successfully!")
}
```

## Migration Versions

### 1. Initial Schema (20240615023552_init_schema)

Creates the initial product, order, and order item tables.

**Up Migration:**
```sql
CREATE TABLE `products` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `description` text,
  `rating` int NOT NULL,
  `num_reviews` int NOT NULL DEFAULT 0,
  `price` decimal(10,2) NOT NULL,
  `count_in_stock` int NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `payment_method` varchar(255) NOT NULL,
  `tax_price` decimal(10,2) NOT NULL,
  `shipping_price` decimal(10,2) NOT NULL,
  `total_price` decimal(10,2) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime
);

CREATE TABLE `order_items` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `quantity` int NOT NULL,
  `image` varchar(255) NOT NULL,
  `price` int NOT NULL
);

ALTER TABLE `order_items` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);
ALTER TABLE `order_items` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
```

### 2. Add Users Table (20240716220538_add_users_table)

Adds user management functionality.

**Up Migration:**
```sql
CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `is_admin` bool NOT NULL DEFAULT false,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  UNIQUE (email)
);
```

### 3. Add Sessions Table (20240716233848_add_sessions_table)

Adds session management for authentication.

**Up Migration:**
```sql
CREATE TABLE `sessions` (
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `user_email` varchar(255) NOT NULL,
  `refresh_token` varchar(512) NOT NULL,
  `is_revoked` bool NOT NULL DEFAULT false,
  `created_at` datetime DEFAULT (now()),
  `expires_at` datetime
);
```

### 4. Add User-Order Relationship (20240724232703_add_user_id_fk)

Links orders to users.

**Up Migration:**
```sql
ALTER TABLE `orders`
  ADD COLUMN `user_id` int NOT NULL,
  ADD CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`);
```

### 5. Add Order Status (20240817135742_add_order_status)

Adds order status tracking.

**Up Migration:**
```sql
ALTER TABLE `orders`
  ADD COLUMN `status` ENUM('pending', 'shipped', 'delivered') NOT NULL DEFAULT 'pending';
```

### 6. Add Notification Tables (20240818130930_add_notification_events_tables)

Adds notification system for order status changes.

**Up Migration:**
```sql
CREATE TABLE `notification_states` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `state` enum('not sent', 'sent', 'failed') NOT NULL,
  `message` varchar(512),  
  `requested_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `completed_at` datetime
);

CREATE TABLE `notification_events_queue` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_email` varchar(256) NOT NULL,
  `order_status` varchar(256) NOT NULL,
  `order_id` int NOT NULL, 
  `state_id` int NOT NULL,
  `attempts` int,  
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime
);

ALTER TABLE `notification_states`
    ADD CONSTRAINT `notification_states_order_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `notification_events_queue`
    ADD CONSTRAINT `notification_events_queue_order_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
    ADD CONSTRAINT `notification_events_queue_state_id_fk` FOREIGN KEY (`state_id`) REFERENCES `notification_states` (`id`);
```

## Rolling Back Migrations

To roll back the most recent migration:

```bash
go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate \
  -path ./db/migrations \
  -database "mysql://<USERNAME>:<PASSWORD>@tcp(<HOST>:<PORT>)/<DATABASE>" \
  down 1
```

To roll back all migrations:

```bash
go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate \
  -path ./db/migrations \
  -database "mysql://<USERNAME>:<PASSWORD>@tcp(<HOST>:<PORT>)/<DATABASE>" \
  down
```
