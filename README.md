# Wallet Topup Service

A Go-based wallet topup service following Clean Architecture principles.

## Project Structure

```
├── cmd/
│   └── server/              # main.go
├── config/                  # .env, config loader
├── internal/
│     ├── api/
│     │   └── wallet/                  # Wallet API module
│     │       ├── handlers/            # Request handlers
│     │       ├── interface.go         # Service and repository interfaces
│     │       ├── repo/                # Repository implementations
│     │       ├── routes/              # Route definitions
│     │       └── services/            # Business logic implementations
│     ├── database/                    # Database configuration and migrations
│     ├── entities/                   # Domain entities and models
│     │   ├── transaction.go         # Transaction entity
│     │   └── user.go               # User entity
│     ├── global/                    # Global configurations and utilities
│     │   └── responses/         # Standard API responses
├── test/                    # Unit tests
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Features

- Wallet topup functionality
- Transaction verification and confirmation
- User balance management
- Clean Architecture implementation
- PostgreSQL database integration
- RESTful API endpoints

## API Endpoints

### Wallet
- `POST /wallet/verify` - Verify a transaction
- `POST /wallet/confirm/:transaction_id` - Confirm a transaction


## Example API Requests
 Verify a transaction
```
  curl --location 'http://localhost:8080/wallet/verify' \
--header 'Content-Type: application/json' \
--data '{"user_id": 1, "amount": 100.50, "payment_method": "credit_card"}'

```
Confirm a transaction    
```
    curl --location --request POST 'http://localhost:8080/wallet/confirm/1139b5bc-704d-4afe-894c-1480fc0e6e5d'
```

## Getting Started

1. Clone the repository
2. Install dependencies
3. Set up the database
4. Run migrations
5. Start the server

## Dependencies

- Go 1.21+
- PostgreSQL
- Gin Web Framework
- GORM ORM
- Goose Migrations

## License

MIT