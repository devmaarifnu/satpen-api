# TODO BACKEND - SATUAN PENDIDIKAN API

## 📋 Overview
Dokumen ini berisi **API Endpoints untuk Data Satuan Pendidikan (Satpen)**. Ini adalah fitur prioritas tinggi yang dibutuhkan untuk menampilkan database sekolah-sekolah Ma'arif di seluruh Indonesia.

---

## 🎯 API Endpoints - Satuan Pendidikan

### 1. Get All Satuan Pendidikan
```http
GET /api/v1/satpen
```

**Description:** Mendapatkan daftar satuan pendidikan Ma'arif dengan filtering dan pagination

**Query Parameters:**
- `page` (optional, default: 1): Page number
- `limit` (optional, default: 20): Items per page (max: 100)
- `jenjang` (optional): Filter by jenjang (PAUD, MI, MTs, MA, Pesantren, Perguruan Tinggi)
- `provinsi` (optional): Filter by provinsi name or ID
- `kabupaten` (optional): Filter by kabupaten/kota name or ID
- `search` (optional): Search by nama satpen or alamat (fulltext search)
- `akreditasi` (optional): Filter by akreditasi (A, B, C, Belum Terakreditasi)
- `status` (optional, default: aktif): Filter by status (aktif, non-aktif)
- `verified` (optional): Filter by verification status (true/false)
- `sort` (optional, default: -created_at): Sort field (nama, jumlah_siswa, tahun_berdiri, akreditasi)

**Example Requests:**
```bash
# Get all MI schools in DKI Jakarta
curl "https://api.lpmaarifnu.or.id/v1/satpen?jenjang=MI&provinsi=DKI Jakarta&page=1&limit=20"

# Search schools by name
curl "https://api.lpmaarifnu.or.id/v1/satpen?search=Al-Maarif&page=1"

# Get schools with A accreditation
curl "https://api.lpmaarifnu.or.id/v1/satpen?akreditasi=A&page=1"

# Get schools sorted by number of students
curl "https://api.lpmaarifnu.or.id/v1/satpen?sort=-jumlah_siswa&page=1"
```

**Response Success (200):**
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
        "jenjang": "MI",
        "alamat": "Jl. Kramat Raya No. 123",
        "kabupaten": {
          "id": 1,
          "nama": "Jakarta Pusat",
          "kode": "3171"
        },
        "provinsi": {
          "id": 1,
          "nama": "DKI Jakarta",
          "kode": "31"
        },
        "kode_pos": "10450",
        "kepala_sekolah": "Dr. Ahmad Fauzi, S.Pd.I",
        "email": "mi01jakarta@maarifnu.or.id",
        "phone": "021-3920123",
        "website": null,
        "jumlah_siswa": 450,
        "jumlah_guru": 28,
        "jumlah_rombel": 12,
        "akreditasi": "A",
        "tahun_berdiri": 1985,
        "coordinates": {
          "latitude": -6.1754,
          "longitude": 106.8272
        },
        "status": "aktif",
        "is_verified": true,
        "verified_at": "2024-01-15T10:00:00Z",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-15T10:00:00Z"
      },
      {
        "id": 2,
        "npsn": "20102702",
        "nama": "MTs Ma'arif NU 02 Jakarta",
        "jenjang": "MTs",
        "alamat": "Jl. Tanah Abang II No. 45",
        "kabupaten": {
          "id": 1,
          "nama": "Jakarta Pusat",
          "kode": "3171"
        },
        "provinsi": {
          "id": 1,
          "nama": "DKI Jakarta",
          "kode": "31"
        },
        "kode_pos": "10160",
        "kepala_sekolah": "Dra. Hj. Siti Nurjanah, M.Pd",
        "email": "mts02jakarta@maarifnu.or.id",
        "phone": "021-3841234",
        "website": null,
        "jumlah_siswa": 580,
        "jumlah_guru": 35,
        "jumlah_rombel": 18,
        "akreditasi": "A",
        "tahun_berdiri": 1990,
        "coordinates": {
          "latitude": -6.1844,
          "longitude": 106.8122
        },
        "status": "aktif",
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
      "total_siswa": 2500000,
      "by_jenjang": {
        "PAUD": 1200,
        "MI": 5600,
        "MTs": 4200,
        "MA": 2800,
        "Pesantren": 180,
        "Perguruan Tinggi": 20
      },
      "by_akreditasi": {
        "A": 6500,
        "B": 5200,
        "C": 1800,
        "Belum Terakreditasi": 500
      }
    }
  }
}
```

**Response Error (400):**
```json
{
  "success": false,
  "message": "Invalid query parameters",
  "errors": {
    "page": "must be a positive integer",
    "limit": "must not exceed 100"
  }
}
```

**Response Error (500):**
```json
{
  "success": false,
  "message": "Internal server error",
  "error": "Database query failed"
}
```

---

### 2. Get Single Satuan Pendidikan
```http
GET /api/v1/satpen/:id
```

**Path Parameters:**
- `id` (required): Satpen ID or NPSN

**Example Requests:**
```bash
# Get by ID
curl https://api.lpmaarifnu.or.id/v1/satpen/1

