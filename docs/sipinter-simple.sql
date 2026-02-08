-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for testing_lpmaarif1
CREATE DATABASE IF NOT EXISTS `testing_lpmaarif1` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `testing_lpmaarif1`;

-- Dumping structure for table testing_lpmaarif1.jenjang_pendidikan
CREATE TABLE IF NOT EXISTS `jenjang_pendidikan` (
  `id_jenjang` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nm_jenjang` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `keterangan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lembaga` enum('MADRASAH','SEKOLAH') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_jenjang`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.kabupaten
CREATE TABLE IF NOT EXISTS `kabupaten` (
  `id_kab` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_prov` bigint unsigned NOT NULL,
  `nama_kab` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_kab`),
  KEY `kabupaten_id_prov_foreign` (`id_prov`),
  CONSTRAINT `kabupaten_id_prov_foreign` FOREIGN KEY (`id_prov`) REFERENCES `provinsi` (`id_prov`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=516 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.kategori_satpen
CREATE TABLE IF NOT EXISTS `kategori_satpen` (
  `id_kategori` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nm_kategori` enum('A','B','C','D') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `konotasi` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `keterangan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_kategori`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.migrations
CREATE TABLE IF NOT EXISTS `migrations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.pdptk
CREATE TABLE IF NOT EXISTS `pdptk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_satpen` bigint unsigned DEFAULT NULL,
  `tapel` varchar(10) DEFAULT NULL,
  `pd_lk` int DEFAULT NULL,
  `pd_pr` int DEFAULT NULL,
  `jml_pd` int DEFAULT NULL,
  `guru_lk` int DEFAULT NULL,
  `guru_pr` int DEFAULT NULL,
  `jml_guru` int DEFAULT NULL,
  `tendik_lk` int DEFAULT NULL,
  `tendik_pr` int DEFAULT NULL,
  `jml_tendik` int DEFAULT NULL,
  `last_sinkron` datetime DEFAULT NULL,
  `status_sinkron` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_satpen` (`id_satpen`),
  KEY `tapel` (`tapel`),
  CONSTRAINT `FK_pdptk_satpen` FOREIGN KEY (`id_satpen`) REFERENCES `satpen` (`id_satpen`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_pdptk_tahun_pelajaran` FOREIGN KEY (`tapel`) REFERENCES `tahun_pelajaran` (`tapel_dapo`)
) ENGINE=InnoDB AUTO_INCREMENT=4065 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.pengurus_cabang
CREATE TABLE IF NOT EXISTS `pengurus_cabang` (
  `id_pc` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_prov` bigint unsigned NOT NULL,
  `kode_kab` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_pc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_pc`),
  KEY `pengurus_cabang_id_prov_foreign` (`id_prov`),
  CONSTRAINT `pengurus_cabang_id_prov_foreign` FOREIGN KEY (`id_prov`) REFERENCES `provinsi` (`id_prov`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=522 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.personal_access_tokens
CREATE TABLE IF NOT EXISTS `personal_access_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint unsigned NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `expires_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.profile_pengurus_cabang
CREATE TABLE IF NOT EXISTS `profile_pengurus_cabang` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_pc` bigint unsigned NOT NULL,
  `alamat` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `kelurahan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `kecamatan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `kabupaten` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `lintang` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `bujur` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `website` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ketua` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_ketua` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `wakil_ketua` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_wakil` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `bendahara` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_bendahara` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sekretaris` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_sekretaris` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `masa_khidmat` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id_pc` (`id_pc`) USING BTREE,
  CONSTRAINT `FK_profile_pengurus_cabang_pengurus_cabang` FOREIGN KEY (`id_pc`) REFERENCES `pengurus_cabang` (`id_pc`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.profile_pengurus_wilayah
CREATE TABLE IF NOT EXISTS `profile_pengurus_wilayah` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_pw` bigint unsigned NOT NULL,
  `alamat` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `kelurahan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `kecamatan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `kabupaten` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `lintang` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `bujur` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `website` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ketua` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_ketua` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `wakil_ketua` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_wakil` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `bendahara` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_bendahara` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sekretaris` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `telp_sekretaris` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `masa_khidmat` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id_pw` (`id_pw`) USING BTREE,
  CONSTRAINT `FK_profile_pengurus_wilayah_provinsi` FOREIGN KEY (`id_pw`) REFERENCES `provinsi` (`id_prov`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.provinsi
CREATE TABLE IF NOT EXISTS `provinsi` (
  `id_prov` bigint unsigned NOT NULL AUTO_INCREMENT,
  `map` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kode_prov` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nm_prov` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_prov`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.ptk
CREATE TABLE IF NOT EXISTS `ptk` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_npyp` int DEFAULT NULL,
  `id_satpen` bigint unsigned NOT NULL,
  `nik` varchar(16) NOT NULL,
  `nama_ptk` varchar(255) NOT NULL,
  `tempat_lahir` varchar(255) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `jenis_kelamin` enum('Laki-Laki','Perempuan') NOT NULL,
  `nama_ibu` varchar(255) NOT NULL,
  `agama` enum('Islam','Kristen','Katolik','Hindu','Buddha','Konghucu') NOT NULL,
  `kebutuhan_khusus` enum('Tidak ada','A - Tuna Netra','B - Tuna Rungu','C - Tuna Grahita Ringan','C1 - Tuna Grahita Sedang','D - Tuna Daksa Ringan','E - Tuna Laras','F - Tuna Wicara','H - Hiperaktif','I - Cerdas Istimewa','J - Bakat Istimewa','K - Kesulitan Belajar','N - Narkoba','O - Indigo','P - Down Sindrome','Q - Autis','Lainnya') DEFAULT 'Tidak ada',
  `status_perkawinan` enum('Menikah','Belum Menikah','Duda atau Lajang') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(255) NOT NULL,
  `kabupaten_kota` varchar(255) NOT NULL,
  `kecamatan` varchar(255) NOT NULL,
  `desa_kelurahan` varchar(255) NOT NULL,
  `alamat` text NOT NULL,
  `kode_pos` varchar(5) NOT NULL,
  `jenis_ptk` enum('Guru Kelas','Guru Mapel','Guru BK','Guru Pendamping Khusus','Tenaga Administrasi Sekolah','Guru TIK','Laboran','Tenaga Perpustakaan','Academic Advisor','Academic Spesialis','Curiculum Development Advisor','Kindegarten Teacher','Management Advisor','Playgroup Teacher','Principal','Teaching Assistant','Vice Principal','Tukang Kebun','Penjaga Sekolah','Petugas Keamanan','Pesuruh/Office Boy','Kepala Sekolah','Terapis','Guru Pengganti','Pengawas Paud Dikmas','Penilik','Guru Pembimbing Khusus','Instruktur Kejuruan','Instruktur','Penguji','Master Penguji','Tutor','Pamong Belajar','Tenaga Kependidikan','Pengawas') NOT NULL,
  `status_kepegawaian` enum('PNS','PNS Diperbantukan','PNS Depag','GTY/PTY','Honor Daerah Tk. 1 Provinsi','Honor Daerah Tk. 2 Kab/Kota','Guru Honor Sekolah','Tenaga Honor Sekolah','CPNS','PPPK','PPNPN','Guru Pengganti','Kontrak Kerja WNA') NOT NULL,
  `nip` varchar(50) DEFAULT NULL,
  `lembaga_pengangkat` enum('Pemerintah Pusat','Pemerintah Provinsi','Pemerintah Kab/Kota','Ketua Yayasan','Kepala Sekolah','Lainnya') NOT NULL,
  `no_sk_pengangkatan` varchar(255) NOT NULL,
  `tmt_pengangkatan` date NOT NULL,
  `sumber_gaji` enum('APBN','APBD Provinsi','APBD Kab/Kota','Yayasan','Sekolah','Lembaga Donor','Lainnya') NOT NULL,
  `lisensi_kepala_sekolah` enum('Sudah','Belum') DEFAULT 'Belum',
  `nomor_surat_tugas` varchar(255) NOT NULL,
  `tanggal_surat_tugas` date NOT NULL,
  `tmt_tugas` date NOT NULL,
  `upload_sk` varchar(255) NOT NULL,
  `status_ajuan` enum('verifikasi','revisi','proses','approve','dikeluarkan') DEFAULT 'verifikasi',
  `tanggal_verifikasi` datetime DEFAULT NULL,
  `tanggal_revisi` datetime DEFAULT NULL,
  `tanggal_proses` datetime DEFAULT NULL,
  `tanggal_approve` datetime DEFAULT NULL,
  `tanggal_dikeluarkan` datetime DEFAULT NULL,
  `keterangan_revisi` text,
  `nomor_sk_keluar` varchar(255) DEFAULT NULL,
  `catatan_verifikator` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nik` (`nik`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_ptk_satpen` (`id_satpen`),
  KEY `idx_ptk_status` (`status_ajuan`),
  KEY `idx_ptk_nik` (`nik`),
  KEY `idx_ptk_email` (`email`),
  KEY `idx_ptk_created` (`created_at`),
  KEY `idx_ptk_jenis` (`jenis_ptk`),
  KEY `FK_ptk_npyp` (`id_npyp`),
  CONSTRAINT `FK_ptk_npyp` FOREIGN KEY (`id_npyp`) REFERENCES `npyp` (`id_npyp`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_ptk_satpen` FOREIGN KEY (`id_satpen`) REFERENCES `satpen` (`id_satpen`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.ptk_status_history
CREATE TABLE IF NOT EXISTS `ptk_status_history` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptk_id` bigint unsigned NOT NULL,
  `status_from` enum('verifikasi','revisi','proses','approve','dikeluarkan') DEFAULT NULL,
  `status_to` enum('verifikasi','revisi','proses','approve','dikeluarkan') NOT NULL,
  `keterangan` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_ptk_status_history_ptk` (`ptk_id`),
  KEY `idx_ptk_status_history_date` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.satpen
CREATE TABLE IF NOT EXISTS `satpen` (
  `id_satpen` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` bigint unsigned NOT NULL,
  `id_prov` bigint unsigned NOT NULL,
  `id_kab` bigint unsigned NOT NULL,
  `id_pc` bigint unsigned NOT NULL,
  `id_kategori` bigint unsigned DEFAULT NULL,
  `id_jenjang` bigint unsigned NOT NULL,
  `npsn` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `no_registrasi` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `no_urut` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nm_satpen` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `yayasan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `kepsek` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `telpon` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `fax` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `thn_berdiri` year DEFAULT NULL,
  `kecamatan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `kelurahan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `alamat` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `aset_tanah` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nm_pemilik` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tgl_registrasi` datetime NOT NULL,
  `actived_date` datetime DEFAULT NULL,
  `status` enum('permohonan','revisi','proses dokumen','setujui','expired','perpanjangan') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_satpen`),
  UNIQUE KEY `satpen_id_user_unique` (`id_user`),
  UNIQUE KEY `satpen_npsn_unique` (`npsn`),
  UNIQUE KEY `satpen_no_registrasi_unique` (`no_registrasi`),
  UNIQUE KEY `satpen_no_urut_unique` (`no_urut`),
  KEY `satpen_id_prov_foreign` (`id_prov`),
  KEY `satpen_id_kab_foreign` (`id_kab`),
  KEY `satpen_id_pc_foreign` (`id_pc`),
  KEY `satpen_id_kategori_foreign` (`id_kategori`),
  KEY `satpen_id_jenjang_foreign` (`id_jenjang`),
  CONSTRAINT `satpen_id_jenjang_foreign` FOREIGN KEY (`id_jenjang`) REFERENCES `jenjang_pendidikan` (`id_jenjang`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `satpen_id_kab_foreign` FOREIGN KEY (`id_kab`) REFERENCES `kabupaten` (`id_kab`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `satpen_id_kategori_foreign` FOREIGN KEY (`id_kategori`) REFERENCES `kategori_satpen` (`id_kategori`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `satpen_id_pc_foreign` FOREIGN KEY (`id_pc`) REFERENCES `pengurus_cabang` (`id_pc`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `satpen_id_prov_foreign` FOREIGN KEY (`id_prov`) REFERENCES `provinsi` (`id_prov`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `satpen_id_user_foreign` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9363 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.settings
CREATE TABLE IF NOT EXISTS `settings` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `describe` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `lookup` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `settings_lookup_unique` (`lookup`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.tahun_pelajaran
CREATE TABLE IF NOT EXISTS `tahun_pelajaran` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tapel_dapo` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `nama_tapel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tapel_dapo` (`tapel_dapo`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.timeline_reg
CREATE TABLE IF NOT EXISTS `timeline_reg` (
  `id_timeline` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_satpen` bigint unsigned NOT NULL,
  `status_verifikasi` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tgl_status` datetime NOT NULL,
  `keterangan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_timeline`),
  KEY `timeline_reg_id_satpen_foreign` (`id_satpen`),
  CONSTRAINT `timeline_reg_id_satpen_foreign` FOREIGN KEY (`id_satpen`) REFERENCES `satpen` (`id_satpen`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=12268 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table testing_lpmaarif1.users
CREATE TABLE IF NOT EXISTS `users` (
  `id_user` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` enum('super admin','admin pusat','admin wilayah','admin cabang','operator') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `status_active` enum('active','block') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `provId` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cabangId` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=9791 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
