package main

import (
	"fmt"
	"os"
	"strconv"
)

// Teman struct untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// DataTeman sebagai database sementara
var DataTeman = map[int]Teman{
	1: {"Budi", "Jakarta", "Nganggur", "seru"},
	2: {"Yoga", "Bandung", "Serabutan", "Penuh tantangan"},
	3: {"Jambrut", "Surabaya", "Petani", "asyik"},
}

// TampilkanDataTeman untuk menampilkan data teman berdasarkan nomor absen
func TampilkanDataTeman(absen int) {
	teman, found := DataTeman[absen]
	if !found {
		fmt.Println("Teman dengan absen nomor", absen, "tidak ditemukan.")
		return
	}

	fmt.Println("Untuk Absen Nomor", absen)
	fmt.Println("● Nama =", teman.Nama)
	fmt.Println("● Alamat =", teman.Alamat)
	fmt.Println("● Pekerjaan =", teman.Pekerjaan)
	fmt.Println("● Alasan memilih kelas Golang =", teman.Alasan)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <nomor_absen>")
		os.Exit(1)
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		os.Exit(1)
	}

	TampilkanDataTeman(absen)
}
