# Satpen API - Implementation Guide

## ✅ Status Implementasi

API Satuan Pendidikan telah berhasil dibuat dan disesuaikan dengan database yang sudah ada (`testing_lpmaarif1`).

## 📊 Yang Sudah Dikerjakan

### 1. ✅ Models (Entity Layer)
Semua model telah disesuaikan dengan schema database yang ada:

- **[Satpen](internal/models/satuan_pendidikan.go)** - Tabel `satpen` dengan relasi lengkap
- **[Provinsi](internal/models/provinsi.go)** - Tabel `provinsi`
- **[Kabupaten](internal/models/kabupaten.go)** - Tabel `kabupaten`
- **[JenjangPendidikan](internal/models/jenjang_pendidikan.go)** - Tabel `jenjang_pendidikan`
- **[KategoriSatpen](internal/models/kategori_satpen.go)** - Tabel `kategori_satpen`
- **[PengurusCabang](internal/models/pengurus_cabang.go)** - Tabel `pengurus_cabang`
- **[PDPTK](internal/models/pdptk.go)** - Tabel `pdptk` untuk data siswa dan guru
- **[Statistics](internal/models/statistics.go)** - Model untuk statistics response

### 2. ✅ Repository Layer
File: [internal/repository/satpen_repository.go](internal/repository/satpen_repository.go)

Methods:
- `FindAll()` - Get list satpen dengan filtering, pagination, dan sorting
- `FindByID()` - Get satpen by ID
- `FindByNPSN()` - Get satpen by NPSN
- `GetStatistics()` - Get statistik lengkap
- `CountByJenjang()` - Count per jenjang
- `CountByAkreditasi()` - Count per akreditasi/kategori
- `GetTopProvinsi()` - Top provinsi berdasarkan jumlah satpen

**Features:**
- ✅ Preload semua relasi (Provinsi, Kabupaten, Jenjang, Kategori, PC, PDPTK)
- ✅ Filter by jenjang, provinsi, kabupaten, akreditasi
- ✅ Search by nama dan alamat
- ✅ Status mapping (setujui=aktif, lainnya=non-aktif)
- ✅ Verified mapping (setujui=verified)
- ✅ Sorting dengan field mapping
- ✅ Get latest PDPTK data per satpen

### 3. ✅ Service Layer
File: [internal/service/satpen_service.go](internal/service/satpen_service.go)

Methods:
- `GetAllSatpen()` - Business logic untuk list satpen
- `GetSatpenByID()` - Support ID dan NPSN
- `GetStatistics()` - Aggregate statistics

**Features:**
- ✅ Pagination validation
- ✅ Default values dari config
- ✅ ID atau NPSN lookup
- ✅ Error handling

### 4. ✅ Handler Layer
Files:
- [internal/handler/satpen_handler.go](internal/handler/satpen_handler.go)
- [internal/handler/health_handler.go](internal/handler/health_handler.go)

Endpoints:
- `GET /api/v1/satpen` - List satpen
- `GET /api/v1/satpen/:id` - Single satpen
- `GET /api/v1/satpen/statistics` - Statistics
- `GET /health` - Health check

### 5. ✅ Middleware
Files:
- [internal/middleware/cors.go](internal/middleware/cors.go) - CORS handling
- [internal/middleware/rate_limit.go](internal/middleware/rate_limit.go) - Rate limiting
- [internal/middleware/logger.go](internal/middleware/logger.go) - Request logging

### 6. ✅ Routes
File: [internal/routes/routes.go](internal/routes/routes.go)

Setup lengkap dengan:
- ✅ CORS middleware
- ✅ Logger middleware
- ✅ Rate limiter
- ✅ Route grouping `/api/v1`

### 7. ✅ Configuration
Files:
- [config.yaml](config.yaml) - Main configuration
- [internal/config/config.go](internal/config/config.go) - Config loader
- [.env.example](.env.example) - Environment variables template

**Config Features:**
- ✅ YAML-based configuration
- ✅ Environment variable override
- ✅ Database settings
- ✅ Rate limiting settings
- ✅ Pagination defaults
- ✅ CORS settings
- ✅ Logging configuration

### 8. ✅ Database
File: [internal/database/database.go](internal/database/database.go)

