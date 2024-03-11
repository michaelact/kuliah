#include <iostream>
#include <unordered_map>
#include <string>

int main() {
    std::unordered_map<std::string, int> PeopleAid = {
        {"Bayi", 1000000},
        {"Batita", 1000000},
        {"Balita", 1000000},
        {"Anak", 1000000},
        {"Remaja", 1500000},
        {"Pemuda", 1500000},
        {"Dewasa", 1500000},
        {"Lansia", 2000000}
    };

    // Membuat variable untuk menyimpan umur dan kategori
    int age;
    std::string people;

    // Meminta pengguna memasukkan umur dia
    std::cout << "Masukkan umur anda: ";
    std::cin >> age;

    // Menentukan kategori berdasarkan umur
    if (age < 0) {
        // Menghentikan program jika umur tidak sesuai dengan kategori yang ditentukan.
        std::cout << "Kamu tidak mendapatkan apa-apa." << std::endl;
        return 0;
    } else if (age <= 1) {
        people = "Bayi";
    } else if (age <= 3) {
        people = "Batita";
    } else if (age <= 5) {
        people = "Balita";
    } else if (age <= 12) {
        people = "Anak";
    } else if (age <= 17) {
        people = "Remaja";
    } else if (age <= 30) {
        people = "Pemuda";
    } else if (age <= 60) {
        people = "Dewasa";
    } else if (age > 60) {
        people = "Lansia";
    }

    // Menampilkan bantuan dana yang didapatkan
    std::cout << "Kamu mendapatkan " << PeopleAid[people] << std::endl;
    return 0;
}
