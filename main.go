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

func tambahData(T *dataSampah, n *int) {
	var id string
	var idxID int
	fmt.Print("ID sampah: ")
	fmt.Scan(&id)
	idxID = sequentialSearch(*T, *n, id)
	if idxID > -1 {
		for id == T[idxID].ID {
			fmt.Println("ID sudah ada, masukkan ID baru")
			fmt.Scan(&id)
		}
	}
	T[*n].ID = id
	fmt.Print("Jenis sampah: ")
	fmt.Scan(&T[*n].Jenis)
	fmt.Print("Jumlah sampah: ")
	fmt.Scan(&T[*n].Jumlah)
	fmt.Print("Metode daur ulang ('-' jika tidak ada): ")
	fmt.Scan(&T[*n].DaurUlang)
	*n++
	fmt.Println("Data berhasil ditambahkan.")
}

func binarySearch(T dataSampah, n int, id string) int {
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if T[mid].ID == id {
			return mid
		} else if T[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func hapusData(T *dataSampah, n *int) {
	var id string
	var i, idx int
	var val int
	fmt.Print("Masukkan ID sampah yang akan dihapus: ")
	fmt.Scan(&id)
	val = sequentialSearch(*T, *n, id)

	if val != -1 {
		idx = val
		fmt.Println("Data yang akan dihapus :")
		fmt.Printf("%s %d %s\n", T[idx].Jenis, T[idx].Jumlah, T[idx].DaurUlang)
		for i = idx; i < *n-1; i++ {
			T[i] = T[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
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

// Prosedur insertion sort jumlah menaik
func insertionSortJumlahAsc(T *dataSampah, n int) {
	var i, j int
	var key Sampah
	for i = 1; i < n; i++ {
		key = T[i]
		j = i - 1
		for j >= 0 && T[j].Jumlah > key.Jumlah {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

// Prosedur insertion sort jumlah menurun
func insertionSortJumlahDesc(T *dataSampah, n int) {
	var i, j int
	var key Sampah
	for i = 1; i < n; i++ {
		key = T[i]
		j = i - 1
		for j >= 0 && T[j].Jumlah < key.Jumlah {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

// Prosedur insertion sort jenis menaik
func insertionSortJenisAsc(T *dataSampah, n int) {
	var i, j int
	var key Sampah
	for i = 1; i < n; i++ {
		key = T[i]
		j = i - 1
		for j >= 0 && T[j].Jenis > key.Jenis {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

// Prosedur insertion sort jenis menurun
func insertionSortJenisDesc(T *dataSampah, n int) {
	var i, j int
	var key Sampah
	for i = 1; i < n; i++ {
		key = T[i]
		j = i - 1
		for j >= 0 && T[j].Jenis < key.Jenis {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

func cetakData(T dataSampah, n int) {
	var out string = " "
	var i int
	var end bool = false

	fmt.Println("\n===============================================================")
	fmt.Println("|                          DATA SAMPAH                        |")
	fmt.Println("===============================================================")
	fmt.Println("| ID       | Jenis                | Jumlah | Metode Daur      |")
	fmt.Println("===============================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-8s | %-20s | %6d | %-16s |\n", T[i].ID, T[i].Jenis, T[i].Jumlah, T[i].DaurUlang)
	}
	fmt.Println("===============================================================")

	for end == false {
		fmt.Println("kembali ke menu? press x")
		fmt.Scan(&out)
		if out == "x" || out == "X" {
			end = true
		}
	}
}

func menu() {
	fmt.Println("\n============ MENU UTAMA ================")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Ubah Data Sampah")
	fmt.Println("3. Hapus Data Sampah")
	fmt.Println("4. Cari Data Sampah")
	fmt.Println("5. Urutkan Data Sampah")
	fmt.Println("6. Tampilkan Statistik Daur Ulang")
	fmt.Println("7. Cetak Data Daur Ulang")
	fmt.Println("0. Keluar")
	fmt.Println("========================================")
}
