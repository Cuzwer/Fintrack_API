# FinTrack API

> Personal Budget & Expense Tracking System — Backend REST API

Built with **Go**, **Fiber**, **PostgreSQL** and deployed with **Docker**

---

## 📌 Overview

FinTrack API คือระบบ Backend API สำหรับจัดการการเงินส่วนตัว ช่วยให้ผู้ใช้สามารถติดตามรายรับรายจ่าย ตั้งงบประมาณรายเดือน และดูสรุปภาพรวมทางการเงินได้ครับ

---

## 🛠 Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Framework | Fiber |
| Database | PostgreSQL |
| Query | Raw SQL |
| Authentication | JWT + HTTP Cookie |
| Container | Docker + Docker Compose |
| Documentation | Swagger |
| Testing | Go testing + testify |

---

## ✨ Features

### Core Features
- ✅ User registration / login (JWT + HTTP Cookie)
- ❌ Account management (บัญชีธนาคาร, เงินสด, บัตรเครดิต)
- ❌ Income / Expense transaction
- ❌ Category management
- ❌ Monthly budget setting
- ❌ Balance calculation

### Nice to Have
- ⭐ Transaction summary report
- ⭐ Budget vs Actual comparison
- ⭐ Recurring transaction

---

## 📁 Project Structure

```
fintrack-api/
│
├── cmd/
│   └── main.go                    # entry point
│
├── internal/
│   ├── domain/                    # Business entities
│   │   ├── user.go
│   │   ├── account.go
│   │   ├── transaction.go
│   │   ├── category.go
│   │   └── budget.go
│   │
│   ├── repository/                # Raw SQL queries
│   │   ├── user_repo.go
│   │   ├── account_repo.go
│   │   ├── transaction_repo.go
│   │   ├── category_repo.go
│   │   └── budget_repo.go
│   │
│   ├── service/                   # Business logic
│   │   ├── user_service.go
│   │   ├── account_service.go
│   │   ├── transaction_service.go
│   │   ├── category_service.go
│   │   └── budget_service.go
│   │
│   ├── handler/                   # HTTP handlers
│   │   ├── user_handler.go
│   │   ├── account_handler.go
│   │   ├── transaction_handler.go
│   │   ├── category_handler.go
│   │   └── budget_handler.go
│   │
│   └── middleware/
│       └── auth.go                # JWT middleware
│
├── pkg/
│   ├── database/
│   │   └── postgres.go            # DB connection
│   ├── config/
│   │   └── config.go              # อ่าน .env
│   └── utils/
│       ├── hash.go                # bcrypt password
│       └── jwt.go                 # generate/validate token
│
├── migrations/
│   ├── 001_create_users.sql
│   ├── 002_create_accounts.sql
│   ├── 003_create_categories.sql
│   ├── 004_create_transactions.sql
│   └── 005_create_budgets.sql
│
├── docs/                          # Swagger auto-generated
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
```

---

## 🗄 Database Schema

```sql
CREATE TABLE users (
  id_user       SERIAL PRIMARY KEY,
  email_user    VARCHAR(100) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
  id_account   SERIAL PRIMARY KEY,
  id_user      INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  name         VARCHAR(100) NOT NULL,
  type_account VARCHAR(50) NOT NULL,
  balance      DECIMAL(15, 2) DEFAULT 0.00,
  currency     VARCHAR(10) DEFAULT 'THB'
);

CREATE TABLE categories (
  id_category   SERIAL PRIMARY KEY,
  id_user       INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  name_category VARCHAR(100) NOT NULL,
  type_category VARCHAR(20) NOT NULL CHECK (type_category IN ('income', 'expense'))
);

CREATE TABLE transactions (
  id_trans         SERIAL PRIMARY KEY,
  id_account       INT NOT NULL REFERENCES accounts(id_account) ON DELETE CASCADE,
  id_category      INT REFERENCES categories(id_category),
  amount_trans     DECIMAL(15, 2) NOT NULL,
  type_trans       VARCHAR(20) NOT NULL CHECK (type_trans IN ('income', 'expense')),
  descrip_trans    VARCHAR(255),
  transaction_date DATE NOT NULL DEFAULT CURRENT_DATE,
  created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE budgets (
  id_budg     SERIAL PRIMARY KEY,
  id_user     INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  id_category INT NOT NULL REFERENCES categories(id_category),
  amount_budg DECIMAL(15, 2) NOT NULL,
  month_budg  INT NOT NULL CHECK (month_budg BETWEEN 1 AND 12),
  year_budg   INT NOT NULL
);
```

---

## 🔌 API Endpoints

### Auth
| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/auth/register` | สมัครสมาชิก |
| POST | `/api/v1/auth/login` | เข้าสู่ระบบ |
| POST | `/api/v1/auth/logout` | ออกจากระบบ |

### Accounts
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/accounts` | ดูบัญชีทั้งหมด |
| POST | `/api/v1/accounts` | สร้างบัญชีใหม่ |
| PUT | `/api/v1/accounts/:id` | แก้ไขบัญชี |
| DELETE | `/api/v1/accounts/:id` | ลบบัญชี |

### Transactions
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/transactions` | ดู transaction ทั้งหมด |
| POST | `/api/v1/transactions` | บันทึก transaction |
| PUT | `/api/v1/transactions/:id` | แก้ไข transaction |
| DELETE | `/api/v1/transactions/:id` | ลบ transaction |

### Categories
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/categories` | ดู category ทั้งหมด |
| POST | `/api/v1/categories` | สร้าง category |
| PUT | `/api/v1/categories/:id` | แก้ไข category |
| DELETE | `/api/v1/categories/:id` | ลบ category |

### Budgets
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/budgets` | ดู budget ทั้งหมด |
| POST | `/api/v1/budgets` | ตั้ง budget |
| PUT | `/api/v1/budgets/:id` | แก้ไข budget |
| DELETE | `/api/v1/budgets/:id` | ลบ budget |

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Docker + Docker Compose
- PostgreSQL 15+

### Run with Docker

```bash
# 1. clone project
git clone https://github.com/yourusername/fintrack-api.git
cd fintrack-api

# 2. copy env
cp .env.example .env

# 3. run
docker-compose up -d
```

### Run locally

```bash
# install dependencies
go mod tidy

# run migration
psql -U postgres -d fintrack -f migrations/001_create_users.sql
# ... run all migrations

# start server
go run cmd/main.go
```

---

## 🏗 Architecture

ใช้ **Clean Architecture** แยก layer ชัดเจนครับ

```
HTTP Request
    ↓
Handler     → รับ request, validate input, ส่ง response
    ↓
Service     → business logic ทั้งหมด
    ↓
Repository  → Raw SQL query
    ↓
PostgreSQL
```

แต่ละ layer รู้จักแค่ layer ข้างล่างตัวเองเท่านั้น ทำให้ง่ายต่อการ test และ maintain ครับ

---

## 🔐 Authentication Flow

```
Register → bcrypt password → store in DB
Login    → verify password → generate JWT → set HTTP Cookie
Request  → middleware ตรวจ Cookie → validate JWT → allow/deny
Logout   → clear Cookie
```

---

## 🧪 Testing

```bash
go test ./...

# with coverage
go test ./... -cover
```

---

## 📦 Environment Variables

```env
# Server
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=fintrack

# JWT
JWT_SECRET=your_secret_key
JWT_EXPIRE_HOURS=24
```

---

## 👤 Author
