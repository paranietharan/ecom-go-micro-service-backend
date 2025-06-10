# Postman Testing Guide

This document provides detailed information about testing the E-Commerce Microservice Backend using Postman.

## Base URL

```
http://localhost:8080
```

## Authentication Flow

1. Register a new user
2. Login to obtain access and refresh tokens
3. Use the access token in the Authorization header for subsequent requests

## Endpoints Reference

### Authentication & User Management

#### Register User

- **Endpoint**: `POST /users`
- **Description**: Creates a new user account.
- **Auth Required**: No
- **Sample Payload**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword123",
    "is_admin": false
  }
  ```
- **Expected Response** (201 Created):
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "is_admin": false
  }
  ```

#### Register Admin User

- **Endpoint**: `POST /users`
- **Description**: Creates a new admin account.
- **Auth Required**: No
- **Sample Payload**:
  ```json
  {
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "adminpassword123",
    "is_admin": true
  }
  ```
- **Expected Response** (201 Created):
  ```json
  {
    "name": "Admin User",
    "email": "admin@example.com",
    "is_admin": true
  }
  ```

#### Login

- **Endpoint**: `POST /users/login`
- **Description**: Authenticates a user and provides access and refresh tokens.
- **Auth Required**: No
- **Sample Payload**:
  ```json
  {
    "email": "john@example.com",
    "password": "securepassword123"
  }
  ```
- **Expected Response** (200 OK):
  ```json
  {
    "session_id": "123e4567-e89b-12d3-a456-426614174000",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "access_token_expires_at": "2025-06-10T14:15:22Z",
    "refresh_token_expires_at": "2025-06-11T14:00:22Z",
    "user": {
      "name": "John Doe",
      "email": "john@example.com",
      "is_admin": false
    }
  }
  ```

#### List All Users (Admin Only)

- **Endpoint**: `GET /users`
- **Description**: Retrieves all registered users.
- **Auth Required**: Yes (Admin only)
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (200 OK):
  ```json
  {
    "users": [
      {
        "name": "John Doe",
        "email": "john@example.com",
        "is_admin": false
      },
      {
        "name": "Admin User",
        "email": "admin@example.com",
        "is_admin": true
      }
    ]
  }
  ```

#### Update User Profile

- **Endpoint**: `PATCH /users`
- **Description**: Updates the authenticated user's profile.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Sample Payload**:
  ```json
  {
    "name": "John Updated",
    "password": "newpassword123"
  }
  ```
- **Expected Response** (200 OK):
  ```json
  {
    "name": "John Updated",
    "email": "john@example.com",
    "is_admin": false
  }
  ```

#### Delete User (Admin Only)

- **Endpoint**: `DELETE /users/{id}`
- **Description**: Deletes a user by ID.
- **Auth Required**: Yes (Admin only)
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (204 No Content)

#### Logout

- **Endpoint**: `POST /users/logout`
- **Description**: Logs out the current user and invalidates their session.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (204 No Content)

### Token Management

#### Renew Access Token

