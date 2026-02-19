# Satpen API - LP Ma'arif NU

REST API untuk Data Satuan Pendidikan Ma'arif NU di seluruh Indonesia.

## 🚀 Features

### Satuan Pendidikan
✅ **GET /api/v1/satpen** - Get all satuan pendidikan dengan filtering & pagination
✅ **GET /api/v1/satpen/:id** - Get satuan pendidikan berdasarkan ID atau NPSN
✅ **GET /api/v1/satpen/statistics** - Get statistik ringkasan

### Master Data
✅ **GET /api/v1/provinsi** - Get all provinsi
✅ **GET /api/v1/provinsi/:id** - Get provinsi by ID
✅ **GET /api/v1/kabupaten** - Get all kabupaten/kota
✅ **GET /api/v1/kabupaten/:id** - Get kabupaten by ID
✅ **GET /api/v1/pengurus-cabang** - Get all pengurus cabang
✅ **GET /api/v1/pengurus-cabang/:id** - Get pengurus cabang by ID
✅ **GET /api/v1/jenjang-pendidikan** - Get all jenjang pendidikan
✅ **GET /api/v1/jenjang-pendidikan/:id** - Get jenjang pendidikan by ID

### System
✅ **GET /health** - Health check endpoint

**Total:** 19 endpoints

## 🛠️ Tech Stack

- **Go 1.21+** - Programming language
- **Gin** - Web framework
- **GORM** - ORM untuk database operations
- **MySQL 8.0+** - Database (menggunakan database `testing_lpmaarif1`)
- **Logrus** - Structured logging
- **YAML** - Configuration management

## 📋 Prerequisites

- Go 1.21 atau lebih tinggi
- MySQL 8.0+ dengan database `testing_lpmaarif1`
- Git

## ⚙️ Installation

### 1. Install dependencies

```bash
go mod download
go mod tidy
```

### 2. Configure database

Edit `config.yaml`:
```yaml
database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "your_password"
  database: "testing_lpmaarif1"
```

Atau gunakan `.env`:
```bash
cp .env.example .env
```

Edit `.env`:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=your_password
DB_DATABASE=testing_lpmaarif1
```

### 3. Run application

```bash
go run cmd/api/main.go
```

API akan tersedia di `http://localhost:8080`

## 📖 API Documentation

### 📚 Complete Documentation

- **[API Quick Reference](API_QUICK_REFERENCE.md)** - Quick lookup for all endpoints
- **[API Documentation](API_DOCUMENTATION.md)** - Complete API documentation
- **[API Master Data](API_MASTER_DATA.md)** - Master data endpoints details
- **[Postman Guide](POSTMAN_GUIDE.md)** - Postman collection usage guide
- **[Implementation Details](IMPLEMENTATION.md)** - Technical implementation
- **[Changelog](CHANGELOG.md)** - Version history

### Quick Examples

#### 1. Get All Satuan Pendidikan

```http
GET /api/v1/satpen?page=1&limit=20&jenjang=MI&provinsi=DKI Jakarta
```

**Query Parameters:**
- `page` - Nomor halaman (default: 1)
- `limit` - Items per halaman (default: 20, max: 100)
- `jenjang` - Filter jenjang (MI, MTs, MA, PAUD, dll)
- `provinsi` - Filter provinsi
- `kabupaten` - Filter kabupaten
- `search` - Pencarian nama/alamat
- `akreditasi` - Filter akreditasi (A, B, C, D)
- `status` - Filter status (aktif/non-aktif, default: aktif)
- `verified` - Filter verifikasi (true/false)
- `sort` - Sorting (default: -created_at)

#### 2. Get Single Satpen

```http
GET /api/v1/satpen/:id
GET /api/v1/satpen/20102701
```

#### 3. Get Statistics

```http
GET /api/v1/satpen/statistics?provinsi=DKI Jakarta
```

#### 4. Get Provinsi

```http
GET /api/v1/provinsi
GET /api/v1/provinsi/1
```

#### 5. Get Kabupaten

```http
GET /api/v1/kabupaten?provinsi_id=32&search=Bandung
GET /api/v1/kabupaten/1
```

#### 6. Get Pengurus Cabang

```http
GET /api/v1/pengurus-cabang?provinsi_id=32&page=1&limit=10
GET /api/v1/pengurus-cabang/1
```

#### 7. Get Jenjang Pendidikan

```http
GET /api/v1/jenjang-pendidikan
GET /api/v1/jenjang-pendidikan/1
GET /api/v1/jenjang-pendidikan?search=MI
```

#### 8. Health Check

```http
GET /health
```

**📖 For complete API documentation, see [API_DOCUMENTATION.md](API_DOCUMENTATION.md)**

## 📁 Project Structure

```
satpen-api/
├── cmd/api/main.go              # Entry point
├── internal/
│   ├── config/                  # Configuration
│   ├── database/                # Database connection
│   ├── models/                  # Data models
│   ├── repository/              # Data access
│   ├── service/                 # Business logic
│   ├── handler/                 # HTTP handlers
│   ├── middleware/              # Middleware
│   ├── routes/                  # Routes
│   └── utils/                   # Utilities
├── config.yaml                   # Configuration
├── go.mod                        # Dependencies
└── README.md
```

## 🗄️ Database Schema

API menggunakan database `testing_lpmaarif1` dengan tabel:

- **satpen** - Data satuan pendidikan
- **provinsi** - Master provinsi
- **kabupaten** - Master kabupaten
- **jenjang_pendidikan** - Master jenjang
- **kategori_satpen** - Master kategori/akreditasi
- **pengurus_cabang** - Data pengurus cabang
- **pdptk** - Data siswa & guru

## 🔄 Field Mapping

Database → API Response:

- `nm_satpen` → `nama`
- `thn_berdiri` → `tahun_berdiri`
- `kepsek` → `kepala_sekolah`
- `telpon` → `phone`
- `jml_pd` (pdptk) → `jumlah_siswa`
- `jml_guru` (pdptk) → `jumlah_guru`
- `nm_kategori` → `akreditasi`
- `status='setujui'` → `is_verified=true`

## 🔧 Configuration

Edit `config.yaml` untuk konfigurasi:

```yaml
app:
  port: 8080
  env: "development"

database:
  database: "testing_lpmaarif1"

rate_limit:
  enabled: true
  satpen:
    requests: 60
    window: 60

pagination:
  default_limit: 20
  max_limit: 100
```

## 📝 Development

### Build

```bash
go build -o bin/satpen-api cmd/api/main.go
```

### Run

```bash
.\bin\satpen-api.exe
```

### Using Makefile

```bash
make install    # Install dependencies
make build      # Build binary
make run        # Run app
make test       # Run tests
```

## 📄 License

Copyright © 2025 LP Ma'arif NU
