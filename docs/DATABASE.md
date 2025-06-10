# Database Schema

## Tables Overview

The database schema consists of the following tables:

- `products`: Stores product information
- `orders`: Stores order information
- `order_items`: Links products to orders
- `users`: Stores user information
- `sessions`: Manages user authentication sessions
- `notification_states`: Tracks notification delivery status
- `notification_events_queue`: Queues notifications for delivery

## Schema Evolution

The schema is managed through a series of migrations:

### 1. Initial Schema (20240615023552_init_schema)

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

```sql
ALTER TABLE `orders`
  ADD COLUMN `user_id` int NOT NULL,
  ADD CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`);
```

### 5. Add Order Status (20240817135742_add_order_status)

```sql
ALTER TABLE `orders`
  ADD COLUMN `status` ENUM('pending', 'shipped', 'delivered') NOT NULL DEFAULT 'pending';
```

### 6. Add Notification Tables (20240818130930_add_notification_events_tables)

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

## Entity-Relationship Diagram

```
┌─────────────┐       ┌─────────────┐       ┌─────────────┐
│   products  │       │    orders   │       │    users    │
├─────────────┤       ├─────────────┤       ├─────────────┤
│ id          │       │ id          │       │ id          │
│ name        │       │ user_id     │◄──────┤ name        │
│ image       │       │ payment_met │       │ email       │
│ category    │       │ tax_price   │       │ password    │
│ description │       │ shipping_p  │       │ is_admin    │
│ rating      │       │ total_price │       │ created_at  │
│ num_reviews │       │ status      │       │ updated_at  │
│ price       │       │ created_at  │       └─────────────┘
│ count_stock │       │ updated_at  │
│ created_at  │       └─────────────┘
│ updated_at  │             │
└─────────────┘             │
      │                     │
      │                     │
      ▼                     ▼
┌─────────────┐       ┌─────────────┐       ┌─────────────┐
│ order_items │       │notification_│       │  sessions   │
├─────────────┤       │   states    │       ├─────────────┤
│ id          │       ├─────────────┤       │ id          │
│ order_id    │       │ id          │       │ user_email  │
│ product_id  │◄──────┤ order_id    │       │ refresh_tok │
│ name        │       │ state       │       │ is_revoked  │
│ quantity    │       │ message     │       │ created_at  │
│ image       │       │ requested_at│       │ expires_at  │
│ price       │       │ completed_at│       └─────────────┘
└─────────────┘       └─────────────┘
                             │
                             │
                             ▼
                      ┌─────────────┐
                      │notification_│
                      │events_queue │
                      ├─────────────┤
                      │ id          │
                      │ user_email  │
                      │ order_status│
                      │ order_id    │
                      │ state_id    │
                      │ attempts    │
                      │ created_at  │
                      │ updated_at  │
                      └─────────────┘
```
