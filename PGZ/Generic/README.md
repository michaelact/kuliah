# Generic

Catatan pribadi:
- _Type Approximation_ bisa digunakan pada awal tipe data, misal `~int`. Sehingga jika kita membuat tipe data `type Age int` saat menggunakan function generic, keduanya bisa dianggap sama.
- _Type Inference_ secara otomatis mendeteksi tipe data yang diberikan. Hanya bisa digunakan pada _function/method generic_.
- https://go.dev/doc/tutorial/generics
- Gunakan package-package berikut untuk beroperasi dengan _slices and maps of generic type_.
  - https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220613132600-b0d781184e0d/slices
  - https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220613132600-b0d781184e0d/maps
