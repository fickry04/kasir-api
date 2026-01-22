# Kasir API

Kasir API adalah REST API yang dibangun menggunakan Go untuk sistem manajemen kasir/penjualan. API ini menyediakan endpoint untuk mengelola produk dan kategori dengan database PostgreSQL.

## Fitur Utama

- **Product Management**: CRUD operations untuk produk (Create, Read, Update, Delete)
- **Category Management**: CRUD operations untuk kategori (Create, Read, Update, Delete)
- **Health Check**: Endpoint untuk mengecek status API
- **Database Integration**: Terintegrasi dengan PostgreSQL (Neon Serverless)

## Tech Stack

- **Language**: Go 1.25.6
- **Database**: PostgreSQL (Neon Serverless)
- **Dependencies**:
  - `github.com/lib/pq` - PostgreSQL driver untuk Go
  - `github.com/joho/godotenv` - Environment variable loader

## Instalasi & Setup

### Prerequisites
- Go 1.25.6 atau lebih tinggi
- PostgreSQL database (atau Neon Serverless)
- Git

### Step-by-Step

1. **Clone repository**
```bash
git clone <repository-url>
cd kasir-api
```

2. **Setup environment variables**
Buat file `.env` di root directory:
```
DATABASE_URL=postgresql://username:password@host:port/database
```

3. **Install dependencies**
```bash
go mod download
```

4. **Run server**
```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

## Project Structure

```
kasir-api/
├── cmd/
│   └── server/
│       └── main.go           # Entry point aplikasi
├── internal/
│   ├── config/
│   │   └── database.go       # Database configuration
│   ├── handler/
│   │   ├── product_handler.go
│   │   └── category_handler.go
│   ├── model/
│   │   ├── product.go
│   │   └── category.go
│   ├── repository/
│   │   ├── product_repository.go
│   │   └── category_repository.go
│   ├── router/
│   │   └── router.go         # Route definitions
│   └── service/
│       ├── product_service.go
│       └── category_service.go
├── docs/                      # API Documentation
├── go.mod
└── README.md
```

## Architecture

Aplikasi ini mengikuti clean architecture pattern dengan layers:

1. **Handler Layer**: Menangani HTTP requests dan responses
2. **Service Layer**: Business logic dan validasi
3. **Repository Layer**: Data access layer untuk database operations
4. **Model Layer**: Data structures/entities

## API Endpoints

### Health Check
- **GET** `/health` - Check API status

### Products
- **GET** `/api/products` - Dapatkan semua produk
- **POST** `/api/products` - Buat produk baru
- **GET** `/api/products/{id}` - Dapatkan produk berdasarkan ID
- **PUT** `/api/products/{id}` - Update produk
- **DELETE** `/api/products/{id}` - Hapus produk

### Categories
- **GET** `/api/categories` - Dapatkan semua kategori
- **POST** `/api/categories` - Buat kategori baru
- **GET** `/api/categories/{id}` - Dapatkan kategori berdasarkan ID
- **PUT** `/api/categories/{id}` - Update kategori
- **DELETE** `/api/categories/{id}` - Hapus kategori

Untuk dokumentasi lengkap setiap endpoint, lihat folder [docs/](docs/)

## Request/Response Examples

### Product Response Format
```json
{
  "id": 1,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

### Category Response Format
```json
{
  "id": 1,
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

## Development

### Struktur Error Handling
- `400 Bad Request` - Invalid request format atau ID
- `404 Not Found` - Resource tidak ditemukan
- `201 Created` - Resource berhasil dibuat
- `200 OK` - Request berhasil

### Menjalankan Project
```bash
# Terminal 1 - Run server
go run cmd/server/main.go

# Terminal 2 - Test API (contoh dengan curl)
curl http://localhost:8080/api/products
```

## Testing

Untuk testing API, Anda dapat menggunakan:
- **Postman** - GUI client untuk HTTP requests
- **curl** - Command line tool
- **Thunder Client** - VS Code extension

## Environment Variables

| Variable | Description |
|----------|-------------|
| `DATABASE_URL` | PostgreSQL connection string |

## Contributing

1. Create feature branch (`git checkout -b feature/amazing-feature`)
2. Commit changes (`git commit -m 'Add amazing feature'`)
3. Push to branch (`git push origin feature/amazing-feature`)
4. Open Pull Request

## Notes

- Category service saat ini bersifat statik dan tidak terhubung langsung ke database (untuk pembelajaran)
- Product service terhubung dengan database PostgreSQL
- Semua endpoint mengembalikan JSON response

## License

Proyek ini adalah untuk pembelajaran.

## Contact & Support

Untuk pertanyaan atau issues, silakan buat issue di repository ini.

---

Last Updated: January 2026
