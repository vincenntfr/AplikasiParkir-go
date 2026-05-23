package main

import "fmt"

const MAX int = 100

// TIPE DATA
type Petugas struct {
	usn      string
	password string
	nama     string
}
type Transaksi struct {
	nopol   string
	jenis   string
	durasi  int
	biaya   int
	petugas string
}
type arrPetugas [MAX]Petugas
type arrTransaksi [MAX]Transaksi

//LOGIN
func login(A arrPetugas, n int) bool {
	var user, pass string
	fmt.Print("Username: ")
	fmt.Scan(&user)

	fmt.Print("Password: ")
	fmt.Scan(&pass)

	for i := 0; i < n; i++ {
		if A[i].usn == user && A[i].password == pass {
			fmt.Println("Login Berhasil!")
			return true
		}
	}
	fmt.Println("Login Gagal")
	return false
}

//HITUNG BIAYA
func hitungBiaya(jenis string, durasi int) int {
	var awal, lanjut int
	if jenis == "motor" {
		awal = 3000
		lanjut = 2000
	} else {
		awal = 5000
		lanjut = 3000
	}
	if durasi <= 1 {
		return awal
	}
	return awal + ((durasi - 1) * lanjut)
}

//PETUGAS
func tambahPetugas(A *arrPetugas, n *int) {
	var x Petugas

	fmt.Print("Nama: ")
	fmt.Scan(&x.nama)

	fmt.Print("Username: ")
	fmt.Scan(&x.usn)

	fmt.Print("Password: ")
	fmt.Scan(&x.password)
	A[*n] = x
	*n++

	fmt.Println("Petugas Berhasil Ditambahkan!!")
}

func cariPetugas(A arrPetugas, n int, user string) int {
	for i := 0; i < n; i++ {
		if A[i].usn == user {
			return i
		}
	}
	return -1
}
