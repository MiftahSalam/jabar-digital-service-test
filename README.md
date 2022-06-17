# jabar-digital-service-test
Aplikasi REST API user authentication sederhana sebagai syarat tes masuk Jabar Digital Service 2022
   
# Fitur
- User Register
- User Login
- Token Validation
- Authentication (JWT)
- Routing
- Input Validation
- Testing

# Demo
Aplikasi ini dideploy di heroku. [Cek demo](https://simple-api-dservice-jabar-test.herokuapp.com/api/v1) 

# Stacks
- Golang/Gin
- GORM (Golang ORM)
- Postgresql
- JWT-GO
- godotenv
- testify


# Getting started

## Install Golang

Pastikan Go versi 1.13 atau lebih tinggi sudah terinstal.

https://golang.org/doc/install

## Golang Environment Config

Pengaturan Go environment variable dapat dilihat di link https://golang.org/doc/install#install.

## Install Git
Download dan instal git versi terbaru di https://git-scm.com/downloads

## Install Postgresql
Download dan instal postgresql dan pgAdmin versi terbaru di https://www.postgresql.org/download/ dan https://www.pgadmin.org/download/

## Clone Project
Buka terminal/command prompt dan jalankan,:
```
git clone https://github.com/MiftahSalam/jabar-digital-service-test
```

## Install Dependencies
Setelah melakukan projek klon, masuk ke folder root project dan jalankan:
```
go mod tidy
```

## Create Database
Buat database dapat menggunakan pgAdmin atau command line. Berikut salah satu link tutorial yang bisa digunakan https://www.postgresqltutorial.com/postgresql-administration/postgresql-create-database/

## Application Environment Config
- Pastikan database sudah siap (terinstal dan terkonfigurasi)
- buat file .env file di project root
- buat dan isi variabel - variable dibawah di file .env (sesuaikan dengan parameter database yang terkonfigurasi)
```
JWT_SECRET=
JWT_EXPIRED_IN=

DATABASE_HOST=
DATABASE_PORT=
DATABASE_USERNAME=
DATABASE_PASSWORD=
DATABASE_NAME=
```

## Running Application
- Buka terminal dan masuk ke folder project root
- jalankan perintah
```
go run ./main.go
```

## Build Application
- Buka terminal dan masuk ke folder project root
- jalankan perintah
```
go build .
```

## Testing
Untuk sementara, karena tes menggunakan koneksi real database dan beberapa test case menggunakan data yangsama (yang bisa menyebabkan data konflik), maka testing hanya bisa dilakukan per package saja (contoh : model, service, router). Test root (go test ./...) tidak available. 

Contoh menjalankan test:
- buka terminal dan masuk ke folder package yang akan dites
- jalankan perintah
    ```
    go test ./users/model
    or
    go test ./users/service
    or
    go test ./users/router
    ```

## Deploy
Untuk deploy aplikasi Go di heroku dapat dilihat di link berikut https://devcenter.heroku.com/articles/getting-started-with-go
