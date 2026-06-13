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
func loginPetugas(A arrPetugas, n int) bool {
	// I.S. :
	// F.S. :
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
func loginAdmin() bool {
	// I.S. :
	// F.S. :
	var user, pass string
	fmt.Print("Username Admin: ")
	fmt.Scan(&user)

	fmt.Print("Password Admin: ")
	fmt.Scan(&pass)

	if user == "admin" && pass == "123" {
		fmt.Println("Login Admin Berhasil!")
		return true
	}

	fmt.Println("Login Admin Gagal")
	return false
}

//HITUNG BIAYA
func hitungBiaya(jenis string, durasi int) int {
	// I.S. :
	// F.S. :
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
	// I.S. :
	// F.S. :
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
	// I.S. :
	// F.S. :
	for i := 0; i < n; i++ {
		if A[i].usn == user {
			return i
		}
	}
	return -1
}
func editPetugas(A *arrPetugas, n int) {
	// I.S. :
	// F.S. :
	var user string
	fmt.Print("Username dicari: ")
	fmt.Scan(&user)
	idx := cariPetugas(*A, n, user)

	if idx == -1 {
		fmt.Println("Petugas tidak ditemukan")
		return
	}

	fmt.Print("Nama Baru: ")
	fmt.Scan(&A[idx].nama)
	fmt.Print("Password Baru: ")
	fmt.Scan(&A[idx].password)

	fmt.Print("Data Berhasil Diubah!")
}

func hapusPetugas(A *arrPetugas, n *int) {
	// I.S. :
	// F.S. :
	var user string

	fmt.Print("Username: ")
	fmt.Scan(&user)

	idx := cariPetugas(*A, *n, user)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for i := idx; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--

	fmt.Println("Petugas Berhasil Dihapus!")
}

// TRANSAKSI
func tambahTransaksi(T *arrTransaksi, n *int) {
	// I.S. :
	// F.S. :
	var x Transaksi
	fmt.Print("Nomor Polisi: ")
	fmt.Scan(&x.nopol)

	fmt.Print("Jenis Kendaraan (motor/mobil): ")
	fmt.Scan(&x.jenis)

	for x.jenis != "motor" && x.jenis != "mobil" {
		fmt.Println("Input Salah!")
		fmt.Print("Jenis: ")
		fmt.Scan(&x.jenis)
	}
	fmt.Print("Durasi Parkir: ")
	fmt.Scan(&x.durasi)

	for x.durasi <= 0 {
		fmt.Println("Durasi harus > 0")
		fmt.Print("Durasi: ")
		fmt.Scan(&x.durasi)
	}
	x.biaya = hitungBiaya(x.jenis, x.durasi)

	fmt.Print("Nama Petugas: ")
	fmt.Scan(&x.petugas)

	T[*n] = x
	*n++
	fmt.Println("Transaksi berhasil ditambah")
}
func cariTransaksi(T arrTransaksi, n int, nopol string) int {
	// I.S. :
	// F.S. :
	for i := 0; i < n; i++ {
		if T[i].nopol == nopol {
			return i
		}
	}
	return -1
}
func editTransaksi(T *arrTransaksi, n int) {
	// I.S. :
	// F.S. :
	var nopol string
	fmt.Print("Masukan nopol: ")
	fmt.Scan(&nopol)

	idx := cariTransaksi(*T, n, nopol)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	fmt.Print("Jenis Baru: ")
	fmt.Scan(&T[idx].jenis)

	fmt.Print("Durasi Baru: ")
	fmt.Scan(&T[idx].durasi)

	T[idx].biaya = hitungBiaya(T[idx].jenis, T[idx].durasi)
	fmt.Println("Data berhasil diubah")
}
func hapusTransaksi(T *arrTransaksi, n *int) {
	// I.S. :
	// F.S. :
	var nopol string

	fmt.Print("Masukan nopol: ")
	fmt.Scan(&nopol)

	idx := cariTransaksi(*T, *n, nopol)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for i := idx; i < *n-1; i++ {
		T[i] = T[i+1]
	}
	*n--
	fmt.Println("Data berhasil dihapus")
}

// SEARCH
func sequentialSearch(T arrTransaksi, n int, nopol string) {
	// I.S. :
	// F.S. :
	found := false
	for i := 0; i < n; i++ {
		if T[i].nopol == nopol {
			fmt.Println("Data ditemukan")
			fmt.Println(T[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

//SORTING BIAYA
func sortingBiayaAsc(T *arrTransaksi, n int) {
	// I.S. :
	// F.S. :
	var pass, idx, i int
	var temp Transaksi

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if T[i].biaya < T[idx].biaya {
				idx = i
			}
		}
		temp = T[pass]
		T[pass] = T[idx]
		T[idx] = temp
	}
}
func sortingBiayaDesc(T *arrTransaksi, n int) {
	// I.S. :
	// F.S. :
	var pass, idx, i int
	var temp Transaksi

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if T[i].biaya > T[idx].biaya {
				idx = i
			}
		}
		temp = T[pass]
		T[pass] = T[idx]
		T[idx] = temp
	}
}

//SORTING NOPOL
func sortingNopolAsc(T *arrTransaksi, n int) {
	// I.S. :
	// F.S. :
	var pass, i int
	var temp Transaksi

	for pass = 1; pass < n; pass++ {
		temp = T[pass]
		i = pass
		for i > 0 && temp.nopol < T[i-1].nopol {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
	}

}
func sortingNopolDesc(T *arrTransaksi, n int) {
	// I.S. :
	// F.S. :
	var pass, i int
	var temp Transaksi

	for pass = 1; pass < n; pass++ {
		temp = T[pass]
		i = pass
		for i > 0 && temp.nopol > T[i-1].nopol {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
	}

}

//REKURSIF
func totalUang(T arrTransaksi, n int) int {
	// I.S. :
	// F.S. :
	if n == 0 {
		return 0
	}
	return T[n-1].biaya + totalUang(T, n-1)
}

//NILAI EKSTREM
func nilaiMax(T arrTransaksi, n int) int {
	// I.S. :
	// F.S. :
	max := T[0].biaya

	for i := 1; i < n; i++ {
		if T[i].biaya > max {
			max = T[i].biaya
		}
	}
	return max
}
func nilaiMin(T arrTransaksi, n int) int {
	// I.S. :
	// F.S. :
	min := T[0].biaya

	for i := 1; i < n; i++ {
		if T[i].biaya < min {
			min = T[i].biaya
		}
	}
	return min
}

// LAPORAN
func laporan(T arrTransaksi, n int) {
	// I.S. :
	// F.S. :
	fmt.Println("===== LAPORAN =====")

	for i := 0; i < n; i++ {
		fmt.Println("Nopol    :", T[i].nopol)
		fmt.Println("Jenis    :", T[i].jenis)
		fmt.Println("Durasi   :", T[i].durasi)
		fmt.Println("Biaya    :", T[i].biaya)
		fmt.Println("Petugas  :", T[i].petugas)
		fmt.Println("=======================")
	}

	if n > 0 {
		fmt.Println("Biaya Maksimum: ", nilaiMax(T, n))
		fmt.Println("Biaya Minimum: ", nilaiMin(T, n))
	}
	fmt.Println("Total Uang: ", totalUang(T, n))
}

//MAIN
func main() {
	var petugas arrPetugas
	var transaksi arrTransaksi
	var nPetugas, nTransaksi, menu int
	var cari string

	petugas[0].usn = "petugas"
	petugas[0].password = "123"
	petugas[0].nama = "Tono"

	nPetugas = 1

	for menu != 3 {

		fmt.Println()
		fmt.Println("===== MENU UTAMA =====")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login Petugas")
		fmt.Println("3. Exit")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		switch menu {

		case 1:
			if loginAdmin() {
				adminMenu := 0
				for adminMenu != 5 {

					fmt.Println()
					fmt.Println("===== MENU ADMIN =====")
					fmt.Println("1. Tambah Petugas")
					fmt.Println("2. Edit Petugas")
					fmt.Println("3. Hapus Petugas")
					fmt.Println("4. Lihat Petugas")
					fmt.Println("5. Logout")

					fmt.Print("Pilih menu: ")
					fmt.Scan(&adminMenu)

					switch adminMenu {
					case 1:
						tambahPetugas(&petugas, &nPetugas)
					case 2:
						editPetugas(&petugas, nPetugas)
					case 3:
						hapusPetugas(&petugas, &nPetugas)
					case 4:
						for i := 0; i < nPetugas; i++ {
							fmt.Println("Nama     :", petugas[i].nama)
							fmt.Println("Username :", petugas[i].usn)
							fmt.Println("======================")
						}
					case 5:
						fmt.Println("Logout admin berhasil")
					}
				}
			}

		case 2:
			if loginPetugas(petugas, nPetugas) {
				petugasMenu := 0
				for petugasMenu != 10 {

					fmt.Println()
					fmt.Println("===== MENU PETUGAS =====")
					fmt.Println("1. Tambah Transaksi")
					fmt.Println("2. Edit Transaksi")
					fmt.Println("3. Hapus Transaksi")
					fmt.Println("4. Cari Kendaraan")
					fmt.Println("5. Sorting Biaya ASC")
					fmt.Println("6. Sorting Biaya DESC")
					fmt.Println("7. Sorting Nopol ASC")
					fmt.Println("8. Sorting Nopol DESC")
					fmt.Println("9. Laporan")
					fmt.Println("10. Logout")

					fmt.Print("Pilih menu: ")
					fmt.Scan(&petugasMenu)

					switch petugasMenu {

					case 1:
						tambahTransaksi(&transaksi, &nTransaksi)
					case 2:
						editTransaksi(&transaksi, nTransaksi)
					case 3:
						hapusTransaksi(&transaksi, &nTransaksi)
					case 4:
						fmt.Print("Cari nopol: ")
						fmt.Scan(&cari)
						sequentialSearch(transaksi, nTransaksi, cari)
					case 5:
						sortingBiayaAsc(&transaksi, nTransaksi)
						fmt.Println("Sorting biaya ASC berhasil")
					case 6:
						sortingBiayaDesc(&transaksi, nTransaksi)
						fmt.Println("Sorting biaya DESC berhasil")
					case 7:
						sortingNopolAsc(&transaksi, nTransaksi)
						fmt.Println("Sorting nopol ASC berhasil")
					case 8:
						sortingNopolDesc(&transaksi, nTransaksi)
						fmt.Println("Sorting nopol DESC berhasil")
					case 9:
						laporan(transaksi, nTransaksi)
					case 10:
						fmt.Println("Logout petugas berhasil")
					}
				}
			}
		case 3:
			fmt.Println("Program selesai")
		}
	}
}
