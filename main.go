package main

import (
	"database/sql"
	"errors"
	"fmt"
	"laundy/entity"
	"regexp"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Whobay123@"
	dbname   = "laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	// MST CUSTOMER
	// customers := viewCustomers()
	// for _, customer := range customers {
	// 	fmt.Println(customer.Id, customer.Nama, customer.NoHp, customer.Alamat)
	// }

	// customer := entity.Customer{Id: 2, Nama: "Bian", NoHp: "087727726319", Alamat: "Bandung"}

	// insertCustomer(customer)
	// updateCustomer(customer)
	// deleteCustomer("2")

	// MST SERVICE
	// services := viewService()
	// for _, service := range services {
	// 	fmt.Println(service.Id, service.Pelayanan, service.Harga)
	// }

	// service := entity.Service{Id: 4, Pelayanan: "Laundry Karpet", Harga: "50000.00"}

	// insertService(service)
	// updateService(service)
	// deleteService("1")

	// MST TRANSACTION
	// transactions := viewTransaction()
	// for _, transaction := range transactions {
	// 	fmt.Println(transaction.Id, transaction.IdCustomer, transaction.TanggalMasuk, transaction.TanggalKEluar, transaction.DiterimaOleh)
	// }

	// transaction := entity.Transaction{Id: 2, IdCustomer: 2, TanggalMasuk: "2022-08-18", TanggalKEluar: "2022-08-20", DiterimaOleh: "Mirna"}

	// insertTransaction(transaction)

	// TX TRANSACTION ENROLLMENT
	transactionEnrollment := entity.TransactionEnrollment{Id: 4, IdTransaction: 2, IdService: 4, Jumlah: "1", Satuan: "Buah", Total: 50000}

	enrollmentSubject(transactionEnrollment)

}

// MST_CUSTOMER
func viewCustomers() []entity.Customer {
	db := connectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM mst_customer;"

	rows, err := db.Query(sqlStatment)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	customers := scanCustomer(rows)

	return customers
}

func scanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Nama, &customer.NoHp, &customer.Alamat)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return customers
}

func validateCustomer(customer entity.Customer) error {
	if customer.Nama == "" {
		return errors.New("Nama Tidak Boleh Kosong")
	}

	if customer.Alamat == "" {
		return errors.New("Alamat Tidak Boleh Kosong")
	}

	re := regexp.MustCompile(`^\+?(\d.*){3,}$`)
	if !re.MatchString(customer.NoHp) {
		return errors.New("Nomor Hp harus Nomor yang Valid")
	}

	return nil
}

func insertCustomer(customer entity.Customer) {
	if err := validateCustomer(customer); err != nil {
		fmt.Println("Error:", err)
		return
	}

	db := connectDb()
	defer db.Close()

	sqlStatment := "INSERT INTO mst_customer (id, nama, no_hp, alamat) VALUES ($1, $2, $3, $4)"

	_, err := db.Exec(sqlStatment, customer.Id, customer.Nama, customer.NoHp, customer.Alamat)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

func updateCustomer(customer entity.Customer) {
	if err := validateCustomer(customer); err != nil {
		fmt.Println("Error:", err)
		return
	}

	db := connectDb()
	defer db.Close()

	sqlStatment := "UPDATE mst_customer SET nama = $2, no_hp = $3, alamat = $4 WHERE id = $1;"

	_, err := db.Exec(sqlStatment, customer.Id, customer.Nama, customer.NoHp, customer.Alamat)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

func deleteCustomer(id string) {
	db := connectDb()
	defer db.Close()

	sqlStatment := "DELETE FROM mst_customer WHERE id = $1;"

	_, err := db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Delete Data!")
	}
}

// MST_SERVICE
func viewService() []entity.Service {
	db := connectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM mst_service;"

	rows, err := db.Query(sqlStatment)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	services := scanService(rows)

	return services
}

func scanService(rows *sql.Rows) []entity.Service {
	services := []entity.Service{}
	var err error

	for rows.Next() {
		service := entity.Service{}
		err := rows.Scan(&service.Id, &service.Pelayanan, &service.Harga)
		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return services
}

func validateService(service entity.Service) error {
	if service.Pelayanan == "" {
		return errors.New("Pelayanan Tidak Boleh Kosong")
	}

	harga, err := strconv.ParseFloat(service.Harga, 64)
	if err != nil || harga <= 0 {
		return errors.New("Harga Harus Berupa Angka")
	}

	return nil
}

func insertService(service entity.Service) {
	if err := validateService(service); err != nil {
		fmt.Println("Error:", err)
		return
	}

	db := connectDb()
	defer db.Close()
	var err error

	sqlStatment := "INSERT INTO mst_service (id, pelayanan, harga) VALUES ($1, $2, $3);"

	_, err = db.Exec(sqlStatment, service.Id, service.Pelayanan, service.Harga)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Inser Data!")
	}
}

func updateService(service entity.Service) {
	if err := validateService(service); err != nil {
		fmt.Println("Error:", err)
		return
	}

	db := connectDb()
	defer db.Close()
	var err error

	sqlStatment := "UPDATE mst_service SET pelayanan = $2, harga = $3 WHERE id = $1;"

	_, err = db.Exec(sqlStatment, service.Id, service.Pelayanan, service.Harga)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Seuccessfully Update Data!")
	}
}