**Features:**
- ✅ Connection pool configuration
- ✅ Auto migration DISABLED (menggunakan DB existing)
- ✅ Health check
- ✅ Development/Production mode

### 9. ✅ Utils
File: [internal/utils/response.go](internal/utils/response.go)

Helper functions:
- ✅ SuccessResponse
- ✅ ErrorResponse
- ✅ ValidationErrorResponse
- ✅ NotFoundResponse
- ✅ InternalErrorResponse

### 10. ✅ Main Application
File: [cmd/api/main.go](cmd/api/main.go)

**Features:**
- ✅ Config loading
- ✅ Logger setup
- ✅ Database connection
- ✅ Dependency injection
- ✅ Route setup
- ✅ Server start

### 11. ✅ Documentation
Files:
- [README.md](README.md) - User guide
- [IMPLEMENTATION.md](IMPLEMENTATION.md) - This file
- [Makefile](Makefile) - Build automation

## 🔄 Field Mapping (Database → API)

### Satpen Model
```
Database Column         → API Field            → Description
================================================================================
id_satpen               → id                   → Primary key
npsn                    → npsn                 → Nomor Pokok Sekolah Nasional
nm_satpen               → nama                 → Nama satuan pendidikan
thn_berdiri             → tahun_berdiri        → Tahun berdiri
kepsek                  → kepala_sekolah       → Nama kepala sekolah
telpon                  → phone                → Nomor telepon
alamat                  → alamat               → Alamat lengkap
kecamatan               → kecamatan            → Kecamatan
kelurahan               → kelurahan            → Kelurahan
email                   → email                → Email satpen
yayasan                 → yayasan              → Nama yayasan
status                  → status               → Status (permohonan, setujui, dll)
created_at              → created_at           → Tanggal dibuat
updated_at              → updated_at           → Tanggal diupdate

# Virtual/Computed Fields (dari relasi)
pdptk.jml_pd            → jumlah_siswa         → Jumlah siswa dari PDPTK
pdptk.jml_guru          → jumlah_guru          → Jumlah guru dari PDPTK
kategori.nm_kategori    → akreditasi           → A/B/C/D dari kategori
status='setujui'        → is_verified=true     → Verified status
actived_date            → verified_at          → Tanggal verifikasi
```

### Status Mapping
```
Database Status         → API Status           → is_verified
================================================================================
setujui                 → aktif                → true
expired                 → non-aktif            → false
revisi                  → non-aktif            → false
permohonan              → non-aktif            → false
proses dokumen          → non-aktif            → false
perpanjangan            → aktif                → false
```

## 📡 API Response Format

### Success Response
```json
{
  "success": true,
  "message": "Satuan pendidikan retrieved successfully",
  "data": {
    "satpen": [...],
    "pagination": {
      "current_page": 1,
      "total_pages": 100,
      "total_items": 2000,
      "items_per_page": 20,
      "has_next": true,
      "has_prev": false
    },
    "statistics": {
      "total_satpen": 14000,
      "total_provinsi": 34,
      "total_siswa": 2500000,
      "total_guru": 125000,
      "by_jenjang": {...},
      "by_akreditasi": {...},
      "top_provinsi": [...]
    }
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error"
}
```

## 🚀 Quick Start

### 1. Install Dependencies
```bash
go mod download
go mod tidy
```

### 2. Configure Database
Edit `config.yaml`:
```yaml
database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "your_password"
  database: "testing_lpmaarif1"
```

### 3. Run
```bash
go run cmd/api/main.go
```

### 4. Test
```bash
# Health check
curl http://localhost:8080/health

# Get satpen
curl "http://localhost:8080/api/v1/satpen?page=1&limit=10"

# Get satpen by ID
curl http://localhost:8080/api/v1/satpen/1

# Get statistics
curl http://localhost:8080/api/v1/satpen/statistics
```

## 🔍 Query Examples

### Filter by Jenjang
```bash
curl "http://localhost:8080/api/v1/satpen?jenjang=MI&page=1&limit=20"
```

### Filter by Provinsi
```bash
curl "http://localhost:8080/api/v1/satpen?provinsi=DKI Jakarta&page=1"
```

### Search by Name
```bash
curl "http://localhost:8080/api/v1/satpen?search=Ma'arif&page=1"
```

