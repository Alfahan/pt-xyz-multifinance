# PT-XYZ-MULTIFINANCE

Sistem backend untuk PT-XYZ Multifinance, dikembangkan dengan pendekatan **Clean Architecture** dan bahasa pemrograman **Go**.  
Project ini mendukung pengembangan scalable, maintainable, dan mudah di-deploy.

---

## Struktur Project

```
.
├── cmd/           # Entry point aplikasi (main.go)
├── config/        # Konfigurasi aplikasi (.env, config.go)
├── internal/      # Source code utama (domain, handler, repository, usecase, middleware)
├── pkg/           # Library utilitas yang reusable (database, redis, kafka, dsb)
├── docs/          # Dokumentasi API (Swagger/OpenAPI)
├── test/          # Unit & integration test
├── tmp/           # File sementara (binary hasil build, dsb)
├── .env           # Environment variables
├── .air.toml      # Konfigurasi Air (hot reload)
├── go.mod
├── go.sum
└── README.md
```

---

## Cara Menjalankan

1. **Install dependencies**
    ```sh
    go mod tidy
    ```

2. **Jalankan aplikasi**
    ```sh
    go run cmd/main.go
    ```

3. **Set environment variables**
    - Edit file `.env` sesuai kebutuhan database, redis, kafka, dsb.

---

## Menjalankan dengan Docker

Aplikasi dapat dijalankan secara cepat menggunakan Docker & Docker Compose tanpa perlu install Go, PostgreSQL, atau Redis secara manual.

### 1. Siapkan file `.env`
Pastikan file `.env` telah terisi sesuai kebutuhan konfigurasi.

### 2. Build dan Jalankan Aplikasi

```sh
docker compose up --build
```
Perintah ini akan:
- Build image aplikasi Go
- Menyediakan service untuk PostgreSQL dan Redis (jika dikonfigurasi di `docker-compose.yml`)
- Menghubungkan seluruh service dalam satu jaringan

### 3. Mengakses Aplikasi

- API: `http://localhost:8080`
- Swagger UI: `http://localhost:8080/swagger/index.html`  
  (Proteksi basic auth, user & password lihat `.env`)

> **Tips:**  
> Untuk development, gunakan `docker compose up --build` agar perubahan kode terupdate di container.

### 4. Perintah Docker Lainnya

- **Hentikan service:**
  ```sh
  docker compose down
  ```
- **Lihat log service:**
  ```sh
  docker compose logs -f
  ```
- **Masuk ke container aplikasi:**
  ```sh
  docker compose exec app sh
  ```

### 5. Contoh Struktur `docker-compose.yml`

```yaml
version: "3.8"
services:
  app:
    build: .
    env_file: .env
    ports:
      - "8080:8080"
    depends_on:
      - db
      - redis
  db:
    image: postgres:16
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME}
    ports:
      - "5432:5432"
    volumes:
      - dbdata:/var/lib/postgresql/data
  redis:
    image: redis:7
    ports:
      - "6379:6379"
    volumes:
      - redisdata:/data
volumes:
  dbdata:
  redisdata:
```

---

## Konfigurasi `.env` Contoh

```env
# Database configuration
DB_HOST=localhost
DB_PORT=5454
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=xyzdb


# Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Kafka
KAFKA_BROKERS=localhost:9092
KAFKA_CLIENT_ID=pt-xyz-app

# Server configuration
SERVER_PORT=8081

# Swagger user & password
SWAGGER_USER=admin
SWAGGER_PASS=password

# JWT_SECRET
JWT_SECRET=123kaslmca
```

---

## Hot Reload (Development)

Instalasi dan penggunaan [Air](https://github.com/air-verse/air):

1. **Install Air**
    ```sh
    go install github.com/air-verse/air@latest
    ```
2. **Jalankan dengan hot reload**
    ```sh
    air
    ```

> **Otomatis Generate Swagger:**  
> Project ini sudah dikonfigurasi agar setiap ada perubahan kode, dokumentasi Swagger (`docs/`) akan digenerate otomatis sebelum build. Lihat bagian [Swagger Otomatis](#swagger-otomatis).

---

## Testing

```sh
go test ./test/...
```

---

## Dokumentasi API (Swagger)

Dokumentasi API tersedia di folder `docs/` (Swagger/OpenAPI).  
Swagger akan **tergenerate otomatis setiap ada penambahan/ubah route** (lihat konfigurasi Air di bawah).

- **Akses Swagger UI:**  
  Jalankan aplikasi, lalu akses [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- **Proteksi:**  
  Swagger UI diproteksi Basic Auth, username dan password dari `.env` (`SWAGGER_USER`, `SWAGGER_PASS`).

### Menambah/Update Dokumentasi API

1. Tambahkan anotasi Swagger pada handler (contoh):
    ```go
    // @Summary      Health Check
    // @Description  Cek status service
    // @Tags         Utility
    // @Success      200  {object}  map[string]interface{}
    // @Router       /health [get]
    func HealthCheck(c echo.Context) error { ... }
    ```
2. **Swagger akan tergenerate otomatis** saat development dengan Air, atau bisa generate manual:
    ```sh
    swag init --dir cmd,internal
    ```

### Swagger Otomatis

Sudah dikonfigurasi di `.air.toml`:
```toml
[build]
  pre_cmd = "swag init --dir cmd,internal"
  cmd = "go build -o ./tmp/main.exe ./cmd/main.go"
  bin = "tmp/main.exe"
```
Setiap ada perubahan kode, dokumentasi API akan **selalu update**.

---

## Aturan Git & Commit

Agar kolaborasi dan histori kode rapi, gunakan aturan berikut:

### 1. **Branching**
- Gunakan branch terpisah untuk setiap fitur/bugfix, misal:  
  - `feature/nama-fitur`
  - `bugfix/penjelasan-singkat-bug`
- Jangan commit langsung ke `main` kecuali untuk hotfix sangat kritis.

### 2. **Format Pesan Commit**
Gunakan format konvensional agar mudah dilacak, seperti:
```
<type>(<scope>): <short summary>

# Contoh:
feat(auth): implementasi login JWT
fix(user): validasi data user ketika register
docs(readme): update dokumentasi cara setup
refactor(repo): perbaiki struktur repository user
test(api): tambah unit test untuk endpoint health
chore(ci): update workflow pipeline
```
- **type**: feat, fix, docs, refactor, test, chore, style, perf, build, ci, etc.
- **scope**: bagian/folder/fitur terkait (opsional, gunakan jika perlu)
- **summary**: singkat & jelas (≤50 karakter)

### 3. **Best Practice**
- Satu commit = satu perubahan logis/atomic (hindari commit campur aduk).
- Pull request harus jelas tujuannya, deskripsi lengkap, dan mention issue/fitur jika relevan.
- Lakukan review sebelum merge ke `main`.
- Selalu update branch dari `main` sebelum membuat PR.

### 4. **Contoh Alur Kerja**
1. Buat branch baru:
    ```sh
    git checkout -b feature/login-jwt
    ```
2. Commit perubahan:
    ```sh
    git add .
    git commit -m "feat(auth): implementasi login JWT"
    ```
3. Push branch:
    ```sh
    git push origin feature/login-jwt
    ```
4. Buat Pull Request ke `main`, sertakan deskripsi dan link issue (jika ada).

---

## Kontribusi

1. Fork repository ini
2. Buat branch fitur/bugfix
3. Commit perubahan dengan pesan yang jelas
4. Pull request ke branch `main`

---

## Lisensi

Lisensi terbuka untuk pengembangan internal PT-XYZ Multifinance.