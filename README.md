# Kasir API

Kasir API adalah REST API yang dibangun menggunakan Go untuk penugasan project bootcamp Jago Golang Dasar - CodeWithUmam.

## Notes

- Category service saat ini bersifat statik (in-memory) dan tidak terhubung langsung ke database — ini sengaja dibuat untuk tujuan pembelajaran. Data kategori hanya tersimpan selama aplikasi berjalan dan akan direset ketika server direstart.
- Product service terhubung dengan database PostgreSQL (contoh: Neon Serverless) untuk menunjukan bagaimana menyimpan data yang persisten. Pastikan `DATABASE_URL` di file `.env` dikonfigurasi jika ingin menyimpan produk ke database.
- Perilaku penyimpanan:
  - Produk: disimpan ke PostgreSQL (persisted)
  - Kategori: disimpan di memory (non-persisted)
- Semua endpoint mengembalikan JSON response

## Ada apa disini

- **CRUD Product**: CRUD operations untuk produk (Create, Read, Update, Delete)
- **CRUD Category**: CRUD operations untuk kategori (Create, Read, Update, Delete)
- **Health Check**: Endpoint untuk mengecek status API
- **Database Integration**: Terintegrasi dengan PostgreSQL (Neon Serverless)

## Tech Stack

- **Language**: Go 1.25.6
- **Database**: PostgreSQL (Neon Serverless)
- **Dependencies**:
  - `github.com/lib/pq` - PostgreSQL driver untuk Go
  - `github.com/joho/godotenv` - Environment variable loader

## Project Structure
Struktur project ini hanya hasil generate chatgpt :v, tidak ada kerangka yang tetap untuk struktur project golang.
 
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

### Build (Compile) untuk Linux & Windows

Anda bisa meng-compile aplikasi Go menjadi binary untuk platform target. Contoh-perintah di bawah menghasilkan binary statis untuk `linux` dan `windows`.

Build untuk Linux (amd64):
```bash
GOOS=linux GOARCH=amd64 go build -o bin/kasir-linux ./cmd/server
```

Build untuk Windows (amd64):
```bash
GOOS=windows GOARCH=amd64 go build -o bin/kasir-windows.exe ./cmd/server
```

## Ringkasan API Endpoints (Buka docs untuk lebih lengkap)

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
- **KeyRunner** - VS Code extension

## Environment Variables

| Variable | Description |
|----------|-------------|
| `DATABASE_URL` | PostgreSQL connection string |

Sekian dan Terima Kasih.
