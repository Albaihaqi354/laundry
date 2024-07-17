INSERT INTO mst_customer (nama, no_hp, alamat) 
VALUES
('Jesicca', '0812654987', 'Bandung');

INSERT INTO mst_service (pelayanan, harga) 
VALUES
('Cuci + setirka', 7000.00),
('Laundry BedCover', 50000.00),
('Laundry Boneka', 25000.00);

INSERT INTO mst_transaction(id_customer, tanggal_masuk, tanggal_keluar, diterima_oleh, total_harga)
VALUES
(1, '2022-08-18', '2022-08-20', 'Mirna', 0);

INSERT INTO tx_transaction_enrollment (id_transaction, id_service, jumlah, satuan, total)
VALUES
(1, 1, 5, 'KG', 35000),
(1, 2, 1, 'Buah', 50000),
(1, 3, 2, 'Buah', 50000);