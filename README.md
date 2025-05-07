# E-Commerce Platform (Microservices)

## Overview

This project is a basic e-commerce platform built with Go, structured using Clean Architecture and composed of three microservices:

- **API Gateway** – Routes incoming requests to the appropriate service, handles logging and basic authentication.

- **Inventory Service** – Manages product data, categories, stock, and prices. Listens for order events through RabbitMQ to update product inventory.

- **Order Service** – Handles order creation, status updates, and product quantities per order. Publishes order events to RabbitMQ when orders are created.

All services are written in Go using gRPC for service-to-service communication and connected to PostgreSQL for persistence. The services communicate asynchronously using RabbitMQ for event-driven architecture.

## Architecture

### Message Queue Pattern

The project implements an event-driven architecture using RabbitMQ:

1. **Order Service** publishes `order.created` events when a new order is created
2. **Inventory Service** subscribes to `order.created` events and updates product stock accordingly

This approach provides:
- Decoupling between services
- Asynchronous processing
- Increased resilience (services can continue to function independently)

```
┌──────────────┐         ┌───────────────┐         ┌───────────────┐
│              │ 1.Create │               │ 2.Publish │              │
│    Client    │────────►│  Order Service │─────────►│    RabbitMQ   │
│              │         │               │         │              │
└──────────────┘         └───────────────┘         └───────┬───────┘
                                                          │
                                                          │ 3.Consume
                                                          ▼
                                               ┌───────────────────┐
                                               │                   │
                                               │ Inventory Service │
                                               │                   │
                                               └───────────────────┘
```

## Project Structure

Each microservice follows a clean architecture pattern:

```
service/
├── internal/
│   ├── delivery/
│   │   └── grpc/         # gRPC handlers
│   ├── domain/           # Domain models and interfaces
│   ├── message/          # Message queue implementation
│   ├── repository/       # Database access layer
│   └── usecase/          # Business logic
└── main.go               # Service entry point
```

## Requirements

- Go 1.18 or higher
- PostgreSQL
- RabbitMQ 3.8+ 
- Basic understanding of RESTful APIs and gRPC

## Installation

### 1. Install Go Dependencies:

   In each service directory:
   ```
    cd inventoryService
    go mod tidy

    cd ../orderService
    go mod tidy

    cd ../apiGateway
    go mod tidy
   ```

### 2. Install and Run RabbitMQ:

   Using Docker:
   ```
   docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
   ```

   Or install directly from the [RabbitMQ website](https://www.rabbitmq.com/download.html).

### 3. Set Up PostgreSQL Tables:

Inventory Service:
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    stock INT NOT NULL,
    category_id INT NOT NULL
);
```

Order Service:
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL,
    quantity INT NOT NULL
);
```

### 4. Running the services:

Inventory Service:
```
cd inventoryService
go run main.go
```
Runs gRPC server on port 50051

Order Service:
```
cd orderService
go run main.go
```
Runs gRPC server on port 50052

API-Gateway:
```
cd apiGateway
go run cmd/main.go
```
Runs on http://localhost:8080

## Authentication

All API Gateway endpoints are protected by Basic Auth:

- Username: admin
- Password: 1234
- Base64 encoded: YWRtaW46MTIzNA==

Include this in your Authorization header:
```
Authorization: Basic YWRtaW46MTIzNA==
```

## Sample API Requests (via API Gateway)

### Create Product:
```
curl -X POST http://localhost:8080/products \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{"name": "iPhone", "description": "Apple smartphone", "price": 999.99, "stock": 10, "category_id": 1}'
```

### List Products:
```
curl -X GET http://localhost:8080/products \
  -H "Authorization: Basic YWRtaW46MTIzNA=="
```

### Create Order:
```
curl -X POST http://localhost:8080/orders \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {"product_id": 1, "quantity": 2}
    ]
  }'
```

### Get Order by ID:
```
curl -X GET http://localhost:8080/orders/1 \
  -H "Authorization: Basic YWRtaW46MTIzNA=="
```

### Update Order Status:
```
curl -X PATCH http://localhost:8080/orders/1 \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{"status": "completed"}'
```

## Testing the Message Queue Integration

### 1. Direct gRPC Testing

You can test the services using a gRPC client like [BloomRPC](https://github.com/bloomrpc/bloomrpc) or [grpcurl](https://github.com/fullstorydev/grpcurl).

#### Test Order Service:

1. Create a new order:

```bash
grpcurl -plaintext -d '{"user_id": 1, "items": [{"product_id": 1, "quantity": 2}]}' localhost:50052 pb.OrderService/CreateOrder
```

2. Get an order:

```bash
grpcurl -plaintext -d '{"id": 1}' localhost:50052 pb.OrderService/GetOrder
```

#### Test Inventory Service:

1. Create a new product:

```bash
grpcurl -plaintext -d '{"name": "Test Product", "description": "A test product", "price": 19.99, "stock": 100}' localhost:50051 pb.InventoryService/CreateProduct
```

2. Get a product:

```bash
grpcurl -plaintext -d '{"id": 1}' localhost:50051 pb.InventoryService/GetProduct
```

### 2. End-to-End Message Queue Testing

To test the entire message queue flow:

1. Create a product with initial stock through API Gateway or directly via gRPC:

```bash
curl -X POST http://localhost:8080/products \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Product", "description": "A test product", "price": 19.99, "stock": 10, "category_id": 1}'
```

2. Create an order that includes this product:

```bash
curl -X POST http://localhost:8080/orders \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {"product_id": 1, "quantity": 3}
    ]
  }'
```

3. Check the product stock to verify it was automatically updated:

```bash
curl -X GET http://localhost:8080/products/1 \
  -H "Authorization: Basic YWRtaW46MTIzNA=="
```

The stock should be reduced by the ordered quantity (10 - 3 = 7).

4. Monitor service logs to observe the message flow:
   - Order Service will show logs about publishing the order event
   - RabbitMQ will show the message being processed
   - Inventory Service will show logs about consuming the event and updating stock

## Logging

Each request is logged to the console with:

- IP address
- Method
- Endpoint
- Status code
- Duration

The RabbitMQ integration adds additional logging:
- Message publishing events
- Message consumption events
- Inventory updates triggered by messages

## HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Create successful
- `400 Bad Request`: Missing required parameters or invalid input
- `401 Unauthorized`: Authentication failed
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server-side error

## Error Handling

The system includes error handling strategies:

1. **Database Transactions**: Ensure data consistency
2. **Message Acknowledgements**: Prevent message loss
3. **Message Requeuing**: Retry processing of failed messages
4. **Comprehensive Logging**: For debugging and issue diagnosis

## Notes

- All API calls should go through the API Gateway
- The message queue handles inventory updates automatically when orders are created
- Services can run independently, providing resilience if one service is temporarily unavailable
- RabbitMQ must be running for full functionality of the order creation process

## Future Improvements

- Add circuit breakers for increased resilience
- Implement distributed tracing
- Add metrics collection for monitoring
- Implement the Outbox pattern for guaranteed message delivery
- Add user authentication and authorization
