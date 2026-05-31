package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Tipe Bentukan (Struct) untuk Bahan Makanan
type BahanMakanan struct {
	Nama        string
	Kategori    string
	Jumlah      int
	Satuan      string
	TglExpired  int // Format: YYYYMMDD
	MinimalStok int
}

// Aturan Array dan Tracker Data
const MAX_BAHAN = 1000
type DaftarBahan [MAX_BAHAN]BahanMakanan

var dataStok DaftarBahan
var jumlahData int = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lanjut := true

	// Dummy Data Awal untuk simulasi pengerjaan awal
	dataStok[0] = BahanMakanan{"Daging Sapi", "Daging", 5, "kg", 20260605, 2}
	dataStok[1] = BahanMakanan{"Wortel", "Sayur", 2, "kg", 20260602, 3}
	jumlahData = 2

	for lanjut {
		fmt.Println("\n=============================================")
		fmt.Println("   APLIKASI MANAJEMEN STOK (PROGRES 50%)    ")
		fmt.Println("=============================================")
		fmt.Println("1. Tambah Data Bahan (Prosedur)")
		fmt.Println("2. Cari Bahan (Sequential & Binary Search)")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu (1-3): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahBahan(scanner)
		case 2:
			menuSearching(scanner)
		case 3:
			fmt.Println("\nProgram dihentikan. Siap dilanjutkan ke tahap sorting.")
			lanjut = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ==================== PROSEDUR INPUT DATA ====================
func tambahBahan(scanner *bufio.Scanner) {
	if jumlahData >= MAX_BAHAN {
		fmt.Println("Gagal: Array penuh!")
		return
	}

	fmt.Println("\n--- INPUT DATA BAHAN ---")
	var b BahanMakanan

	fmt.Print("Nama Bahan: ")
	scanner.Scan()
	b.Nama = strings.TrimSpace(scanner.Text())

	fmt.Print("Kategori: ")
	scanner.Scan()
	b.Kategori = strings.TrimSpace(scanner.Text())

	fmt.Print("Jumlah Stok: ")
	fmt.Scanln(&b.Jumlah)

	fmt.Print("Satuan: ")
	fmt.Scanln(&b.Satuan)

	fmt.Print("Tanggal Expired (YYYYMMDD): ")
	fmt.Scanln(&b.TglExpired)

	fmt.Print("Minimal Stok: ")
	fmt.Scanln(&b.MinimalStok)

	dataStok[jumlahData] = b
	jumlahData++
	fmt.Println("Data berhasil disimpan sementara di array!")
}

// ==================== METODE PENCARIAN (SEARCHING) ====================
func menuSearching(scanner *bufio.Scanner) {
	fmt.Println("\n--- FITUR PENCARIAN ---")
	fmt.Println("1. Sequential Search (Nama/Kategori)")
	fmt.Println("2. Binary Search (Nama Spesifik - Butuh Data Terurut)")
	fmt.Print("Pilih metode (1-2): ")
	var opsi int
	fmt.Scanln(&opsi)

	fmt.Print("Masukkan kata kunci: ")
	scanner.Scan()
	keyword := strings.TrimSpace(scanner.Text())

	if opsi == 1 {
		sequentialSearch(keyword)
	} else if opsi == 2 {
		// Catatan Tubes: Binary search idealnya berjalan setelah proses sorting diimplementasikan
		idx := binarySearch(keyword)
		if idx != -1 {
			fmt.Printf("\n[Ditemukan via Binary Search] Nama: %s | Stok: %d\n", dataStok[idx].Nama, dataStok[idx].Jumlah)
		} else {
			fmt.Println("Data tidak ditemukan atau data belum terurut secara alfabetis.")
		}
	}
}

// Fungsi Sequential Search
func sequentialSearch(keyword string) {
	found := false
	keyLower := strings.ToLower(keyword)

	for i := 0; i < jumlahData; i++ {
		if strings.Contains(strings.ToLower(dataStok[i].Nama), keyLower) {
			fmt.Printf("- %s, Kategori: %s, Stok: %d\n", dataStok[i].Nama, dataStok[i].Kategori, dataStok[i].Jumlah)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Fungsi Binary Search
func binarySearch(keyword string) int {
	low := 0
	high := jumlahData - 1
	keyLower := strings.ToLower(keyword)

	for low <= high {
		mid := (low + high) / 2
		midLower := strings.ToLower(dataStok[mid].Nama)

		if midLower == keyLower {
			return mid
		} else if midLower < keyLower {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
