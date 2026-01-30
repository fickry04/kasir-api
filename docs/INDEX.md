# Kasir API - Documentation Index

Selamat datang di dokumentasi Kasir API! Dokumentasi ini berisi semua informasi yang Anda butuhkan untuk menggunakan dan mengembangkan API ini.

## 📚 Documentation Structure

### Getting Started
**[GETTING_STARTED.md](GETTING_STARTED.md)** - Start here!
- Prerequisites & installation
- Configuration & setup
- Running the server
- First API calls
- Common tasks & troubleshooting

---

## 🔌 API Endpoint Documentation

### Health Check
**[HEALTH_CHECK_API.md](HEALTH_CHECK_API.md)**
- Monitoring API status
- Implementation examples
- Best practices

**Endpoint:**
```
GET /health
```

---

### Products Management
**[PRODUCTS_API.md](PRODUCTS_API.md)**

Complete CRUD operations untuk Product:

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/products` | Dapatkan semua produk |
| POST | `/api/products` | Buat produk baru |
| GET | `/api/products/{id}` | Dapatkan produk by ID |
| PUT | `/api/products/{id}` | Update produk |
| DELETE | `/api/products/{id}` | Hapus produk |

**Product Object:**
```json
{
  "id": 1,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

---

### Categories Management
**[CATEGORIES_API.md](CATEGORIES_API.md)**

Complete CRUD operations untuk Category:

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/categories` | Dapatkan semua kategori |
| POST | `/api/categories` | Buat kategori baru |
| GET | `/api/categories/{id}` | Dapatkan kategori by ID |
| PUT | `/api/categories/{id}` | Update kategori |
| DELETE | `/api/categories/{id}` | Hapus kategori |

**⚠️ Important:** Category service adalah static/memory-based (untuk learning)

**Category Object:**
```json
{
  "id": 1,
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

---

## 🏗️ Architecture Overview

```
HTTP Request
    ↓
Handler Layer (HTTP handling)
    ↓
Service Layer (Business logic)
    ↓
Repository Layer (Data access)
    ↓
PostgreSQL Database
```

---

## 📝 Quick Reference

### Installation Commands

```bash
# Clone repository
git clone <repository-url>
cd kasir-api

# Download dependencies
go mod download

# Create .env file
echo "DB_CONN=your_connection_string" > .env

# Run server
go run cmd/server/main.go
```

### Basic curl Commands

```bash
# Health check
curl http://localhost:8080/health

# Get all products
curl http://localhost:8080/api/products

# Create product
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"nama":"Kopi","harga":25000,"stok":100}'

# Get product by ID
curl http://localhost:8080/api/products/1

# Update product
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Kopi Premium","harga":30000,"stok":150}'

# Delete product
curl -X DELETE http://localhost:8080/api/products/1
```

---

## 🛠️ Tech Stack

| Component | Version | Purpose |
|-----------|---------|---------|
| Go | 1.25.6 | Programming language |
| PostgreSQL | Latest | Database |
| lib/pq | 1.10.9 | PostgreSQL driver |
| viper | 1.21.0 | Environment variables |

---

## 📋 Feature List

### Implemented
- ✅ Product CRUD operations (Database-connected)
- ✅ Category CRUD operations (Static/Memory-based)
- ✅ Health check endpoint
- ✅ JSON request/response
- ✅ Error handling
- ✅ Environment configuration

### Future Enhancements
- 🔄 Category database integration
- 🔄 Input validation
- 🔄 Authentication & authorization
- 🔄 Rate limiting
- 🔄 Logging & monitoring
- 🔄 API versioning
- 🔄 Database migrations

---

## 🔍 Common Use Cases

### 1. Verify API is Running
```bash
curl http://localhost:8080/health
```

### 2. View All Products
```bash
curl http://localhost:8080/api/products
```

### 3. Add New Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"nama":"Produk Baru","harga":50000,"stok":10}'
```

### 4. Update Stock
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Kopi","harga":25000,"stok":75}'
```

### 5. Remove Discontinued Product
```bash
curl -X DELETE http://localhost:8080/api/products/1
```

---

## 🐛 Troubleshooting Quick Guide

| Issue | Solution |
|-------|----------|
| "Connection refused" | Server belum running, jalankan `go run cmd/server/main.go` |
| "Database connection failed" | Check DB_CONN di .env, verify database server berjalan |
| "Invalid request body" | Pastikan JSON valid dan Content-Type header benar |
| "Product not found" | Verify product ID exists |
| "Port 8080 already in use" | Ubah port di main.go atau stop aplikasi lain menggunakan port 8080 |

---

## 📞 Support & Resources

### Documentation Files
- [Getting Started Guide](GETTING_STARTED.md) - Untuk pemula
- [Products API](PRODUCTS_API.md) - Dokumentasi products endpoint
- [Categories API](CATEGORIES_API.md) - Dokumentasi categories endpoint
- [Health Check API](HEALTH_CHECK_API.md) - Status monitoring
- [Main README](../README.md) - Project overview

### External Resources
- Go Documentation: https://golang.org/doc/
- PostgreSQL Manual: https://www.postgresql.org/docs/
- RESTful API Design: https://restfulapi.net/
- HTTP Status Codes: https://httpwg.org/specs/rfc7231.html#status.codes

---

## 📅 Documentation Info

- **Last Updated**: January 2026
- **API Version**: 1.0
- **Documentation Version**: 1.0
- **Status**: Active

---

## 🎯 Documentation Roadmap

- [x] Getting Started Guide
- [x] Products API Documentation
- [x] Categories API Documentation
- [x] Health Check Documentation
- [x] Documentation Index
- [ ] Code examples untuk berbagai bahasa
- [ ] Database schema documentation
- [ ] Deployment guide

---

## 💡 Tips

1. **Mulai dari sini**: Baca [GETTING_STARTED.md](GETTING_STARTED.md) terlebih dahulu
2. **Test API**: Gunakan curl atau Postman untuk test endpoints
3. **Read Examples**: Setiap dokumentasi endpoint berisi contoh penggunaan
4. **Check Errors**: Baca section "Common Errors" di setiap endpoint
5. **Ask Questions**: Jika ada yang tidak jelas, refer ke documentation atau check code comments

---

## 🚀 Getting Started Now

### Quick Start (5 minutes)

1. Clone repo: `git clone <url>`
2. Setup .env: `echo "DB_CONN=your_connection_string" > .env`
3. Run server: `go run cmd/server/main.go`
4. Test health: `curl http://localhost:8080/health`
5. Create product: `curl -X POST http://localhost:8080/api/products -H "Content-Type: application/json" -d '{"nama":"Kopi","harga":25000,"stok":100}'`

👉 **Ready?** [Go to Getting Started Guide →](GETTING_STARTED.md)

---

**Questions?** Refer ke dokumentasi yang relevan atau check code comments.

Happy coding! 🎉
