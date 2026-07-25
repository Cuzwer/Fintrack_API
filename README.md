# FinTrack API

> Personal Budget & Expense Tracking System — Backend REST API

A powerful and easy-to-use financial management backend built with **Go**, **Fiber**, and **PostgreSQL**. Track your income, expenses, accounts, categories, and monthly budgets all in one place.

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Prerequisites](#-prerequisites)
- [Getting Started](#-getting-started)
  - [Clone the Project](#1-clone-the-project)
  - [Setup Environment](#2-setup-environment)
  - [Run with Docker](#3-run-with-docker)
  - [Run Locally](#4-run-locally)
- [API Endpoints](#-api-endpoints)
- [Project Structure](#-project-structure)
- [Environment Variables](#-environment-variables)
- [API Documentation](#-api-documentation)

---

## ✨ Features

- **User Authentication** — Register, login, and logout with JWT tokens stored in HTTP-only cookies
- **Account Management** — Track multiple accounts (bank accounts, cash, credit cards)
- **Transaction Tracking** — Record income and expenses with categories
- **Category Management** — Organize transactions by custom categories
- **Budget Planning** — Set monthly budgets per category and track spending
- **Balance Calculation** — Automatic balance updates for all accounts
- **Swagger Documentation** — Interactive API docs available at `/swagger/*`

---

## 🛠 Tech Stack

| Category | Technology |
|----------|------------|
| Language | Go 1.26+ |
| Framework | Fiber v2 |
| Database | PostgreSQL |
| ORM | GORM |
| Authentication | JWT + HTTP Cookie |
| API Docs | Swagger |
| Migration | golang-migrate |
| Container | Docker + Docker Compose |

---

## 📌 Prerequisites

Before you start, make sure you have:

- **Go** 1.21 or higher
- **PostgreSQL** 15+ (or Docker)
- **Docker** & **Docker Compose** (for containerized setup)
- **Git**

---

## 🚀 Getting Started

### 1. Clone the Project

```bash
git clone https://github.com/Cuzwer/Fintrack_API.git
cd Fintrack_API
```

### 2. Setup Environment

Create a `.env` file in the project root with the following variables:

```env
# Server Configuration
APP_PORT=7070

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=fintrack
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your_super_secret_key_here
JWT_EXPIRE_HOURS=24
```

> **Note:** If you're using Docker, the database credentials in `.env` must match those in `docker-compose.yml`.

### 3. Run with Docker

The easiest way to get started:

```bash
# Build and start all services (API + Database)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

The API will be available at `http://localhost:7070`

### 4. Run Locally

#### Option A: With Docker Database Only

```bash
# Start only PostgreSQL
docker-compose up -d db

# Install dependencies
go mod tidy

# Run the API
go run ./cmd/main.go
```

#### Option B: Local PostgreSQL

```bash
# Create the database
psql -U postgres -c "CREATE DATABASE fintrack;"

# Run migrations
psql -U postgres -d fintrack -f migrations/001_create_users_table.up.sql
psql -U postgres -d fintrack -f migrations/002_create_accounts_table.up.sql
psql -U postgres -d fintrack -f migrations/003_create_categories_table.up.sql
psql -U postgres -d fintrack -f migrations/004_create_transactions_table.up.sql
psql -U postgres -d fintrack -f migrations/005_create_budgets_table.up.sql

# Install dependencies
go mod tidy

# Start the server
go run ./cmd/main.go
```

#### Option C: Development Mode (with auto-reload)

```bash
# Install nodemon for Go
go install github.com/nickromo/nodemon@latest

# Run with hot-reload
make dev
# or
nodemon --watch . --ext go --exec "go run ./cmd/main.go"
```

---

## 🔌 API Endpoints

### Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/user/register` | Register a new user |
| POST | `/api/v1/user/login` | Login and get JWT cookie |
| DELETE | `/api/v1/user/{id}` | Delete user account |

### Accounts
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/accounts/:id` | Get all accounts for a user |
| POST | `/api/v1/accounts` | Create a new account |

### Transactions
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/transactions` | Get all transactions |
| POST | `/api/v1/transactions` | Create a new transaction |

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/categories` | Get all categories |
| POST | `/api/v1/categories` | Create a new category |
| PUT | `/api/v1/categories/:id` | Update a category |
| DELETE | `/api/v1/categories/:id` | Delete a category |

### Budgets
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/budgets` | Get all budgets |
| POST | `/api/v1/budgets` | Create a new budget |
| PUT | `/api/v1/budgets/:id` | Update a budget |
| DELETE | `/api/v1/budgets/:id` | Delete a budget |

### Documentation
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/swagger/*` | Swagger API Documentation |

---

## 📁 Project Structure

```
Fintrack_API/
├── cmd/
│   └── main.go              # Application entry point
│
├── internal/
│   ├── domain/              # Data models/entities
│   │   ├── user.go
│   │   ├── account.go
│   │   ├── transaction.go
│   │   ├── categories.go
│   │   └── budgets.go
│   │
│   ├── repository/          # Database queries (Raw SQL)
│   │   ├── user_repo.go
│   │   ├── account_repo.go
│   │   ├── transactionRepo.go
│   │   ├── categories_repo.go
│   │   └── budget_repo.go
│   │
│   ├── service/             # Business logic layer
│   │   ├── user_service.go
│   │   ├── account_service.go
│   │   ├── transactionService.go
│   │   ├── categories_service.go
│   │   └── budget_sevice.go
│   │
│   ├── handler/             # HTTP request handlers
│   │   ├── user_handler.go
│   │   ├── account_handler.go
│   │   ├── transaction_handler.go
│   │   ├── categories_handler.go
│   │   └── budget_hanlder.go
│   │
│   ├── middleware/         # Custom middleware
│   │   ├── auth.go         # JWT authentication
│   │   ├── cors.go
│   │   └── limiter.go
│   │
│   └── routes/              # Route definitions
│       ├── router.go
│       ├── user_routes.go
│       ├── account_routes.go
│       ├── transaction_routes.go
│       ├── categories_routes.go
│       └── budget_routes.go
│
├── pkg/
│   ├── database/            # Database connection
│   │   └── postgres.go
│   ├── config/             # Configuration loader
│   │   └── config.go
│   └── utils/               # Utility functions
│       └── hash.go
│
├── migrations/              # SQL migration files
│   ├── 001_create_users_table.up.sql
│   ├── 002_create_accounts_table.up.sql
│   ├── 003_create_categories_table.up.sql
│   ├── 004_create_transactions_table.up.sql
│   └── 005_create_budgets_table.up.sql
│
├── docs/                    # Swagger generated docs
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
└── README.md
```

---

## 🔐 Authentication Flow

```
┌─────────────────────────────────────────────────────────┐
│                    AUTHENTICATION FLOW                  │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  REGISTER                                                │
│  ├─ User submits email & password                       │
│  ├─ Password hashed with bcrypt                        │
│  └─ User saved to database                              │
│                                                         │
│  LOGIN                                                  │
│  ├─ User submits credentials                            │
│  ├─ Password verified against hash                      │
│  ├─ JWT token generated                                 │
│  └─ Token stored in HTTP-only cookie                   │
│                                                         │
│  PROTECTED ROUTES                                       │
│  ├─ Middleware reads JWT from cookie                    │
│  ├─ Token validated & user ID extracted                │
│  └─ Request proceeds or 401 returned                    │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

---

## 🗄 Database Schema

```sql
-- Users table
CREATE TABLE users (
    id_user       SERIAL PRIMARY KEY,
    email_user    VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Accounts table (bank accounts, cash, credit cards)
CREATE TABLE accounts (
    id_account   SERIAL PRIMARY KEY,
    id_user      INT NOT NULL REFERENCES users(id_user),
    name         VARCHAR(100) NOT NULL,
    type_account VARCHAR(50) NOT NULL,
    balance      DECIMAL(15, 2) DEFAULT 0.00,
    currency     VARCHAR(10) DEFAULT 'THB'
);

-- Categories table (income/expense categories)
CREATE TABLE categories (
    id_category   SERIAL PRIMARY KEY,
    id_user       INT NOT NULL REFERENCES users(id_user),
    name_category VARCHAR(100) NOT NULL,
    type_category VARCHAR(20) NOT NULL CHECK (type_category IN ('income', 'expense'))
);

-- Transactions table
CREATE TABLE transactions (
    id_trans         SERIAL PRIMARY KEY,
    id_account       INT NOT NULL REFERENCES accounts(id_account),
    id_category      INT REFERENCES categories(id_category),
    amount_trans     DECIMAL(15, 2) NOT NULL,
    type_trans       VARCHAR(20) NOT NULL CHECK (type_trans IN ('income', 'expense')),
    descrip_trans    VARCHAR(255),
    transaction_date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Budgets table
CREATE TABLE budgets (
    id_budg     SERIAL PRIMARY KEY,
    id_user     INT NOT NULL REFERENCES users(id_user),
    id_category INT NOT NULL REFERENCES categories(id_category),
    amount_budg DECIMAL(15, 2) NOT NULL,
    month_budg  INT NOT NULL CHECK (month_budg BETWEEN 1 AND 12),
    year_budg   INT NOT NULL
);
```

---

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Server port | `7070` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database username | `postgres` |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | `fintrack` |
| `DB_SSLMODE` | SSL mode | `disable` |
| `JWT_SECRET` | Secret key for JWT signing | - |
| `JWT_EXPIRE_HOURS` | Token expiration in hours | `24` |

---

## 📚 API Documentation

Once the server is running, visit:

- **Swagger UI:** `http://localhost:7070/swagger/`

This provides an interactive interface to test all API endpoints.

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests with verbose output
go test -v ./...
```

---

## 🏗 Architecture

This project follows **Clean Architecture** principles:

```
HTTP Request
     │
     ▼
┌─────────────┐
│   Handler   │  →  Input validation, Response formatting
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Service   │  →  Business logic
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Repository  │  →  Database queries
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ PostgreSQL  │  →  Data storage
└─────────────┘
```

Each layer only knows about the layer directly below it, making the code modular, testable, and maintainable.

---

## 👤 Author

**Azfar Mamu**

---

## 📄 License

This project is open source and available under the MIT License.

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](https://github.com/Cuzwer/Fintrack_API/issues).

---

<p align="center">
  <strong>Happy tracking your finances! 💰</strong>
</p>
