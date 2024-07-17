CREATE TABLE mst_customer (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(100) NOT NULL,
	no_hp VARCHAR(15) NOT NULL,
	alamat VARCHAR(20) NOT NULL
);

CREATE TABLE mst_service (
	id SERIAL PRIMARY KEY,
	pelayanan VARCHAR(100) NOT NULL,
	harga NUMERIC(10, 2) NOT NULL
);

CREATE TABLE mst_transaction (
	id SERIAL PRIMARY KEY,
	id_customer INT REFERENCES mst_customer(id),
	tanggal_masuk DATE NOT NULL,
	tanggal_keluar DATE NOT NULL,
	diterima_oleh VARCHAR(100) NOT NULL,
	total_harga INT( default 0
);

CREATE TABLE tx_transaction_enrollment (
	id SERIAL PRIMARY KEY,
	id_transaction INT REFERENCES mst_transaction(id),
	id_service INT REFERENCES mst_service(id),
	jumlah NUMERIC(10, 2) NOT NULL,
	satuan VARCHAR(10) NOT NULL,
	total INT NOT NULL
);