- **Endpoint**: `POST /tokens/renew`
- **Description**: Generates a new access token using a valid refresh token.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Sample Payload**:
  ```json
  {
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```
- **Expected Response** (200 OK):
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "access_token_expires_at": "2025-06-10T14:30:22Z"
  }
  ```

#### Revoke Session

- **Endpoint**: `POST /tokens/revoke`
- **Description**: Revokes the current session, invalidating all associated tokens.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (204 No Content)

### Product Management

#### Create Product (Admin Only)

- **Endpoint**: `POST /products`
- **Description**: Creates a new product in the catalog.
- **Auth Required**: Yes (Admin only)
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Sample Payload**:
  ```json
  {
    "name": "Wireless Headphones",
    "image": "headphones.jpg",
    "category": "Electronics",
    "description": "Premium wireless headphones with noise cancellation",
    "rating": 4,
    "num_reviews": 0,
    "price": 99.99,
    "count_in_stock": 50
  }
  ```
- **Expected Response** (201 Created):
  ```json
  {
    "id": 1,
    "name": "Wireless Headphones",
    "image": "headphones.jpg",
    "category": "Electronics",
    "description": "Premium wireless headphones with noise cancellation",
    "rating": 4,
    "num_reviews": 0,
    "price": 99.99,
    "count_in_stock": 50,
    "created_at": "2025-06-10T14:00:00Z",
    "updated_at": null
  }
  ```

#### List All Products

- **Endpoint**: `GET /products`
- **Description**: Retrieves all products in the catalog.
- **Auth Required**: No
- **Expected Response** (200 OK):
  ```json
  [
    {
      "id": 1,
      "name": "Wireless Headphones",
      "image": "headphones.jpg",
      "category": "Electronics",
      "description": "Premium wireless headphones with noise cancellation",
      "rating": 4,
      "num_reviews": 0,
      "price": 99.99,
      "count_in_stock": 50,
      "created_at": "2025-06-10T14:00:00Z",
      "updated_at": null
    },
    {
      "id": 2,
      "name": "Smart Watch",
      "image": "smartwatch.jpg",
      "category": "Electronics",
      "description": "Feature-rich smartwatch with health monitoring",
      "rating": 5,
      "num_reviews": 0,
      "price": 199.99,
      "count_in_stock": 25,
      "created_at": "2025-06-10T14:15:00Z",
      "updated_at": null
    }
  ]
  ```

#### Get Product by ID

- **Endpoint**: `GET /products/{id}`
- **Description**: Retrieves a specific product by ID.
- **Auth Required**: No
- **Expected Response** (200 OK):
  ```json
  {
    "id": 1,
    "name": "Wireless Headphones",
    "image": "headphones.jpg",
    "category": "Electronics",
    "description": "Premium wireless headphones with noise cancellation",
    "rating": 4,
    "num_reviews": 0,
    "price": 99.99,
    "count_in_stock": 50,
    "created_at": "2025-06-10T14:00:00Z",
    "updated_at": null
  }
  ```

#### Update Product (Admin Only)

- **Endpoint**: `PATCH /products/{id}`
- **Description**: Updates a product by ID.
- **Auth Required**: Yes (Admin only)
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Sample Payload**:
  ```json
  {
    "price": 89.99,
    "count_in_stock": 45,
    "description": "Updated premium wireless headphones with enhanced noise cancellation"
  }
  ```
- **Expected Response** (200 OK):
  ```json
  {
    "id": 1,
    "name": "Wireless Headphones",
    "image": "headphones.jpg",
    "category": "Electronics",
    "description": "Updated premium wireless headphones with enhanced noise cancellation",
    "rating": 4,
    "num_reviews": 0,
    "price": 89.99,
    "count_in_stock": 45,
    "created_at": "2025-06-10T14:00:00Z",
    "updated_at": "2025-06-10T15:30:00Z"
  }
  ```

#### Delete Product (Admin Only)

- **Endpoint**: `DELETE /products/{id}`
- **Description**: Deletes a product by ID.
- **Auth Required**: Yes (Admin only)
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (204 No Content)

### Order Management

#### Create Order

- **Endpoint**: `POST /orders`
- **Description**: Creates a new order with the specified items.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Sample Payload**:
  ```json
  {
    "items": [
      {
        "name": "Wireless Headphones",
        "quantity": 2,
        "image": "headphones.jpg",
        "price": 89.99,
        "product_id": 1
      },
      {
        "name": "Smart Watch",
        "quantity": 1,
        "image": "smartwatch.jpg",
        "price": 199.99,
        "product_id": 2
      }
    ],
    "payment_method": "PayPal",
    "tax_price": 38.00,
    "shipping_price": 10.00,
    "total_price": 427.97
  }
  ```
- **Expected Response** (201 Created):
  ```json
  {
    "id": 1,
    "items": [
      {
        "name": "Wireless Headphones",
        "quantity": 2,
        "image": "headphones.jpg",
        "price": 89.99,
        "product_id": 1
      },
      {
        "name": "Smart Watch",
        "quantity": 1,
        "image": "smartwatch.jpg",
        "price": 199.99,
        "product_id": 2
      }
    ],
    "payment_method": "PayPal",
    "tax_price": 38.00,
    "shipping_price": 10.00,
    "total_price": 427.97,
    "created_at": "2025-06-10T16:00:00Z",
    "updated_at": null
  }
  ```

#### Get User's Order

- **Endpoint**: `GET /myorder`
- **Description**: Retrieves the authenticated user's order.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (200 OK):
  ```json
  {
    "id": 1,
    "items": [
      {
        "name": "Wireless Headphones",
        "quantity": 2,
        "image": "headphones.jpg",
        "price": 89.99,
        "product_id": 1
      },
      {
        "name": "Smart Watch",
        "quantity": 1,
        "image": "smartwatch.jpg",
        "price": 199.99,
        "product_id": 2
      }
    ],
    "payment_method": "PayPal",
    "tax_price": 38.00,
    "shipping_price": 10.00,
    "total_price": 427.97,
    "created_at": "2025-06-10T16:00:00Z",
    "updated_at": null
  }
  ```

#### List All Orders (Admin Only)

- **Endpoint**: `GET /orders`
- **Description**: Retrieves all orders in the system.
- **Auth Required**: Yes (Admin only)
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (200 OK):
  ```json
  [
    {
      "id": 1,
      "items": [
        {
          "name": "Wireless Headphones",
          "quantity": 2,
          "image": "headphones.jpg",
          "price": 89.99,
          "product_id": 1
        },
        {
          "name": "Smart Watch",
          "quantity": 1,
          "image": "smartwatch.jpg",
          "price": 199.99,
          "product_id": 2
        }
      ],
      "payment_method": "PayPal",
      "tax_price": 38.00,
      "shipping_price": 10.00,
      "total_price": 427.97,
      "created_at": "2025-06-10T16:00:00Z",
      "updated_at": null
    },
    {
      "id": 2,
      "items": [
        {
          "name": "Smart Watch",
          "quantity": 1,
          "image": "smartwatch.jpg",
          "price": 199.99,
          "product_id": 2
        }
      ],
      "payment_method": "Credit Card",
      "tax_price": 20.00,
      "shipping_price": 5.00,
      "total_price": 224.99,
      "created_at": "2025-06-10T16:30:00Z",
      "updated_at": null
    }
  ]
  ```

#### Delete Order

- **Endpoint**: `DELETE /orders/{id}`
- **Description**: Deletes a specific order by ID.
- **Auth Required**: Yes
- **Headers**:
  ```
  Authorization: Bearer {access_token}
  ```
- **Expected Response** (204 No Content)

## Testing Sequence

For comprehensive testing, follow this sequence:

1. **User Registration and Authentication**:
   - Register regular and admin users
   - Login with each user to obtain tokens

2. **Product Management**:
   - Create products (admin)
   - List products (all users)
   - Get individual products (all users)
   - Update products (admin)
   - Delete products (admin)

3. **Order Management**:
   - Create orders
   - Get user's orders
   - List all orders (admin)
   - Delete orders

4. **Token Management**:
   - Renew access token
   - Revoke session
   - Logout

## Postman Collection

You can import the following Postman collection to quickly get started with testing:

### Collection Schema (v2.1.0)

```json
{
  "info": {
    "_postman_id": "f8e2c3d1-b5a6-4a87-8c3e-9d1b2a3c4d5e",
    "name": "E-Commerce Microservice Backend",
    "description": "API collection for testing the E-Commerce Microservice Backend",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Authentication",
      "item": [
        {
          "name": "Register User",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"John Doe\",\n  \"email\": \"john@example.com\",\n  \"password\": \"securepassword123\",\n  \"is_admin\": false\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/users",
              "host": ["{{baseUrl}}"],
              "path": ["users"]
            },
            "description": "Register a new user"
          }
        },
        {
          "name": "Register Admin User",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"Admin User\",\n  \"email\": \"admin@example.com\",\n  \"password\": \"adminpassword123\",\n  \"is_admin\": true\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/users",
              "host": ["{{baseUrl}}"],
              "path": ["users"]
            },
            "description": "Register an admin user"
          }
        },
        {
          "name": "Login User",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "var jsonData = pm.response.json();",
                  "pm.environment.set(\"userAccessToken\", jsonData.access_token);",
                  "pm.environment.set(\"userRefreshToken\", jsonData.refresh_token);"
                ],
                "type": "text/javascript"
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"email\": \"john@example.com\",\n  \"password\": \"securepassword123\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/users/login",
              "host": ["{{baseUrl}}"],
              "path": ["users", "login"]
            },
            "description": "Login as regular user"
          }
        },
        {
          "name": "Login Admin",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "var jsonData = pm.response.json();",
                  "pm.environment.set(\"adminAccessToken\", jsonData.access_token);",
                  "pm.environment.set(\"adminRefreshToken\", jsonData.refresh_token);"
                ],
                "type": "text/javascript"
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"email\": \"admin@example.com\",\n  \"password\": \"adminpassword123\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/users/login",
              "host": ["{{baseUrl}}"],
              "path": ["users", "login"]
            },
            "description": "Login as admin user"
          }
        },
        {
          "name": "Renew Access Token",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "var jsonData = pm.response.json();",
                  "pm.environment.set(\"userAccessToken\", jsonData.access_token);"
                ],
                "type": "text/javascript"
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              },
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"refresh_token\": \"{{userRefreshToken}}\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/tokens/renew",
              "host": ["{{baseUrl}}"],
              "path": ["tokens", "renew"]
            },
            "description": "Renew access token using refresh token"
          }
        },
        {
          "name": "Revoke Session",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/tokens/revoke",
              "host": ["{{baseUrl}}"],
              "path": ["tokens", "revoke"]
            },
            "description": "Revoke the current session"
          }
        },
        {
          "name": "Logout",
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/users/logout",
              "host": ["{{baseUrl}}"],
              "path": ["users", "logout"]
            },
            "description": "Logout current user"
          }
        }
      ]
    },
    {
      "name": "Users",
      "item": [
        {
          "name": "List All Users (Admin)",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{adminAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/users",
              "host": ["{{baseUrl}}"],
              "path": ["users"]
            },
            "description": "List all users (admin only)"
          }
        },
        {
          "name": "Update User Profile",
          "request": {
            "method": "PATCH",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              },
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"John Updated\",\n  \"password\": \"newpassword123\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/users",
              "host": ["{{baseUrl}}"],
              "path": ["users"]
            },
            "description": "Update user profile"
          }
        },
        {
          "name": "Delete User (Admin)",
          "request": {
            "method": "DELETE",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{adminAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/users/3",
              "host": ["{{baseUrl}}"],
              "path": ["users", "3"]
            },
            "description": "Delete a user (admin only)"
          }
        }
      ]
    },
    {
      "name": "Products",
      "item": [
        {
          "name": "Create Product (Admin)",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "var jsonData = pm.response.json();",
                  "pm.environment.set(\"productId\", jsonData.id);"
                ],
                "type": "text/javascript"
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              },
              {
                "key": "Authorization",
                "value": "Bearer {{adminAccessToken}}"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"name\": \"Wireless Headphones\",\n  \"image\": \"headphones.jpg\",\n  \"category\": \"Electronics\",\n  \"description\": \"Premium wireless headphones with noise cancellation\",\n  \"rating\": 4,\n  \"num_reviews\": 0,\n  \"price\": 99.99,\n  \"count_in_stock\": 50\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/products",
              "host": ["{{baseUrl}}"],
              "path": ["products"]
            },
            "description": "Create a new product (admin only)"
          }
        },
        {
          "name": "List All Products",
          "request": {
            "method": "GET",
            "url": {
              "raw": "{{baseUrl}}/products",
              "host": ["{{baseUrl}}"],
              "path": ["products"]
            },
            "description": "List all products"
          }
        },
        {
          "name": "Get Product by ID",
          "request": {
            "method": "GET",
            "url": {
              "raw": "{{baseUrl}}/products/{{productId}}",
              "host": ["{{baseUrl}}"],
              "path": ["products", "{{productId}}"]
            },
            "description": "Get a specific product by ID"
          }
        },
        {
          "name": "Update Product (Admin)",
          "request": {
            "method": "PATCH",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              },
              {
                "key": "Authorization",
                "value": "Bearer {{adminAccessToken}}"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"price\": 89.99,\n  \"count_in_stock\": 45,\n  \"description\": \"Updated premium wireless headphones with enhanced noise cancellation\"\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/products/{{productId}}",
              "host": ["{{baseUrl}}"],
              "path": ["products", "{{productId}}"]
            },
            "description": "Update a product (admin only)"
          }
        },
        {
          "name": "Delete Product (Admin)",
          "request": {
            "method": "DELETE",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{adminAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/products/{{productId}}",
              "host": ["{{baseUrl}}"],
              "path": ["products", "{{productId}}"]
            },
            "description": "Delete a product (admin only)"
          }
        }
      ]
    },
    {
      "name": "Orders",
      "item": [
        {
          "name": "Create Order",
          "event": [
            {
              "listen": "test",
              "script": {
                "exec": [
                  "var jsonData = pm.response.json();",
                  "pm.environment.set(\"orderId\", jsonData.id);"
                ],
                "type": "text/javascript"
              }
            }
          ],
          "request": {
            "method": "POST",
            "header": [
              {
                "key": "Content-Type",
                "value": "application/json"
              },
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "body": {
              "mode": "raw",
              "raw": "{\n  \"items\": [\n    {\n      \"name\": \"Wireless Headphones\",\n      \"quantity\": 2,\n      \"image\": \"headphones.jpg\",\n      \"price\": 89.99,\n      \"product_id\": {{productId}}\n    }\n  ],\n  \"payment_method\": \"PayPal\",\n  \"tax_price\": 18.00,\n  \"shipping_price\": 10.00,\n  \"total_price\": 207.98\n}"
            },
            "url": {
              "raw": "{{baseUrl}}/orders",
              "host": ["{{baseUrl}}"],
              "path": ["orders"]
            },
            "description": "Create a new order"
          }
        },
        {
          "name": "Get User's Order",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/myorder",
              "host": ["{{baseUrl}}"],
              "path": ["myorder"]
            },
            "description": "Get the current user's order"
          }
        },
        {
          "name": "List All Orders (Admin)",
          "request": {
            "method": "GET",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{adminAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/orders",
              "host": ["{{baseUrl}}"],
              "path": ["orders"]
            },
            "description": "List all orders (admin only)"
          }
        },
        {
          "name": "Delete Order",
          "request": {
            "method": "DELETE",
            "header": [
              {
                "key": "Authorization",
                "value": "Bearer {{userAccessToken}}"
              }
            ],
            "url": {
              "raw": "{{baseUrl}}/orders/{{orderId}}",
              "host": ["{{baseUrl}}"],
              "path": ["orders", "{{orderId}}"]
            },
            "description": "Delete an order"
          }
        }
      ]
    }
  ],
  "variable": [
    {
      "key": "baseUrl",
      "value": "http://localhost:8080"
    }
  ]
}
```

## Environment Variables

Set up the following environment variables in Postman:

- `baseUrl`: `http://localhost:8080`
- `userAccessToken`: Automatically set after successful user login
- `userRefreshToken`: Automatically set after successful user login
- `adminAccessToken`: Automatically set after successful admin login
- `adminRefreshToken`: Automatically set after successful admin login
- `productId`: Automatically set after creating a product
- `orderId`: Automatically set after creating an order

## Troubleshooting

### Common Issues

1. **Authentication Failed**: Ensure the access token is valid and has not expired.
2. **Permission Denied**: Check if your user has the required permissions (admin status) for the operation.
3. **Item Not Found**: Verify that the ID used in the request actually exists.
4. **Invalid Request Body**: Double-check the payload structure against the provided samples.

### HTTP Status Codes

- `200 OK`: Request succeeded
- `201 Created`: Resource was successfully created
- `204 No Content`: Request succeeded but no content to return
- `400 Bad Request`: Invalid request format or parameters
- `401 Unauthorized`: Authentication failed or token invalid
- `403 Forbidden`: User does not have permission for this action
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server-side error
