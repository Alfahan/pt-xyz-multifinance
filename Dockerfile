# Gunakan base image Go
FROM golang:1.23-alpine

# Set environment variable untuk Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja di dalam container
WORKDIR /app

# Copy go.mod dan go.sum untuk instalasi dependensi
COPY go.mod go.sum ./

# Jalankan instalasi dependensi
RUN go mod download

# Copy seluruh kode aplikasi ke dalam container
COPY . .

# Build aplikasi
RUN go build -o main ./cmd/main.go

# Expose port untuk aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]