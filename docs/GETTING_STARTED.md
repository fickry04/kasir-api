# Getting Started Guide - Kasir API

Panduan lengkap untuk memulai dan menggunakan Kasir API.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Running the Server](#running-the-server)
5. [First API Call](#first-api-call)
6. [Common Tasks](#common-tasks)
7. [Troubleshooting](#troubleshooting)

---

## Prerequisites

Pastikan Anda memiliki:

- **Go**: Version 1.25.6 atau lebih tinggi
  - Download dari https://golang.org/dl/
  - Verify installation: `go version`

- **PostgreSQL**: Database server
  - Download dari https://www.postgresql.org/download/
  - Atau gunakan **Neon Serverless**: https://neon.tech (recommended)

- **Git**: Version control system
  - Download dari https://git-scm.com/

- **Text Editor / IDE**: Untuk edit code
  - Recommended: VS Code, GoLand, atau IntelliJ IDEA

- **API Testing Tool**: Salah satu dari:
  - Postman: https://www.postman.com/
  - curl: Sudah built-in di kebanyakan OS
  - Thunder Client: VS Code extension

---

## Installation

### Step 1: Clone Repository

```bash
git clone <repository-url>
cd kasir-api
```

### Step 2: Verify Project Structure

```bash
tree -L 2
```

Expected output:
```
kasir-api/
├── cmd/
│   └── server/
├── docs/
├── internal/
├── go.mod
└── README.md
```

### Step 3: Download Go Dependencies

```bash
go mod download
```

Atau install dependencies dan tidy module:
```bash
go mod tidy
```

---

## Configuration

### Step 1: Setup Database

#### Option A: Neon Serverless (Recommended)

1. Buka https://neon.tech
2. Sign up untuk akun gratis
3. Buat project database baru
4. Copy connection string

#### Option B: Local PostgreSQL

1. Install PostgreSQL
2. Create database baru:
   ```bash
   createdb kasir_db
   ```
3. Connection string format:
   ```
   postgresql://username:password@localhost:5432/kasir_db
   ```

### Step 2: Create .env File

Buat file `.env` di root directory project:

```bash
touch .env
```

Isi file `.env` dengan connection string dan port:

```
DB_CONN=postgresql://username:password@host:5432/database
PORT=8080
```

**Contoh untuk Neon:**
```
DB_CONN=postgresql://user:password@ep-quiet-grass-123456.us-east-1.aws.neon.tech/neondb?sslmode=require
```

**Contoh untuk Local PostgreSQL:**
```
DB_CONN=postgresql://postgres:password@localhost:5432/kasir_db
```

### Step 3: Test Database Connection

Anda dapat test connection dengan membuka database client:
- pgAdmin: https://www.pgadmin.org/
- DBeaver: https://dbeaver.io/
- DataGrip: https://www.jetbrains.com/datagrip/

---

## Running the Server

### Start the Server

```bash
go run cmd/server/main.go
```

**Expected Output:**
```
Server running di localhost:{PORT}
```

### Verify Server is Running

Buka terminal baru dan jalankan:

```bash
curl http://localhost:8080/health
```

**Expected Response:**
```json
{"status":"OK","message":"Api Running"}
```

---

## First API Call

### 1. Check Health Status

```bash
curl http://localhost:8080/health
```

### 2. Get All Products

```bash
curl http://localhost:8080/api/products
```

Expected response (jika belum ada data):
```json
null
```

### 3. Create a Product

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "nama": "Kopi Arabika",
    "harga": 25000,
    "stok": 100
  }'
```

**Expected Response (201 Created):**
```json
{
  "id": 1,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

### 4. Get Product by ID

```bash
curl http://localhost:8080/api/products/1
```

**Expected Response:**
```json
{
  "id": 1,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

---

## Common Tasks

### Task 1: Create Multiple Products

```bash
# Product 1
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"nama": "Kopi Arabika", "harga": 25000, "stok": 100}'

# Product 2
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"nama": "Teh Pucuk", "harga": 8000, "stok": 50}'

# Product 3
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"nama": "Air Mineral", "harga": 5000, "stok": 200}'
```

### Task 2: Update Product

```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nama": "Kopi Arabika Premium",
    "harga": 30000,
    "stok": 150
  }'
```

### Task 3: Delete Product

```bash
curl -X DELETE http://localhost:8080/api/products/1
```

### Task 4: Manage Categories

Create category:
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minuman",
    "description": "Kategori minuman"
  }'
```

Get all categories:
```bash
curl http://localhost:8080/api/categories
```

---

## Using Postman

### Import API Collection (Optional)

Jika ada file Postman collection, import dengan:
1. Buka Postman
2. Click "Import"
3. Select file collection

### Manual Setup

1. **Health Check**
   - Method: GET
   - URL: `http://localhost:8080/health`
   - Click Send

2. **Get All Products**
   - Method: GET
   - URL: `http://localhost:8080/api/products`
   - Click Send

3. **Create Product**
   - Method: POST
   - URL: `http://localhost:8080/api/products`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "nama": "Kopi Arabika",
       "harga": 25000,
       "stok": 100
     }
     ```
   - Click Send

---

## Project Structure Overview

```
kasir-api/
├── cmd/                          # Command/Application entry point
│   └── server/
│       └── main.go              # Main server file
│
├── internal/                     # Internal packages (not exported)
│   ├── config/
│   │   └── database.go          # DB configuration
│   ├── handler/                 # HTTP handlers
│   │   ├── product_handler.go
│   │   └── category_handler.go
│   ├── model/                   # Data models
│   │   ├── product.go
│   │   └── category.go
│   ├── repository/              # Data access layer
│   │   ├── product_repository.go
│   │   └── category_repository.go
│   ├── router/                  # Route definitions
│   │   └── router.go
│   └── service/                 # Business logic
│       ├── product_service.go
│       └── category_service.go
│
├── docs/                        # API documentation
├── go.mod                       # Go module file
├── Dockerfile
├── .env                         # Environment for development
└── README.md                    # Project README
```

---


## Troubleshooting

### Problem: "Cannot find module"

**Error:**
```
go: missing go.sum entry
```

**Solution:**
```bash
go mod tidy
go mod download
```

---

### Problem: "Connection Refused"

**Error:**
```
Connection refused: connect ECONNREFUSED 127.0.0.1:8080
```

**Solution:**
1. Pastikan server sudah dijalankan: `go run cmd/server/main.go`
2. Check port 8080 availability
3. Gunakan port berbeda jika 8080 sudah terpakai

---

### Problem: "Database Connection Error"

**Error:**
```
Gagal menyambung ke database
```

**Solution:**
1. Verify `DB_CONN` di file `.env`
2. Test database connection string
3. Pastikan database server sedang berjalan
4. Check credentials (username, password)

---

### Problem: "Cannot decode request"

**Error:**
```
Invalid request
```

**Solution:**
1. Pastikan request body adalah valid JSON
2. Set header `Content-Type: application/json`
3. Verify semua required fields ada

---

## Getting Help

- Read [README.md](../README.md) untuk overview project
- Check individual endpoint documentation
- Look at code comments dan structure
- Consult Go documentation: https://golang.org/doc/

---
