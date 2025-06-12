package main

import "fmt"

const NMAX = 99

type Sampah struct {
	ID        string
	Jenis     string
	Jumlah    int
	DaurUlang string
}

type User struct {
	Username string
	Password string
	Nama     string
	Posisi   string
}

var dataSampah [NMAX]Sampah

var data dataSampah
var nData int

// Fungsi utama program
func main() {
	fmt.Scan()
}

func sequentialSearch(T dataSampah, n int, id string) int {
	var i int
	for i = 0; i < n; i++ {
		if T[i].ID == id {
			return i
		}
	}
	return -1
}

func ubahData(T *dataSampah, n *int) {
	var id string
	var idx int
	var val int
	fmt.Print("Masukkan ID sampah yang akan diubah: ")
	fmt.Scan(&id)
	fmt.Println("Data yang akan diubah :")
	val = sequentialSearch(*T, *n, id)
	fmt.Printf("%s %d %s\n", T[idx].Jenis, T[idx].Jumlah, T[idx].DaurUlang)
	if val == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		idx = val
		fmt.Print("Jenis baru: ")
		fmt.Scan(&T[idx].Jenis)
		fmt.Print("Jumlah baru: ")
		fmt.Scan(&T[idx].Jumlah)
		fmt.Print("Metode daur ulang baru: ")
		fmt.Scan(&T[idx].DaurUlang)
		fmt.Println("Data berhasil diperbarui.")
	}
}

func selectionSortJumlahAsc(T *dataSampah, n int) {
	var i, j, idx int
	var temp Sampah
	for i = 0; i < n-1; i++ {
		idx = i
		j = i + 1
		for j < n {
			if T[j].Jumlah < T[idx].Jumlah {
				idx = j
			}
			j++
		}
		temp = T[i]
		T[i] = T[idx]
		T[idx] = temp
	}
}

// Prosedur selection sort jumlah menurun
func selectionSortJumlahDesc(T *dataSampah, n int) {
	var i, j, idx int
	var temp Sampah
	for i = 0; i < n-1; i++ {
		idx = i
		j = i + 1
		for j < n {
			if T[j].Jumlah > T[idx].Jumlah {
				idx = j
			}
			j++
		}
		temp = T[i]
		T[i] = T[idx]
		T[idx] = temp
	}
}

// Prosedur selection sort jenis menaik
func selectionSortJenisAsc(T *dataSampah, n int) {
	var i, j, idx int
	var temp Sampah
	for i = 0; i < n-1; i++ {
		idx = i
		j = i + 1
		for j < n {
			if T[j].Jenis < T[idx].Jenis {
				idx = j
			}
			j++
		}
		temp = T[i]
		T[i] = T[idx]
		T[idx] = temp
	}
}

// Prosedur selection sort jenis menurun
func selectionSortJenisDesc(T *dataSampah, n int) {
	var i, j, idx int
	var temp Sampah
	for i = 0; i < n-1; i++ {
		idx = i
		j = i + 1
		for j < n {
			if T[j].Jenis > T[idx].Jenis {
				idx = j
			}
			j++
		}
		temp = T[i]
		T[i] = T[idx]
		T[idx] = temp
	}
}

func tampilStatistik(T dataSampah, n int, u User) {
	var total, totalDaurUlang, i int
	var persen int = 0
	for i = 0; i < n; i++ {
		total += T[i].Jumlah
		if T[i].DaurUlang != "-" {
			totalDaurUlang += T[i].Jumlah
		}
	}

	if total > 0 {
		persen = totalDaurUlang * 100 / total
	}

	fmt.Println("=========================================")
	fmt.Println("|            DATA STATISTIK             |")
	fmt.Println("=========================================")
	fmt.Printf("| Nama                   | %-12s |\n", u.Nama)
	fmt.Printf("| Posisi                 | %-12s |\n", u.Posisi)
	fmt.Println("=========================================")
	fmt.Printf("| Total sampah terkumpul | %12d |\n", total)
	fmt.Printf("| Total didaur ulang     | %12d |\n", totalDaurUlang)
	fmt.Printf("| Persentase daur ulang  | %11d%% |\n", persen)
	fmt.Println("=========================================")
}