### Filter by Akreditasi
```bash
curl "http://localhost:8080/api/v1/satpen?akreditasi=A&page=1"
```

### Sorting
```bash
# Ascending by nama
curl "http://localhost:8080/api/v1/satpen?sort=nama"

# Descending by jumlah siswa
curl "http://localhost:8080/api/v1/satpen?sort=-jumlah_siswa"
```

### Multiple Filters
```bash
curl "http://localhost:8080/api/v1/satpen?jenjang=MI&provinsi=DKI Jakarta&akreditasi=A&status=aktif&page=1&limit=20&sort=-jumlah_siswa"
```

## 🗂️ Database Tables Used

### Primary Tables
- `satpen` - Main satuan pendidikan data
- `pdptk` - Peserta didik & pendidik/tendik (siswa & guru)
- `provinsi` - Master provinsi
- `kabupaten` - Master kabupaten
- `jenjang_pendidikan` - Master jenjang (MI, MTs, MA, PAUD, dll)
- `kategori_satpen` - Master kategori/akreditasi (A, B, C, D)
- `pengurus_cabang` - Master pengurus cabang

### Relationships
```
satpen
├── provinsi (id_prov)
├── kabupaten (id_kab)
├── jenjang_pendidikan (id_jenjang)
├── kategori_satpen (id_kategori)
├── pengurus_cabang (id_pc)
└── pdptk (id_satpen) - Latest by tapel
```

## ⚙️ Configuration Options

### App Settings
```yaml
app:
  port: 8080                    # Server port
  env: "development"            # development|staging|production
  timezone: "Asia/Jakarta"      # Timezone
```

### Database Settings
```yaml
database:
  max_idle_conns: 10           # Max idle connections
  max_open_conns: 100          # Max open connections
  conn_max_lifetime: 3600      # Connection lifetime (seconds)
```

### Rate Limiting
```yaml
rate_limit:
  enabled: true
  satpen:
    requests: 60               # Max requests
    window: 60                 # Per window (seconds)
```

### Pagination
```yaml
pagination:
  default_page: 1
  default_limit: 20
  max_limit: 100
```

## 📦 Dependencies

```go
require (
    github.com/gin-gonic/gin v1.10.0
    github.com/sirupsen/logrus v1.9.3
    gopkg.in/yaml.v3 v3.0.1
    gorm.io/driver/mysql v1.5.7
    gorm.io/gorm v1.25.12
)
```

## 🎯 API Sesuai dengan TODO BACKEND

✅ **GET /api/v1/satpen** - Mendapatkan daftar satuan pendidikan dengan:
- Filtering (jenjang, provinsi, kabupaten, akreditasi, status, verified)
- Pagination (page, limit)
- Search (nama, alamat)
- Sorting (semua field)
- Response dengan statistics

✅ **GET /api/v1/satpen/:id** - Mendapatkan detail satpen by ID atau NPSN

✅ **GET /api/v1/satpen/statistics** - Statistik lengkap:
- Total satpen, provinsi, siswa, guru
- Breakdown by jenjang (dengan count, siswa, guru)
- Breakdown by akreditasi
- Top provinsi

✅ **Response Format** - Sesuai dengan spesifikasi di TODO BACKEND

## ✨ Features Tambahan

1. **Rate Limiting** - Protect API from abuse
2. **CORS** - Configurable CORS policy
3. **Logging** - Structured logging dengan Logrus
4. **Health Check** - `/health` endpoint
5. **Configurable** - YAML config + env variables
6. **Clean Architecture** - Repository → Service → Handler pattern

## 📝 Notes

- Database `testing_lpmaarif1` digunakan tanpa modifikasi
- Auto migration DISABLED (menggunakan schema yang sudah ada)
- PDPTK data diambil yang terbaru berdasarkan `tapel`
- Status mapping otomatis (setujui=aktif=verified)
- Field mapping otomatis via GORM tags
- Compatible dengan TODO BACKEND requirements

## 🎉 Ready to Use!

API sudah siap digunakan dan sesuai dengan:
- ✅ Database schema yang sudah ada
- ✅ TODO BACKEND requirements
- ✅ Best practices Go + Gin + GORM
- ✅ Clean Architecture
- ✅ Production ready features (rate limiting, CORS, logging)
