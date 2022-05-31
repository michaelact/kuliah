# Go routines

Catatan pribadi:
- `sync.Once` memastikan hanya go routine pertama yang bisa mengeksekusi _function_ tersebut
- Gunakan `sync.Pool` untuk memastikan data yang ada di pool aman dari _race condition_
- Gunakan `sync.Map` dibanding map biasa untuk memastikan aman dari _race condition_
- Gunakan `sync/atomic` untuk menggunakan data primitif secara aman dalam proses _concurrent_
