# API Dokumentasi - Products Endpoints

Dokumentasi lengkap untuk semua endpoint yang berkaitan dengan Product (Produk).

## Base URL
```
http://localhost:8080
```

## Endpoints Overview

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/products` | Dapatkan semua produk |
| POST | `/api/products` | Buat produk baru |
| GET | `/api/products/{id}` | Dapatkan produk berdasarkan ID |
| PUT | `/api/products/{id}` | Update produk |
| DELETE | `/api/products/{id}` | Hapus produk |

---

## 1. Get All Products

Mengambil daftar semua produk yang ada di database.

**Request:**
```
GET /api/products
```

**Parameters:** None

**Headers:**
```
Content-Type: application/json
```

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "nama": "Kopi Arabika",
    "harga": 25000,
    "stok": 100
  },
  {
    "id": 2,
    "nama": "Teh Pucuk",
    "harga": 8000,
    "stok": 50
  }
]
```

**Status Codes:**
- `200 OK` - Request berhasil
- `500 Internal Server Error` - Error database

**Example dengan curl:**
```bash
curl -X GET http://localhost:8080/api/products \
  -H "Content-Type: application/json"
```

---

## 2. Create Product

Membuat produk baru di database.

**Request:**
```
POST /api/products
```

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

**Parameters:**

| Field | Type | Required | Deskripsi |
|-------|------|----------|-----------|
| nama | string | Ya | Nama produk |
| harga | integer | Ya | Harga produk |
| stok | integer | Ya | Jumlah stok produk |

**Response (201 Created):**
```json
{
  "id": 0,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

**Status Codes:**
- `201 Created` - Produk berhasil dibuat
- `400 Bad Request` - Invalid request format
- `500 Internal Server Error` - Error database

**Example dengan curl:**
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "nama": "Kopi Arabika",
    "harga": 25000,
    "stok": 100
  }'
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/products', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    nama: 'Kopi Arabika',
    harga: 25000,
    stok: 100
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## 3. Get Product by ID

Mengambil detail produk berdasarkan ID.

**Request:**
```
GET /api/products/{id}
```

**URL Parameters:**

| Parameter | Type | Required | Deskripsi |
|-----------|------|----------|-----------|
| id | integer | Ya | ID produk |

**Headers:**
```
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "id": 1,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

**Status Codes:**
- `200 OK` - Request berhasil
- `400 Bad Request` - ID tidak valid (bukan integer)
- `404 Not Found` - Produk tidak ditemukan
- `500 Internal Server Error` - Error database

**Error Response (400):**
```json
Invalid Produk ID
```

**Error Response (404):**
```json
Product not found
```

**Example dengan curl:**
```bash
curl -X GET http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json"
```

---

## 4. Update Product

Mengupdate data produk yang sudah ada.

**Request:**
```
PUT /api/products/{id}
```

**URL Parameters:**

| Parameter | Type | Required | Deskripsi |
|-----------|------|----------|-----------|
| id | integer | Ya | ID produk yang akan diupdate |

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "nama": "Kopi Arabika Premium",
  "harga": 30000,
  "stok": 150
}
```

**Parameters:**

| Field | Type | Required | Deskripsi |
|-------|------|----------|-----------|
| nama | string | Ya | Nama produk |
| harga | integer | Ya | Harga produk |
| stok | integer | Ya | Jumlah stok produk |

**Response (200 OK):**
```json
{
  "id": 1,
  "nama": "Kopi Arabika Premium",
  "harga": 30000,
  "stok": 150
}
```

**Status Codes:**
- `200 OK` - Produk berhasil diupdate
- `400 Bad Request` - ID tidak valid atau invalid request body
- `500 Internal Server Error` - Error database

**Error Response (400):**
```json
Invalid Product ID
```

**Example dengan curl:**
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nama": "Kopi Arabika Premium",
    "harga": 30000,
    "stok": 150
  }'
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/products/1', {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    nama: 'Kopi Arabika Premium',
    harga: 30000,
    stok: 150
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## 5. Delete Product

Menghapus produk dari database.

**Request:**
```
DELETE /api/products/{id}
```

**URL Parameters:**

| Parameter | Type | Required | Deskripsi |
|-----------|------|----------|-----------|
| id | integer | Ya | ID produk yang akan dihapus |

**Headers:**
```
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "message": "Sukses delete"
}
```

**Status Codes:**
- `200 OK` - Produk berhasil dihapus
- `400 Bad Request` - ID tidak valid
- `404 Not Found` - Produk tidak ditemukan
- `500 Internal Server Error` - Error database

**Error Response (400):**
```json
Invalid Produk ID
```

**Error Response (404):**
```json
Product not found
```

**Example dengan curl:**
```bash
curl -X DELETE http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json"
```

---

## Data Model

### Product Object

```json
{
  "id": 1,
  "nama": "Kopi Arabika",
  "harga": 25000,
  "stok": 100
}
```

**Field Descriptions:**

| Field | Type | Deskripsi |
|-------|------|-----------|
| id | integer | ID unik produk (auto-generated) |
| nama | string | Nama produk |
| harga | integer | Harga produk dalam Rupiah |
| stok | integer | Jumlah stok produk yang tersedia |

---

## Common Errors

### 400 Bad Request
```
Invalid request format atau parameter yang dikirim tidak sesuai format yang diharapkan
```

**Solutions:**
- Pastikan request body adalah valid JSON
- Pastikan semua required fields ada
- Pastikan data types sesuai (string, integer, dll)

### 404 Not Found
```
Resource (produk) tidak ditemukan di database
```

**Solutions:**
- Verifikasi ID produk yang digunakan
- Pastikan produk sudah ada di database sebelum mengupdate/menghapus

### 500 Internal Server Error
```
Terjadi error di server atau database
```

**Solutions:**
- Cek database connection
- Cek server logs untuk error details
- Pastikan DB_CONN env variable sudah dikonfigurasi dengan benar

---

## Testing

Anda dapat menguji endpoint-endpoint ini menggunakan tools berikut:

### Dengan Postman
1. Buka Postman
2. Buat request baru dengan method dan URL yang sesuai
3. Tambahkan request body jika diperlukan
4. Klik Send

### Dengan curl
Lihat contoh-contoh curl di setiap endpoint di atas.

### Dengan Thunder Client (VS Code)
1. Install Thunder Client extension
2. Buat request baru
3. Pilih method dan masukkan URL
4. Tambahkan body jika diperlukan
5. Klik Send

---

## Rate Limiting

Saat ini API tidak memiliki rate limiting. Penggunaan API unlimited untuk keperluan development dan testing.

---

## Versioning

Dokumentasi ini berlaku untuk:
- Kasir API v1.0
- Go 1.25.6
- PostgreSQL database

---

Last Updated: January 2026