func deleteService(id string) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatment := "DELETE FROM mst_service WHERE id = $1;"

	_, err = db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

// MST_TRANSACTION
func viewTransaction() []entity.Transaction {
	db := connectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM mst_transaction;"

	rows, err := db.Query(sqlStatment)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	transactions := scanTransaction(rows)

	return transactions
}

func scanTransaction(rows *sql.Rows) []entity.Transaction {
	transactions := []entity.Transaction{}
	var err error

	for rows.Next() {
		transaction := entity.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.IdCustomer, &transaction.TanggalMasuk, &transaction.TanggalMasuk, &transaction.DiterimaOleh)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return transactions
}

func validateTransaction(transaction entity.Transaction) error {
	db := connectDb()
	defer db.Close()

	var idExists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_customer WHERE id = $1)", transaction.IdCustomer).Scan(&idExists)
	if err != nil {
		return errors.New("Gagal Memeriksa Id Customer")
	}
	if !idExists {
		return errors.New("Id Customer Tidak Ditemukan!")
	}

	tglMasuk, err := time.Parse("2006-01-02", transaction.TanggalMasuk)
	if err != nil {
		return errors.New("Format Tanggal masuk Tidak Valid!")
	}
	tglKeluar, err := time.Parse("2006-01-02", transaction.TanggalKEluar)
	if err != nil {
		return errors.New("Format tanggal Keluar Tidak Valid!")
	}
	if tglMasuk.After(tglKeluar) {
		return errors.New("Tanggal Masuk Harus Sebelum tanggal Keluar")
	}
	if transaction.DiterimaOleh == "" {
		return errors.New("Penerima Tidak Boleh Kosong")
	}

	var idExistsInTrans bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_transaction WHERE id = $1)", transaction.Id).Scan(&idExistsInTrans)
	if err != nil {
		return errors.New("Gagal Memeriksa Id Transaction")
	}
	if idExistsInTrans {
		return errors.New("Id Sudah Ada dalam Transaction")
	}

	return nil
}

func insertTransaction(transaction entity.Transaction) {
	if err := validateTransaction(transaction); err != nil {
		fmt.Println("Error:", err)
		return
	}

	db := connectDb()
	defer db.Close()
	var err error

	sqlStatment := "INSERT INTO mst_transaction (id , id_customer, tanggal_masuk, tanggal_keluar, diterima_oleh) VALUES ($1, $2, $3, $4, $5);"

	_, err = db.Exec(sqlStatment, transaction.Id, transaction.IdCustomer, transaction.TanggalMasuk, transaction.TanggalKEluar, transaction.DiterimaOleh)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data!")
	}
}

// TX TRANSACTION
func enrollmentSubject(transactionEnrollment entity.TransactionEnrollment) {
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	insertTransactionEnrolment(transactionEnrollment, tx)

	takenTotal := getSumTotalOfTransaction(transactionEnrollment.IdTransaction, tx)

	updateTransactionEnrolment(takenTotal, transactionEnrollment.IdTransaction, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}
}

func insertTransactionEnrolment(transactionEnrollment entity.TransactionEnrollment, tx *sql.Tx) {
	insertTransactionEnrolment := "INSERT INTO tx_transaction_enrollment(id, id_transaction, id_service, jumlah, satuan, total) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := tx.Exec(insertTransactionEnrolment, transactionEnrollment.Id, transactionEnrollment.IdTransaction, transactionEnrollment.IdService, transactionEnrollment.Jumlah, transactionEnrollment.Satuan, transactionEnrollment.Total)
	validate(err, "Insert", tx)
}

func getSumTotalOfTransaction(id int, tx *sql.Tx) int {
	sumTotal := "SELECT SUM(total) FROM tx_transaction_enrollment WHERE id_transaction = $1;"

	takenTotal := 0
	err := tx.QueryRow(sumTotal, id).Scan(&takenTotal)
	validate(err, "Select", tx)

	return takenTotal
}

func updateTransactionEnrolment(takenTotal int, transactionId int, tx *sql.Tx) {
	updateTransactionEnrolment := "UPDATE mst_transaction SET total_harga = $1 WHERE id = $2;"

	_, err := tx.Exec(updateTransactionEnrolment, takenTotal, transactionId)
	validate(err, "Update", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback!")
	} else {
		fmt.Println("Successfully " + message + " data!")
	}
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Conncted!")
	}

	return db
}
