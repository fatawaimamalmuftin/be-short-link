<div align="center">

# 🔗 ShortLink Backend API

  <p>
    <b>API Service Pemendek URL performa tinggi yang dibangun dengan Go, PostgreSQL, Redis, dan Docker.</b>
  </p>

  <p>
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
    <img src="https://img.shields.io/badge/Gin-008080?style=for-the-badge&logo=gin&logoColor=white" alt="Gin" />
    <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
    <img src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis" />
    <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  </p>

</div>

---

## 🚀 Tech Stack

* **Core Engine:** Go (Golang) v1.24
* **Web Framework:** Gin-Gonic
* **Database:** PostgreSQL 18 (via `pgx/v5` Connection Pool)
* **In-Memory Cache & Session:** Redis 8.10.2 (via `go-redis/v9`)
* **Authentication:** JWT (JSON Web Token) dengan Custom Claims & Issuer Validation
* **Password Hashing:** Argon2id (M=: 64MB, T: 2, P: 2)
* **Containerization:** Docker & Docker Compose
* **Database Migration:** Golang Migrate CLI

---

## ✨ Fitur Utama

| Fitur | Deskripsi | Status |
| :--- | :--- | :---: |
| **Authentication** | Registrasi (Argon2id hashing) & Login JWT | ✅ |
| **Token Blacklist** | Invalidation JWT Token saat Logout menggunakan Redis | ✅ |
| **Link Management** | Memuat daftar short link milik user (`GET /api/links/my`) | ✅ |
| **URL Shortener** | Membuat short link baru dengan kustomisasi slug unik | ✅ |
| **Delete Link** | Menghapus short link berdasarkan ID dengan pengecekan otorisasi | ✅ |
| **Security** | Middleware CORS & Protected Routes via JWT Verification | ✅ |

---

## 🛠️ Ringkasan API Endpoints

| Method | Endpoint | Description | Auth Required |
| :---: | :--- | :--- | :---: |
| `POST` | `/api/auth/registration` | Registrasi user baru | ❌ |
| `POST` | `/api/auth/login` | Login user & dapatkan JWT Token | ❌ |
| `POST` | `/api/auth/logout` | Logout & masukkan token ke blacklist | YES |
| `GET` | `/api/links/my` | Mengambil seluruh link milik user | YES |
| `POST` | `/api/links` | Membuat short link baru | YES |
| `DELETE` | `/api/links/:id` | Menghapus short link berdasarkan ID | YES |

---

## ⚙️ Langkah Instalasi

### 1. Clone Repository
```bash
git clone [https://github.com/fatawaimamalmuftin/be-shortlink.git](https://github.com/fatawaimamalmuftin/be-shortlink.git)
cd be-shortlink