# Web

## Study
1. Download dependencies: `go mod download`
2. Run test: `go test -v golang_web -run <Test-Function-Name> ./...`

Catatan pribadi:
- Gunakan `net/http/httptest` untuk test handler kita
- Secara default, jika kita tidak menentukan response code, maka response codenya adalah 200
- Gunakan `{{ template "namafile.ext" }}` dan `.ParseGlob("./path/to/templates/*.ext")` untuk import template lain
- Gunakan `{{ define "NAMA_TEMPLATE" }}` untuk memberikan nama pada file template tersebut alih-alih menggunakan nama file
- https://blog.gopheracademy.com/advent-2017/using-go-templates/
- Gunakan `t.Funcs()` untuk menambah fungsi baru pada `text/template`
- Untuk kebutuhan _template caching_, taruh `template.Parse()` menjadi global variable. Hindari melakukan _parsing_ di _handler_ atau berulang-ulang.
- Jika ingin menggunakan template golang untuk kebutuhan web, gunakan `html/template`. Template tersebut secara otomatis mencegah kita dari serangan XSS.
- Tambahkan header `Content-Disposition: attachment; filename="filename.jpg"` untuk langsung mendownload file tanpa melakukan render.
