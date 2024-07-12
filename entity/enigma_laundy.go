package entity

type Customer struct {
	Id   int
	Nama string
	NoHp string
}

type Transaction struct {
	Id            int
	IdCustomer    int
	TanggalMasuk  string
	TanggalKEluar string
	DiterimaOleh  string
}

type Service struct {
	Id        int
	Pelayanan string
	Harga     string
}

type TransactionEnrollment struct {
	Id            int
	IdTransaction int
	IdService     int
	Jumlah        string
	Satuan        string
	Total         string
}
