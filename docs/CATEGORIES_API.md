# API Dokumentasi - Categories Endpoints

Dokumentasi lengkap untuk semua endpoint yang berkaitan dengan Category (Kategori).

## Base URL
```
http://localhost:8080
```

## Endpoints Overview

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/categories` | Dapatkan semua kategori |
| POST | `/api/categories` | Buat kategori baru |
| GET | `/api/categories/{id}` | Dapatkan kategori berdasarkan ID |
| PUT | `/api/categories/{id}` | Update kategori |
| DELETE | `/api/categories/{id}` | Hapus kategori |

---

## ⚠️ Important Note

Category service saat ini bersifat **statis dan tidak terhubung langsung ke database/repository**. Data yang disimpan hanya tersimpan di memory selama aplikasi berjalan. Ini merupakan learning project untuk memahami architecture pattern.

---

## 1. Get All Categories

Mengambil daftar semua kategori.

**Request:**
```
GET /api/categories
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
    "name": "Minuman",
    "description": "Kategori minuman"
  },
  {
    "id": 2,
    "name": "Makanan",
    "description": "Kategori makanan"
  }
]
```

**Status Codes:**
- `200 OK` - Request berhasil

**Example dengan curl:**
```bash
curl -X GET http://localhost:8080/api/categories \
  -H "Content-Type: application/json"
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/categories')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

---

## 2. Create Category

Membuat kategori baru.

**Request:**
```
POST /api/categories
```

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

**Parameters:**

| Field | Type | Required | Deskripsi |
|-------|------|----------|-----------|
| name | string | Ya | Nama kategori |
| description | string | Ya | Deskripsi kategori |

**Response (201 Created):**
```json
{
  "id": 0,
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

**Status Codes:**
- `201 Created` - Kategori berhasil dibuat
- `400 Bad Request` - Invalid request format

**Example dengan curl:**
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minuman",
    "description": "Kategori minuman"
  }'
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/categories', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    name: 'Minuman',
    description: 'Kategori minuman'
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## 3. Get Category by ID

Mengambil detail kategori berdasarkan ID.

**Request:**
```
GET /api/categories/{id}
```

**URL Parameters:**

| Parameter | Type | Required | Deskripsi |
|-----------|------|----------|-----------|
| id | integer | Ya | ID kategori |

**Headers:**
```
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

**Status Codes:**
- `200 OK` - Request berhasil
- `400 Bad Request` - ID tidak valid (bukan integer)
- `404 Not Found` - Kategori tidak ditemukan

**Error Response (400):**
```
Invalid Category ID
```

**Error Response (404):**
```
Category not found
```

**Example dengan curl:**
```bash
curl -X GET http://localhost:8080/api/categories/1 \
  -H "Content-Type: application/json"
```

---

## 4. Update Category

Mengupdate data kategori yang sudah ada.

**Request:**
```
PUT /api/categories/{id}
```

**URL Parameters:**

| Parameter | Type | Required | Deskripsi |
|-----------|------|----------|-----------|
| id | integer | Ya | ID kategori yang akan diupdate |

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Minuman Premium",
  "description": "Kategori minuman premium"
}
```

**Parameters:**

| Field | Type | Required | Deskripsi |
|-------|------|----------|-----------|
| name | string | Ya | Nama kategori |
| description | string | Ya | Deskripsi kategori |

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Minuman Premium",
  "description": "Kategori minuman premium"
}
```

**Status Codes:**
- `200 OK` - Kategori berhasil diupdate
- `400 Bad Request` - ID tidak valid atau invalid request body

**Error Response (400):**
```
Invalid Category ID
```

**Example dengan curl:**
```bash
curl -X PUT http://localhost:8080/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minuman Premium",
    "description": "Kategori minuman premium"
  }'
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/categories/1', {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    name: 'Minuman Premium',
    description: 'Kategori minuman premium'
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## 5. Delete Category

Menghapus kategori.

**Request:**
```
DELETE /api/categories/{id}
```

**URL Parameters:**

| Parameter | Type | Required | Deskripsi |
|-----------|------|----------|-----------|
| id | integer | Ya | ID kategori yang akan dihapus |

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
- `200 OK` - Kategori berhasil dihapus
- `400 Bad Request` - ID tidak valid
- `404 Not Found` - Kategori tidak ditemukan

**Error Response (400):**
```
Invalid Category ID
```

**Error Response (404):**
```
Category not found
```

**Example dengan curl:**
```bash
curl -X DELETE http://localhost:8080/api/categories/1 \
  -H "Content-Type: application/json"
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/categories/1', {
  method: 'DELETE',
  headers: {
    'Content-Type': 'application/json'
  }
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## Data Model

### Category Object

```json
{
  "id": 1,
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

**Field Descriptions:**

| Field | Type | Deskripsi |
|-------|------|-----------|
| id | integer | ID unik kategori (auto-generated) |
| name | string | Nama kategori |
| description | string | Deskripsi atau penjelasan kategori |

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
Resource (kategori) tidak ditemukan
```

**Solutions:**
- Verifikasi ID kategori yang digunakan
- Pastikan kategori sudah ada di memory sebelum mengupdate/menghapus
- Ingat: Data kategori hanya tersimpan di memory selama aplikasi berjalan

### Important: Memory-based Storage

Karena Category service bersifat statis dan berbasis memory:
- **Data akan hilang** ketika aplikasi di-restart
- **Data hanya tersimpan di sesi aplikasi saat berjalan**
- Ini adalah desain intentional untuk learning purposes

---

## Example Workflow

### 1. Buat kategori baru
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minuman",
    "description": "Kategori minuman"
  }'
```

**Response:**
```json
{
  "id": 1,
  "name": "Minuman",
  "description": "Kategori minuman"
}
```

### 2. Ambil semua kategori
```bash
curl -X GET http://localhost:8080/api/categories
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "Minuman",
    "description": "Kategori minuman"
  }
]
```

### 3. Update kategori
```bash
curl -X PUT http://localhost:8080/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minuman Panas",
    "description": "Kategori minuman panas"
  }'
```

### 4. Hapus kategori
```bash
curl -X DELETE http://localhost:8080/api/categories/1
```

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

## Versioning

Dokumentasi ini berlaku untuk:
- Kasir API v1.0
- Go 1.25.6
- Category Service (Static/Memory-based)

---

Last Updated: January 2026