# Get by NPSN
curl https://api.lpmaarifnu.or.id/v1/satpen/20102701
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Satuan pendidikan retrieved successfully",
  "data": {
    "id": 1,
    "npsn": "20102701",
    "nama": "MI Ma'arif NU 01 Jakarta",
    "jenjang": "MI",
    "alamat": "Jl. Kramat Raya No. 123",
    "kabupaten": {
      "id": 1,
      "nama": "Jakarta Pusat",
      "kode": "3171"
    },
    "provinsi": {
      "id": 1,
      "nama": "DKI Jakarta",
      "kode": "31"
    },
    "kode_pos": "10450",
    "kepala_sekolah": "Dr. Ahmad Fauzi, S.Pd.I",
    "email": "mi01jakarta@maarifnu.or.id",
    "phone": "021-3920123",
    "website": null,
    "jumlah_siswa": 450,
    "jumlah_guru": 28,
    "jumlah_rombel": 12,
    "akreditasi": "A",
    "tahun_berdiri": 1985,
    "coordinates": {
      "latitude": -6.1754,
      "longitude": 106.8272
    },
    "status": "aktif",
    "is_verified": true,
    "verified_at": "2024-01-15T10:00:00Z",
    "verified_by": {
      "id": 1,
      "name": "Super Admin"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Response Error (404):**
```json
{
  "success": false,
  "message": "Satuan pendidikan not found",
  "error": "Satpen with ID '999' does not exist"
}
```

---

### 3. Get Statistics Summary
```http
GET /api/v1/satpen/statistics
```

**Description:** Mendapatkan statistik ringkasan satuan pendidikan

**Query Parameters:**
- `provinsi` (optional): Filter statistics by provinsi
- `jenjang` (optional): Filter statistics by jenjang

**Response Success (200):**
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
      },
      "Pesantren": {
        "count": 180,
        "siswa": 54000,
        "guru": 1350
      },
      "Perguruan Tinggi": {
        "count": 20,
        "siswa": 2000,
        "guru": 50
      }
    },
    "by_akreditasi": {
      "A": 6500,
      "B": 5200,
      "C": 1800,
      "Belum Terakreditasi": 500
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

---

## 🗂️ Database Schema

### Main Table: `satuan_pendidikan`
```sql
CREATE TABLE satuan_pendidikan (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    npsn VARCHAR(20) NOT NULL UNIQUE COMMENT 'Nomor Pokok Sekolah Nasional',
    nama VARCHAR(500) NOT NULL,
    jenjang ENUM('PAUD', 'MI', 'MTs', 'MA', 'Pesantren', 'Perguruan Tinggi') NOT NULL,
    alamat TEXT NOT NULL,
    kabupaten_id INT UNSIGNED,
    provinsi_id INT UNSIGNED NOT NULL,
    kode_pos VARCHAR(10),
    kepala_sekolah VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(20),
    website VARCHAR(500),
    jumlah_siswa INT UNSIGNED DEFAULT 0,
    jumlah_guru INT UNSIGNED DEFAULT 0,
    jumlah_rombel INT UNSIGNED DEFAULT 0 COMMENT 'Rombongan belajar',
    akreditasi ENUM('A', 'B', 'C', 'Belum Terakreditasi') DEFAULT 'Belum Terakreditasi',
    tahun_berdiri YEAR,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    status ENUM('aktif', 'non-aktif') DEFAULT 'aktif',
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP NULL,
    verified_by BIGINT UNSIGNED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- Foreign Keys
    FOREIGN KEY (kabupaten_id) REFERENCES kabupaten_kota(id) ON DELETE SET NULL,
    FOREIGN KEY (provinsi_id) REFERENCES provinsi(id) ON DELETE RESTRICT,
    FOREIGN KEY (verified_by) REFERENCES users(id) ON DELETE SET NULL,

    -- Indexes
    INDEX idx_npsn (npsn),
    INDEX idx_jenjang (jenjang),
    INDEX idx_provinsi_id (provinsi_id),
    INDEX idx_kabupaten_id (kabupaten_id),
    INDEX idx_akreditasi (akreditasi),
    INDEX idx_status (status),
    INDEX idx_coordinates (latitude, longitude),
    FULLTEXT INDEX idx_search (nama, alamat)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### Related Tables:

#### `provinsi`
```sql
CREATE TABLE provinsi (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    kode VARCHAR(5) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_kode (kode)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### `kabupaten_kota`
```sql
CREATE TABLE kabupaten_kota (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    provinsi_id INT UNSIGNED NOT NULL,
    nama VARCHAR(100) NOT NULL,
    kode VARCHAR(10) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (provinsi_id) REFERENCES provinsi(id) ON DELETE CASCADE,
    INDEX idx_provinsi_id (provinsi_id),
    INDEX idx_kode (kode)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### `pengurus_cabang`
```sql
CREATE TABLE pengurus_cabang (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nama_cabang VARCHAR(255) NOT NULL,
    kabupaten_id INT UNSIGNED NOT NULL,
    provinsi_id INT UNSIGNED NOT NULL,
    ketua VARCHAR(255) NOT NULL,
    sekretaris VARCHAR(255),
    bendahara VARCHAR(255),
    alamat TEXT,
    phone VARCHAR(20),
    email VARCHAR(255),
    masa_khidmat VARCHAR(20) COMMENT 'Format: 2023-2028',
    status ENUM('aktif', 'non-aktif') DEFAULT 'aktif',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- Foreign Keys
    FOREIGN KEY (kabupaten_id) REFERENCES kabupaten_kota(id) ON DELETE CASCADE,
    FOREIGN KEY (provinsi_id) REFERENCES provinsi(id) ON DELETE CASCADE,

    -- Indexes
    INDEX idx_kabupaten_id (kabupaten_id),
    INDEX idx_provinsi_id (provinsi_id),
    INDEX idx_status (status),
    FULLTEXT INDEX idx_search (nama_cabang, ketua)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

---

### 4. Get List Provinsi
```http
GET /api/v1/provinsi
```

**Description:** Mendapatkan daftar semua provinsi

**Query Parameters:**
- `search` (optional): Search by nama provinsi

**Example Requests:**
```bash
# Get all provinces
curl https://api.lpmaarifnu.or.id/v1/provinsi

# Search provinces
curl "https://api.lpmaarifnu.or.id/v1/provinsi?search=Jawa"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Provinsi retrieved successfully",
  "data": [
    {
      "id": 1,
      "nama": "DKI Jakarta",
      "kode": "31",
      "total_satpen": 450,
      "total_kabupaten": 6
    },
    {
      "id": 2,
      "nama": "Jawa Barat",
      "kode": "32",
      "total_satpen": 2100,
      "total_kabupaten": 27
    },
    {
      "id": 3,
      "nama": "Jawa Tengah",
      "kode": "33",
      "total_satpen": 2500,
      "total_kabupaten": 35
    }
  ]
}
```

---

### 5. Get List Kabupaten/Kota
```http
GET /api/v1/kabupaten
```

**Description:** Mendapatkan daftar kabupaten/kota

**Query Parameters:**
- `provinsi_id` (optional): Filter by provinsi ID
- `search` (optional): Search by nama kabupaten

**Example Requests:**
```bash
# Get all kabupaten
curl https://api.lpmaarifnu.or.id/v1/kabupaten

# Get kabupaten by provinsi
curl "https://api.lpmaarifnu.or.id/v1/kabupaten?provinsi_id=32"

# Search kabupaten
curl "https://api.lpmaarifnu.or.id/v1/kabupaten?search=Bandung"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Kabupaten retrieved successfully",
  "data": [
    {
      "id": 1,
      "nama": "Jakarta Pusat",
      "kode": "3171",
      "provinsi": {
        "id": 1,
        "nama": "DKI Jakarta",
        "kode": "31"
      },
      "total_satpen": 75
    },
    {
      "id": 2,
      "nama": "Jakarta Utara",
      "kode": "3172",
      "provinsi": {
        "id": 1,
        "nama": "DKI Jakarta",
        "kode": "31"
      },
      "total_satpen": 68
    }
  ]
}
```

---

### 6. Get List Pengurus Cabang
```http
GET /api/v1/pengurus-cabang
```

**Description:** Mendapatkan daftar pengurus cabang LP Ma'arif NU

**Query Parameters:**
- `provinsi_id` (optional): Filter by provinsi ID
- `kabupaten_id` (optional): Filter by kabupaten ID
- `search` (optional): Search by nama cabang atau ketua
- `page` (optional, default: 1): Page number
- `limit` (optional, default: 20): Items per page

**Example Requests:**
```bash
# Get all pengurus cabang
curl https://api.lpmaarifnu.or.id/v1/pengurus-cabang

# Get by provinsi
curl "https://api.lpmaarifnu.or.id/v1/pengurus-cabang?provinsi_id=32"

# Get by kabupaten
curl "https://api.lpmaarifnu.or.id/v1/pengurus-cabang?kabupaten_id=15"

# Search
curl "https://api.lpmaarifnu.or.id/v1/pengurus-cabang?search=Bandung"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Pengurus cabang retrieved successfully",
  "data": {
    "pengurus_cabang": [
      {
        "id": 1,
        "nama_cabang": "LP Ma'arif NU Kota Bandung",
        "kabupaten": {
          "id": 15,
          "nama": "Kota Bandung",
          "kode": "3273"
        },
        "provinsi": {
          "id": 2,
          "nama": "Jawa Barat",
          "kode": "32"
        },
        "ketua": "Dr. H. Ahmad Sanusi, M.Pd.I",
        "sekretaris": "Drs. H. Muhammad Yusuf",
        "bendahara": "Hj. Siti Maryam, S.Pd",
        "alamat": "Jl. Soekarno Hatta No. 456, Bandung",
        "phone": "022-7501234",
        "email": "lpmaarif.bandung@nu.or.id",
        "total_satpen": 156,
        "masa_khidmat": "2023-2028",
        "status": "aktif",
        "created_at": "2023-06-15T10:00:00Z",
        "updated_at": "2024-01-10T15:30:00Z"
      },
      {
        "id": 2,
        "nama_cabang": "LP Ma'arif NU Kabupaten Bandung",
        "kabupaten": {
          "id": 14,
          "nama": "Kabupaten Bandung",
          "kode": "3204"
        },
        "provinsi": {
          "id": 2,
          "nama": "Jawa Barat",
          "kode": "32"
        },
        "ketua": "K.H. Abdullah Syukri, Lc., M.A",
        "sekretaris": "Ust. Rifki Hamdani, S.Pd.I",
        "bendahara": "Hj. Nur Azizah, S.E",
        "alamat": "Jl. Raya Soreang No. 89, Soreang, Kab. Bandung",
        "phone": "022-5890123",
        "email": "lpmaarif.kabbandung@nu.or.id",
        "total_satpen": 243,
        "masa_khidmat": "2022-2027",
        "status": "aktif",
        "created_at": "2022-08-20T10:00:00Z",
        "updated_at": "2024-01-10T15:30:00Z"
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

**Response Error (400):**
```json
{
  "success": false,
  "message": "Invalid query parameters",
  "errors": {
    "provinsi_id": "must be a positive integer"
  }
}
```

---

## 🔧 Implementation Guide

### 1. Repository Layer (Golang + GORM)

```go
// internal/repositories/satpen_repository.go
package repositories

import (
    "gorm.io/gorm"
    "lpmaarifnu-api/internal/models"
)

type SatpenRepository interface {
    FindAll(filters map[string]interface{}, page, limit int) ([]models.SatuanPendidikan, int64, error)
    FindByID(id uint) (*models.SatuanPendidikan, error)
    FindByNPSN(npsn string) (*models.SatuanPendidikan, error)
    GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error)
}

type satpenRepository struct {
    db *gorm.DB
}

func NewSatpenRepository(db *gorm.DB) SatpenRepository {
    return &satpenRepository{db: db}
}

func (r *satpenRepository) FindAll(filters map[string]interface{}, page, limit int) ([]models.SatuanPendidikan, int64, error) {
    var satpen []models.SatuanPendidikan
    var total int64

    query := r.db.Model(&models.SatuanPendidikan{}).
        Preload("Provinsi").
        Preload("Kabupaten")

    // Apply filters
    if jenjang, ok := filters["jenjang"].(string); ok && jenjang != "" {
        query = query.Where("jenjang = ?", jenjang)
    }

    if provinsiID, ok := filters["provinsi_id"].(uint); ok && provinsiID > 0 {
        query = query.Where("provinsi_id = ?", provinsiID)
    }

    if kabupatenID, ok := filters["kabupaten_id"].(uint); ok && kabupatenID > 0 {
        query = query.Where("kabupaten_id = ?", kabupatenID)
    }

    if search, ok := filters["search"].(string); ok && search != "" {
        query = query.Where("MATCH(nama, alamat) AGAINST(? IN NATURAL LANGUAGE MODE)", search)
    }

    if akreditasi, ok := filters["akreditasi"].(string); ok && akreditasi != "" {
        query = query.Where("akreditasi = ?", akreditasi)
    }

    if status, ok := filters["status"].(string); ok && status != "" {
        query = query.Where("status = ?", status)
    } else {
        query = query.Where("status = ?", "aktif")
    }

    if verified, ok := filters["verified"].(bool); ok {
        query = query.Where("is_verified = ?", verified)
    }

    // Count total
    query.Count(&total)

    // Pagination
    offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

    // Order
    if sort, ok := filters["sort"].(string); ok && sort != "" {
        query = query.Order(sort)
    } else {
        query = query.Order("created_at DESC")
    }

    err := query.Find(&satpen).Error
    return satpen, total, err
}

func (r *satpenRepository) FindByID(id uint) (*models.SatuanPendidikan, error) {
    var satpen models.SatuanPendidikan
    err := r.db.Preload("Provinsi").
        Preload("Kabupaten").
        Preload("VerifiedBy").
        First(&satpen, id).Error
    return &satpen, err
}

func (r *satpenRepository) FindByNPSN(npsn string) (*models.SatuanPendidikan, error) {
    var satpen models.SatuanPendidikan
    err := r.db.Preload("Provinsi").
        Preload("Kabupaten").
        Where("npsn = ?", npsn).
        First(&satpen).Error
    return &satpen, err
}

func (r *satpenRepository) GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error) {
    // Implementation for statistics aggregation
    // This would involve multiple queries to aggregate data
    return nil, nil
}
```

### 2. Model (Golang)

```go
// internal/models/satuan_pendidikan.go
package models

import "time"

type SatuanPendidikan struct {
    ID              uint       `json:"id" gorm:"primaryKey"`
    NPSN            string     `json:"npsn" gorm:"uniqueIndex;size:20"`
    Nama            string     `json:"nama" gorm:"size:500"`
    Jenjang         string     `json:"jenjang" gorm:"type:enum('PAUD','MI','MTs','MA','Pesantren','Perguruan Tinggi')"`
    Alamat          string     `json:"alamat" gorm:"type:text"`
    KabupatenID     *uint      `json:"-"`
    Kabupaten       *Kabupaten `json:"kabupaten,omitempty" gorm:"foreignKey:KabupatenID"`
    ProvinsiID      uint       `json:"-"`
    Provinsi        *Provinsi  `json:"provinsi" gorm:"foreignKey:ProvinsiID"`
    KodePos         string     `json:"kode_pos,omitempty" gorm:"size:10"`
    KepalaSekolah   string     `json:"kepala_sekolah,omitempty" gorm:"size:255"`
    Email           string     `json:"email,omitempty" gorm:"size:255"`
    Phone           string     `json:"phone,omitempty" gorm:"size:20"`
    Website         string     `json:"website,omitempty" gorm:"size:500"`
    JumlahSiswa     uint       `json:"jumlah_siswa" gorm:"default:0"`
    JumlahGuru      uint       `json:"jumlah_guru" gorm:"default:0"`
    JumlahRombel    uint       `json:"jumlah_rombel" gorm:"default:0"`
    Akreditasi      string     `json:"akreditasi" gorm:"type:enum('A','B','C','Belum Terakreditasi');default:'Belum Terakreditasi'"`
    TahunBerdiri    int        `json:"tahun_berdiri,omitempty"`
    Latitude        *float64   `json:"-"`
    Longitude       *float64   `json:"-"`
    Coordinates     Coordinates `json:"coordinates,omitempty" gorm:"-"`
    Status          string     `json:"status" gorm:"type:enum('aktif','non-aktif');default:'aktif'"`
    IsVerified      bool       `json:"is_verified" gorm:"default:false"`
    VerifiedAt      *time.Time `json:"verified_at,omitempty"`
    VerifiedByID    *uint      `json:"-"`
    VerifiedBy      *User      `json:"verified_by,omitempty" gorm:"foreignKey:VerifiedByID"`
    CreatedAt       time.Time  `json:"created_at"`
    UpdatedAt       time.Time  `json:"updated_at"`
}

type Coordinates struct {
    Latitude  *float64 `json:"latitude,omitempty"`
    Longitude *float64 `json:"longitude,omitempty"`
}

func (s *SatuanPendidikan) AfterFind() error {
    if s.Latitude != nil && s.Longitude != nil {
        s.Coordinates = Coordinates{
            Latitude:  s.Latitude,
            Longitude: s.Longitude,
        }
    }
    return nil
}
```

---

## 📊 Performance Optimization

### 1. Database Indexing
```sql
-- Already created in schema
CREATE INDEX idx_jenjang ON satuan_pendidikan(jenjang);
CREATE INDEX idx_provinsi_id ON satuan_pendidikan(provinsi_id);
CREATE INDEX idx_kabupaten_id ON satuan_pendidikan(kabupaten_id);
CREATE INDEX idx_akreditasi ON satuan_pendidikan(akreditasi);
CREATE INDEX idx_status ON satuan_pendidikan(status);
CREATE INDEX idx_coordinates ON satuan_pendidikan(latitude, longitude);
CREATE FULLTEXT INDEX idx_search ON satuan_pendidikan(nama, alamat);

-- Composite index for common filters
CREATE INDEX idx_jenjang_provinsi ON satuan_pendidikan(jenjang, provinsi_id, status);
```

### 2. Caching Strategy
```go
// Cache satpen list per filter combination
cacheKey := fmt.Sprintf("satpen:j:%s:p:%d:page:%d", jenjang, provinsiID, page)
cache.Set(cacheKey, satpenData, 5*time.Minute)

// Cache statistics (longer TTL as it changes less frequently)
cache.Set("satpen:statistics", statsData, 1*time.Hour)

// Cache provinsi list (rarely changes)
cache.Set("master:provinsi", provinsiData, 24*time.Hour)
```

### 3. Query Optimization
- Use `LIMIT` and `OFFSET` for pagination
- Use `SELECT` specific columns when not all data is needed
- Use `Preload` carefully (avoid N+1 queries)
- Use database-level fulltext search for better performance

---

## 🎯 Integration with Frontend

### Example Usage (Next.js):
```javascript
// src/app/data-satpen/page.js
const fetchSatpen = async (filters) => {
  const queryParams = new URLSearchParams({
    page: filters.page || 1,
    limit: 20,
    ...(filters.jenjang && { jenjang: filters.jenjang }),
    ...(filters.provinsi && { provinsi: filters.provinsi }),
    ...(filters.search && { search: filters.search }),
  });

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/v1/satpen?${queryParams}`
  );

  return response.json();
};
```

---

## 🔒 Security & Rate Limiting

### Rate Limiting Configuration:
```yaml
# config.yaml
rate_limit:
  satpen:
    requests: 60
    window: 60  # 60 requests per minute
```

### Input Validation:
- Validate `page` and `limit` parameters
- Sanitize search query
- Validate enum values (jenjang, akreditasi, status)

---

## 📞 Support

**Related Files:**
- Main API Docs: `TODO BACKEND - READ ONLY API.md`
- Extension API Docs: `TODO BACKEND EXTEND - READ ONLY API.md`
- Database Schema: `database_schema.sql`
- Database Seeder: `database_seeder.sql`

---

**Last Updated:** 2025-01-14
**Version:** 1.0.0
**Priority:** HIGH (Core Feature)
