# Database

Setup Database:
1. Install Docker dan Docker Compose
2. `docker-compose up -d`
3. Default attribute value:
   - User: `root`
   - Password: `michaelact`
   - Database: `testing`

Catatan:
- `db.SetMaxIdleConns` > Minimal koneksi yang tersedia walaupun sedang tidak digunakan
- `db.SetMaxOpenConns` > Maksimal koneksi yang bisa dibuat ketika permintaan meningkat
- `db.SetConnMaxIdle` > Waktu tunggu yang diberikan untuk menghapus koneksi yang tidak digunakan sampai batas minimum yang ditentukan
- `db.SetConnMaxLifetime` > Batas waktu untuk menggunakan koneksi
- `sql.Null*` Tipe data dalam package `database/sql` untuk data yang bisa memiliki NULL
- Gunakan `?` untuk menerima input dari user dan membungkusnya dengan query disertai parameter sebagai salah satu pencegahan SQL Injection.
- Gunakan `LastInsertId` untuk mendapatkan ID _AutoIncrement_ terakhir
- [Prepare Statement](https://medium.com/pujanggateknologi/prepared-statement-di-go-927b1a8863ec)
