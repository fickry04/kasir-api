# API Dokumentasi - Health Check Endpoint

Dokumentasi untuk Health Check endpoint yang digunakan untuk mengecek status API.

## Base URL
```
http://localhost:8080
```

---

## Health Check

Endpoint ini digunakan untuk memverifikasi bahwa API sudah berjalan dan siap menerima requests.

**Request:**
```
GET /health
```

**Parameters:** None

**Headers:**
```
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "status": "OK",
  "message": "Api Running"
}
```

**Status Codes:**
- `200 OK` - API sedang berjalan dengan baik

**Example dengan curl:**
```bash
curl -X GET http://localhost:8080/health
```

**Output:**
```json
{"status":"OK","message":"Api Running"}
```

**Example dengan JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/health')
  .then(response => response.json())
  .then(data => {
    console.log('Status:', data.status);
    console.log('Message:', data.message);
  })
  .catch(error => console.error('Error:', error));
```

**Example dengan axios:**
```javascript
const axios = require('axios');

axios.get('http://localhost:8080/health')
  .then(response => {
    console.log('Health Check Response:', response.data);
  })
  .catch(error => {
    console.error('Health Check Failed:', error.message);
  });
```

---

## Use Cases

### 1. Aplikasi Startup
Gunakan endpoint ini untuk mengecek apakah API sudah siap sebelum mengirimkan requests penting:

```javascript
async function waitForAPI(maxAttempts = 5) {
  for (let i = 0; i < maxAttempts; i++) {
    try {
      const response = await fetch('http://localhost:8080/health');
      if (response.ok) {
        console.log('API is ready!');
        return true;
      }
    } catch (error) {
      console.log(`Attempt ${i + 1}/${maxAttempts} failed, retrying...`);
      await new Promise(resolve => setTimeout(resolve, 1000));
    }
  }
  console.log('API is not available');
  return false;
}
```

### 2. Monitoring & Alerting
Setup monitoring untuk memastikan API tetap berjalan:

```javascript
setInterval(async () => {
  try {
    const response = await fetch('http://localhost:8080/health');
    if (!response.ok) {
      console.error('API down! Status:', response.status);
      // Send alert/notification
    }
  } catch (error) {
    console.error('API unreachable:', error.message);
    // Send alert/notification
  }
}, 30000); // Check setiap 30 detik
```

### 3. Load Balancer / Reverse Proxy
Health check dapat digunakan untuk:
- Menentukan apakah server siap menerima traffic
- Menghilangkan server yang down dari rotation
- Mendeteksi server yang perlu restart

---

## Response Format

### Success Response

```json
{
  "status": "OK",
  "message": "Api Running"
}
```

**Field Descriptions:**

| Field | Type | Deskripsi |
|-------|------|-----------|
| status | string | Status API (OK = running) |
| message | string | Pesan deskriptif |

---

## Expected Behavior

| Scenario | Status | Response |
|----------|--------|----------|
| API berjalan normal | 200 | `{"status":"OK","message":"Api Running"}` |
| API tidak berjalan | Connection refused | N/A |
| Server error | 500 | Error message |

---

## Integration Examples

### Node.js / Express
```javascript
const express = require('express');
const axios = require('axios');

const app = express();

// Health check sebelum memproses request
app.use(async (req, res, next) => {
  try {
    const healthResponse = await axios.get('http://localhost:8080/health');
    if (healthResponse.data.status === 'OK') {
      next();
    } else {
      res.status(503).json({ error: 'Service unavailable' });
    }
  } catch (error) {
    res.status(503).json({ error: 'Backend service unavailable' });
  }
});

app.listen(3000, () => console.log('Frontend running on port 3000'));
```

### Python
```python
import requests
import time
from datetime import datetime

def check_api_health(url='http://localhost:8080/health'):
    try:
        response = requests.get(url, timeout=5)
        if response.status_code == 200:
            data = response.json()
            print(f"[{datetime.now()}] API Status: {data['status']} - {data['message']}")
            return True
        else:
            print(f"[{datetime.now()}] API returned status {response.status_code}")
            return False
    except requests.exceptions.RequestException as e:
        print(f"[{datetime.now()}] Error checking API: {e}")
        return False

# Continuous monitoring
while True:
    check_api_health()
    time.sleep(30)  # Check every 30 seconds
```

### Go
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func checkHealth() {
	response, err := http.Get("http://localhost:8080/health")
	if err != nil {
		fmt.Println("Error checking health:", err)
		return
	}
	defer response.Body.Close()

	var health HealthResponse
	json.NewDecoder(response.Body).Decode(&health)
	fmt.Printf("API Status: %s - %s\n", health.Status, health.Message)
}

func main() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		checkHealth()
	}
}
```

---

## Best Practices

1. **Timeout**: Set timeout yang reasonable (3-5 detik) untuk health check
2. **Retry Logic**: Implementasikan retry dengan exponential backoff
3. **Frequency**: Jangan check terlalu sering (recommend 30-60 detik)
4. **Logging**: Log semua health check results untuk debugging
5. **Alerting**: Setup alerts jika API tidak respond

---

## Troubleshooting

### "Connection Refused"
```
Error: connect ECONNREFUSED 127.0.0.1:8080
```

**Solutions:**
- Pastikan server sudah dijalankan dengan `go run cmd/server/main.go`
- Check port 8080 apakah sudah digunakan aplikasi lain
- Gunakan `lsof -i :8080` (Unix/Mac) atau `netstat -ano | findstr 8080` (Windows)

### "Request Timeout"
```
Error: request to http://localhost:8080/health timed out after 5000ms
```

**Solutions:**
- Check apakah network connectivity baik
- API mungkin sedang overloaded, coba beberapa saat lagi
- Pastikan localhost:8080 accessible dari mesin tempat request dikirim

### "Connection Reset"
```
Error: socket hang up
```

**Solutions:**
- API mungkin crash, restart server
- Check server logs untuk error details

---

## Versioning

Dokumentasi ini berlaku untuk:
- Kasir API v1.0
- Go 1.25.6

---

Last Updated: January 2026
