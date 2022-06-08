# Embed

Setup Database:
1. Install Docker dan Docker Compose
2. `docker-compose up -d`
3. Default attribute value:
   - User: `root`
   - Password: `michaelact`
   - Database: `testing`

Catatan pribadi:
- Taruh komentar `//go:embed` diatas variabel kosong yang akan menampungnya.
- File yang di _embed_ tidak akan berubah lagi ketika sudah di compile.
