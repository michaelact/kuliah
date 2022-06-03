# Context

Catatan pribadi:
- `basic_context_with_cancel` Ubah initial `counter` value ke 0, saya mendapatkan bahwa jumlah go routine tidak kembali ke awal.
- Gunakan `context.WithTimeout` untuk memberikan signal `cancel()` ketika telah mencapai waktu timeout yang ditentukan
- Gunakan `context.WithDeadline` untuk memberikan signal `cancel()` ketika telah mencapai waktu yang di tentukan
- Pastikan tetap memanggil fungsi `defer cancel()` untuk memastikan tidak terjadi _Go Routine Leak_ ketika proses telah selesai, sebelum waktu yang ditentukan melalui fungsi `.WithTimeout() dan .WithDeadline()`
