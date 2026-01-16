# 📚 Satpen API - Complete API Documentation

## 📋 Table of Contents

1. [API Overview](#api-overview)
2. [Base URL](#base-url)
3. [Response Format](#response-format)
4. [Error Codes](#error-codes)
5. [Authentication](#authentication)
6. [Rate Limiting](#rate-limiting)
7. [Endpoints](#endpoints)
   - [Health Check](#health-check)
   - [Satuan Pendidikan](#satuan-pendidikan)
   - [Master Data - Provinsi](#master-data---provinsi)
   - [Master Data - Kabupaten](#master-data---kabupaten)
   - [Master Data - Pengurus Cabang](#master-data---pengurus-cabang)
   - [Statistics](#statistics)

---

## API Overview

**Satpen API** adalah REST API untuk mengelola data Satuan Pendidikan (Sekolah/Madrasah) LP Ma'arif NU di seluruh Indonesia.

**Version:** 1.1.0
**Last Updated:** 2025-01-16

### Features
- ✅ List satuan pendidikan dengan filtering & pagination
- ✅ Detail satuan pendidikan
- ✅ Master data (Provinsi, Kabupaten, Pengurus Cabang)
- ✅ Statistics & aggregation
- ✅ Search functionality
- ✅ Rate limiting
- ✅ CORS support

---

## Base URL

### Development
```
http://localhost:8080
```

### Production
```
https://api.lpmaarifnu.or.id
```

---

## Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    // Response data
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error information"
}
```

### Paginated Response
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "items": [...],
    "pagination": {
      "current_page": 1,
      "total_pages": 10,
      "total_items": 200,
      "items_per_page": 20,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

---

## Error Codes

| HTTP Status | Description |
|-------------|-------------|
| 200 | OK - Request successful |
| 400 | Bad Request - Invalid parameters |
| 404 | Not Found - Resource not found |
| 429 | Too Many Requests - Rate limit exceeded |
| 500 | Internal Server Error - Server error |

---

## Authentication

**Current Version:** No authentication required (read-only API)

**Future Version:** JWT Bearer token authentication will be implemented.

---

## Rate Limiting

- **Default:** 60 requests per minute per IP
- **Header:** Check `X-RateLimit-Remaining` header
- **When exceeded:** Returns 429 status code

---

## Endpoints

### Health Check

#### Check API Health
```http
GET /health
```

**Description:** Check API and database connection status

**Response (200 OK):**
```json
{
  "success": true,
  "message": "API is healthy",
  "data": {
    "status": "ok",
    "database": "connected",
    "version": "1.1.0",
    "timestamp": "2025-01-16T10:30:00Z"
  }
}
```

**Example:**
```bash
curl http://localhost:8080/health
```

---

### Satuan Pendidikan

#### 1. Get All Satuan Pendidikan

```http
GET /api/v1/satpen
```

**Description:** Mendapatkan daftar satuan pendidikan dengan filtering, pagination, dan sorting

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| page | integer | No | 1 | Page number |
| limit | integer | No | 20 | Items per page (max: 100) |
| jenjang | string | No | - | Filter by jenjang pendidikan |
| provinsi | string | No | - | Filter by provinsi name |
| kabupaten | string | No | - | Filter by kabupaten name |
| search | string | No | - | Search by nama or alamat |
| akreditasi | string | No | - | Filter by akreditasi (A, B, C, D) |
| status | string | No | aktif | Filter by status |
| verified | boolean | No | - | Filter by verified status |
| sort | string | No | -created_at | Sort field (prefix - for desc) |

**Available Jenjang:**
- PAUD
- RA
- MI
- MTs
- MA
- SMK
- Pesantren

**Available Sort Fields:**
- nama
- tahun_berdiri
- jumlah_siswa
- jumlah_guru
- created_at (default: descending)

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Satuan pendidikan retrieved successfully",
  "data": {
    "satpen": [
      {
        "id": 1,
        "npsn": "20102701",
        "nama": "MI Ma'arif NU 01 Jakarta",
        "jenjang": {
          "id": 1,
          "nama": "MI",
          "lembaga": "MADRASAH"
        },
        "alamat": "Jl. Kramat Raya No. 123",
        "kelurahan": "Kramat",
        "kecamatan": "Senen",
        "kabupaten": {
          "id": 1,
          "nama": "Jakarta Pusat"
        },
        "provinsi": {
          "id": 1,
          "kode": "31",
          "nama": "DKI Jakarta"
        },
        "pengurus_cabang": {
          "id": 1,
          "nama": "LP Ma'arif NU Jakarta Pusat"
        },
        "yayasan": "Yayasan Pendidikan Ma'arif NU Jakarta",
        "kepala_sekolah": "Dr. Ahmad Fauzi, S.Pd.I",
        "phone": "021-3920123",
        "email": "mi01jakarta@maarifnu.or.id",
        "tahun_berdiri": 1985,
        "jumlah_siswa": 450,
        "jumlah_guru": 28,
        "akreditasi": "A",
        "status": "setujui",
        "is_verified": true,
        "verified_at": "2024-01-15T10:00:00Z",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-15T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 700,
      "total_items": 14000,
      "items_per_page": 20,
      "has_next": true,
      "has_prev": false
    },
    "statistics": {
      "total_satpen": 14000,
      "total_provinsi": 34,
      "by_jenjang": {
        "MI": 5600,
        "MTs": 4200,
        "MA": 2800
      }
    }
  }
}
```

**Examples:**
```bash
# Get all satpen (page 1, limit 20)
curl "http://localhost:8080/api/v1/satpen"

# Filter by jenjang MI
curl "http://localhost:8080/api/v1/satpen?jenjang=MI&page=1&limit=10"

# Filter by provinsi
curl "http://localhost:8080/api/v1/satpen?provinsi=Jawa%20Barat"

# Search by name
curl "http://localhost:8080/api/v1/satpen?search=Al-Maarif"

# Filter by akreditasi A
curl "http://localhost:8080/api/v1/satpen?akreditasi=A"

# Sort by jumlah siswa (descending)
curl "http://localhost:8080/api/v1/satpen?sort=-jumlah_siswa"

# Multiple filters
curl "http://localhost:8080/api/v1/satpen?jenjang=MI&provinsi=Jawa%20Barat&akreditasi=A&page=1"
```

---

#### 2. Get Satuan Pendidikan by ID/NPSN

```http
GET /api/v1/satpen/:id
```

**Description:** Mendapatkan detail satuan pendidikan berdasarkan ID atau NPSN

**Path Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | string | Yes | Satpen ID (numeric) or NPSN |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Satuan pendidikan retrieved successfully",
  "data": {
    "id": 1,
    "npsn": "20102701",
    "nama": "MI Ma'arif NU 01 Jakarta",
    "jenjang": {
      "id": 1,
      "nama": "MI",
      "lembaga": "MADRASAH"
    },
    "alamat": "Jl. Kramat Raya No. 123",
    "kelurahan": "Kramat",
    "kecamatan": "Senen",
    "kabupaten": {
      "id": 1,
      "nama": "Jakarta Pusat"
    },
    "provinsi": {
      "id": 1,
      "kode": "31",
      "nama": "DKI Jakarta"
    },
    "pengurus_cabang": {
      "id": 1,
      "nama": "LP Ma'arif NU Jakarta Pusat"
    },
    "yayasan": "Yayasan Pendidikan Ma'arif NU Jakarta",
    "kepala_sekolah": "Dr. Ahmad Fauzi, S.Pd.I",
    "phone": "021-3920123",
    "fax": "021-3920124",
    "email": "mi01jakarta@maarifnu.or.id",
    "tahun_berdiri": 1985,
    "jumlah_siswa": 450,
    "jumlah_guru": 28,
    "akreditasi": "A",
    "status": "setujui",
    "is_verified": true,
    "verified_at": "2024-01-15T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "success": false,
  "message": "Satuan pendidikan not found",
  "error": "record not found"
}
```

**Examples:**
```bash
# Get by ID
curl "http://localhost:8080/api/v1/satpen/1"

# Get by NPSN
curl "http://localhost:8080/api/v1/satpen/20102701"
```

---

#### 3. Get Satpen Statistics

```http
GET /api/v1/satpen/statistics
```

**Description:** Mendapatkan statistik ringkasan satuan pendidikan

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| provinsi | string | No | Filter by provinsi name |
| jenjang | string | No | Filter by jenjang |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Statistics retrieved successfully",
  "data": {
    "total_satpen": 14000,
    "total_provinsi": 34,
    "total_siswa": 2500000,
    "total_guru": 125000,
    "by_jenjang": {
      "PAUD": {
        "count": 1200,
        "siswa": 36000,
        "guru": 6000
      },
      "MI": {
        "count": 5600,
        "siswa": 1120000,
        "guru": 56000
      },
      "MTs": {
        "count": 4200,
        "siswa": 840000,
        "guru": 42000
      },
      "MA": {
        "count": 2800,
        "siswa": 448000,
        "guru": 19600
      }
    },
    "by_akreditasi": {
      "A": 6500,
      "B": 5200,
      "C": 1800,
      "D": 500
    },
    "top_provinsi": [
      {
        "provinsi": "Jawa Timur",
        "count": 2800
      },
      {
        "provinsi": "Jawa Tengah",
        "count": 2500
      },
      {
        "provinsi": "Jawa Barat",
        "count": 2100
      }
    ]
  }
}
```

**Examples:**
```bash
# Get all statistics
curl "http://localhost:8080/api/v1/satpen/statistics"

# Get statistics by provinsi
curl "http://localhost:8080/api/v1/satpen/statistics?provinsi=Jawa%20Barat"

# Get statistics by jenjang
curl "http://localhost:8080/api/v1/satpen/statistics?jenjang=MI"
```

---

### Master Data - Provinsi

#### 1. Get All Provinsi

```http
GET /api/v1/provinsi
```

**Description:** Mendapatkan daftar semua provinsi

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| search | string | No | Search by nama provinsi |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Provinsi retrieved successfully",
  "data": [
    {
      "id": 1,
      "map": "31",
      "kode": "31",
      "nama": "DKI Jakarta",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "kode": "32",
      "nama": "Jawa Barat",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Examples:**
```bash
# Get all provinsi
curl "http://localhost:8080/api/v1/provinsi"

# Search provinsi
curl "http://localhost:8080/api/v1/provinsi?search=Jawa"
```

---

#### 2. Get Provinsi by ID

```http
GET /api/v1/provinsi/:id
```

**Description:** Mendapatkan detail provinsi berdasarkan ID

**Path Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | Provinsi ID |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Provinsi retrieved successfully",
  "data": {
    "id": 1,
    "map": "31",
    "kode": "31",
    "nama": "DKI Jakarta",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "success": false,
  "message": "Provinsi not found",
  "error": "record not found"
}
```

**Examples:**
```bash
# Get provinsi by ID
curl "http://localhost:8080/api/v1/provinsi/1"
```

---

### Master Data - Kabupaten

#### 1. Get All Kabupaten

```http
GET /api/v1/kabupaten
```

**Description:** Mendapatkan daftar semua kabupaten/kota

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| provinsi_id | integer | No | Filter by provinsi ID |
| search | string | No | Search by nama kabupaten |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Kabupaten retrieved successfully",
  "data": [
    {
      "id": 1,
      "id_prov": 1,
      "provinsi": {
        "id": 1,
        "kode": "31",
        "nama": "DKI Jakarta"
      },
      "nama": "Jakarta Pusat",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Examples:**
```bash
# Get all kabupaten
curl "http://localhost:8080/api/v1/kabupaten"

# Filter by provinsi
curl "http://localhost:8080/api/v1/kabupaten?provinsi_id=1"

# Search kabupaten
curl "http://localhost:8080/api/v1/kabupaten?search=Bandung"

# Filter and search
curl "http://localhost:8080/api/v1/kabupaten?provinsi_id=32&search=Bandung"
```

---

#### 2. Get Kabupaten by ID

```http
GET /api/v1/kabupaten/:id
```

**Description:** Mendapatkan detail kabupaten berdasarkan ID

**Path Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | Kabupaten ID |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Kabupaten retrieved successfully",
  "data": {
    "id": 1,
    "id_prov": 1,
    "provinsi": {
      "id": 1,
      "kode": "31",
      "nama": "DKI Jakarta"
    },
    "nama": "Jakarta Pusat",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "success": false,
  "message": "Kabupaten not found",
  "error": "record not found"
}
```

**Examples:**
```bash
# Get kabupaten by ID
curl "http://localhost:8080/api/v1/kabupaten/1"
```

---

### Master Data - Pengurus Cabang

#### 1. Get All Pengurus Cabang

```http
GET /api/v1/pengurus-cabang
```

**Description:** Mendapatkan daftar pengurus cabang LP Ma'arif NU dengan pagination

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| page | integer | No | 1 | Page number |
| limit | integer | No | 20 | Items per page (max: 100) |
| provinsi_id | integer | No | - | Filter by provinsi ID |
| search | string | No | - | Search by nama cabang |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Pengurus cabang retrieved successfully",
  "data": {
    "pengurus_cabang": [
      {
        "id": 1,
        "id_prov": 1,
        "provinsi": {
          "id": 1,
          "kode": "31",
          "nama": "DKI Jakarta"
        },
        "kode_kab": "3171",
        "nama": "LP Ma'arif NU Jakarta Pusat",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 25,
      "total_items": 500,
      "items_per_page": 20,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

**Examples:**
```bash
# Get all pengurus cabang
curl "http://localhost:8080/api/v1/pengurus-cabang"

# With pagination
curl "http://localhost:8080/api/v1/pengurus-cabang?page=1&limit=10"

# Filter by provinsi
curl "http://localhost:8080/api/v1/pengurus-cabang?provinsi_id=32"

# Search
curl "http://localhost:8080/api/v1/pengurus-cabang?search=Jakarta"

# Filter and search
curl "http://localhost:8080/api/v1/pengurus-cabang?provinsi_id=32&search=Bandung&page=1&limit=10"
```

---

#### 2. Get Pengurus Cabang by ID

```http
GET /api/v1/pengurus-cabang/:id
```

**Description:** Mendapatkan detail pengurus cabang berdasarkan ID

**Path Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | Pengurus Cabang ID |

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Pengurus cabang retrieved successfully",
  "data": {
    "id": 1,
    "id_prov": 1,
    "provinsi": {
      "id": 1,
      "kode": "31",
      "nama": "DKI Jakarta"
    },
    "kode_kab": "3171",
    "nama": "LP Ma'arif NU Jakarta Pusat",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "success": false,
  "message": "Pengurus cabang not found",
  "error": "record not found"
}
```

**Examples:**
```bash
# Get pengurus cabang by ID
curl "http://localhost:8080/api/v1/pengurus-cabang/1"
```

---

## Quick Reference

### All Endpoints Summary

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/satpen` | List all satpen |
| GET | `/api/v1/satpen/:id` | Get satpen by ID/NPSN |
| GET | `/api/v1/satpen/statistics` | Get satpen statistics |
| GET | `/api/v1/provinsi` | List all provinsi |
| GET | `/api/v1/provinsi/:id` | Get provinsi by ID |
| GET | `/api/v1/kabupaten` | List all kabupaten |
| GET | `/api/v1/kabupaten/:id` | Get kabupaten by ID |
| GET | `/api/v1/pengurus-cabang` | List all pengurus cabang |
| GET | `/api/v1/pengurus-cabang/:id` | Get pengurus cabang by ID |

---

## Testing

### Using cURL

```bash
# Test all endpoints
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/satpen
curl http://localhost:8080/api/v1/satpen/1
curl http://localhost:8080/api/v1/satpen/statistics
curl http://localhost:8080/api/v1/provinsi
curl http://localhost:8080/api/v1/kabupaten
curl http://localhost:8080/api/v1/pengurus-cabang
```

### Using Postman

1. Import collection: `Satpen-API.postman_collection.json`
2. Set base_url variable to `http://localhost:8080`
3. Run requests from folders:
   - Health Check
   - Satuan Pendidikan
   - Master Data - Provinsi
   - Master Data - Kabupaten
   - Master Data - Pengurus Cabang
   - Statistics

---

## Support & Contact

### Documentation
- **API Docs:** [API_DOCUMENTATION.md](API_DOCUMENTATION.md)
- **Master Data Docs:** [API_MASTER_DATA.md](API_MASTER_DATA.md)
- **Postman Guide:** [POSTMAN_GUIDE.md](POSTMAN_GUIDE.md)
- **Implementation:** [IMPLEMENTATION.md](IMPLEMENTATION.md)
- **Changelog:** [CHANGELOG.md](CHANGELOG.md)

### Repository
- **GitHub:** https://github.com/lpmaarifnu/satpen-api

### Contact
- **Email:** dev@lpmaarifnu.or.id
- **Website:** https://lpmaarifnu.or.id

---

**Version:** 1.1.0
**Last Updated:** 2025-01-16
**License:** MIT
