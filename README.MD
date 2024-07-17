# Enigma Laundry

## Deskripsi
Submission Enigma Laundry Mengelola Customer, Service dan Transaction laundry
Aplikasi ini di buat mengumakan Golang

## Prasyarat
1. Golang harus terinstal di sistem Anda
2. PostgreSQL harus terinstal di sistem Anda

## Instalasi
1. Install dependencies menggunakan `go get -u github.com/lib/pq`:

## Konfigurasi
1. Pastikan PostgreSQL berjalan di sistem Anda dan buat database dengan nama `enigma_laundry`.
2. Tabel-tabel yang diperlukan dalam database `enigma_laundry`:
    ```sql
    CREATE TABLE mst_customer (
        id SERIAL PRIMARY KEY,
        nama VARCHAR(100) NOT NULL,
        no_hp VARCHAR(15) NOT NULL,
        alamat VARCHAR(20) NOT NULL
    );

    CREATE TABLE mst_service (
        id SERIAL PRIMARY KEY,
        pelayanan VARCHAR(100) NOT NULL,
        harga VARCHAR(20) NOT NULL
    );

   CREATE TABLE mst_transaction (
	id SERIAL PRIMARY KEY,
	id_customer INT REFERENCES mst_customer(id),
	tanggal_masuk DATE NOT NULL,
	tanggal_keluar DATE NOT NULL,
	diterima_oleh VARCHAR(100) NOT NULL,
	total_harga INT default 0
    );

    CREATE TABLE tx_transaction_enrollment (
        id SERIAL PRIMARY KEY,
        id_transaction INT NOT NULL,
        id_service INT NOT NULL,
        jumlah VARCHAR(20) NOT NULL,
        satuan VARCHAR(20) NOT NULL,
        total VARCHAR(20) NOT NULL,
        FOREIGN KEY (id_transaction) REFERENCES mst_transaction (id),
        FOREIGN KEY (id_service) REFERENCES mst_service (id)
    );
    ```
3. Perbarui konfigurasi koneksi database dalam file `main.go` sesuai dengan pengaturan PostgreSQL Anda:
    ```go
    const (
        host     = "localhost"
        port     = 5432
        user     = "postgres"
        password = "YourPassword"
        dbname   = "enigma_laundry"
    )
    ```

## Menjalankan Aplikasi
1. Jalankan aplikasi menggunakan perintah berikut:

   - go run main.go

## Penggunaan
Aplikasi ini memiliki beberapa fungsi utama untuk mengelola data pelanggan, layanan, dan transaksi. Fungsi-fungsi tersebut dapat dipanggil di dalam `main` function di file `main.go`.

### Mengelola Customer
- Melihat daftar Customer:
    ```go
    customers := viewCustomers()
    for _, customer := range customers {
        fmt.Println(customer.Id, customer.Nama, customer.NoHp, customer.Alamat)
    }
    ```
- Menambahkan Customer baru:
    ```go
    customer := entity.Customer{Id: 2, Nama: "Bian", NoHp: "08772772986", Alamat : "Bandung"}
    insertCustomer(customer)
    ```
- Memperbarui data Customer:
    ```go
    updateCustomer(customer)
    ```
- Menghapus Customer:
    ```go
    deleteCustomer("2")
    ```

### Mengelola Service
- Melihat daftar Service:
    ```go
    services := viewService()
    for _, service := range services {
        fmt.Println(service.Id, service.Pelayanan, service.Harga)
    }
    ```
- Menambahkan Service baru:
    ```go
    service := entity.Service{Id: 4, Pelayanan: "Cuci Karpet", Harga: "25000.00"}
    insertService(service)
    ```
- Memperbarui data Service:
    ```go
    updateService(service)
    ```
- Menghapus Service:
    ```go
    deleteService("4")
    ```

### Mengelola Transaction
- Melihat daftar Transaction:
    ```go
    transactions := viewTransaction()
    for _, transaction := range transactions {
        fmt.Println(transaction.Id, transaction.IdCustomer, transaction.TanggalMasuk, transaction.TanggalKEluar, transaction.DiterimaOleh)
    }
    ```
- Menambahkan Transaction baru:
    ```go
    transaction := entity.Transaction{Id: 2, IdCustomer: 2, TanggalMasuk: "2022-08-18", TanggalKEluar: "2022-08-20", DiterimaOleh: "Mirna"}
    insertTransaction(transaction)
    ```

### Enrollment Transaksi
- Menambahkan enrolmen Transaction:
    ```go
    transactionEnrollment := entity.TransactionEnrollment{Id: 4, IdTransaction: 2, IdService: 4, Jumlah: "1", Satuan: "Buah", Total: "25000.00"}
    enrollmentSubject(transactionEnrollment)
    ```

## Kesimpulan
Aplikasi Enigma Laundry ini menyediakan fungsi-fungsi dasar untuk mengelola data Customer, Service, dan Transaction.
