package main

import "fmt"

const NMAX = 999

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

type dataSampah [NMAX]Sampah

var data dataSampah
var nData int

// Fungsi utama program
func main() {
	var u User
	login(&u)
	var pilihan, idx int
	var end bool = false
	var isAsc bool
	var id string
	var asc string

	for end != true {
		//Fungsi menampilkan menu
		menu()
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1: // Tambah data jika array tidak penuh
			if nData >= NMAX {
				fmt.Print("array sudah penuh")
			} else {
				tambahData(&data, &nData)
			}
		case 2: // Ubah data jika array tidak kosong
			if nData == 0 {
				fmt.Println("Array kosong, tambahkan data terlebih dahulu.")
			} else {
				ubahData(&data, &nData)
			}
		case 3: // Ubah data jika array tidak kosong
			if nData == 0 {
				fmt.Println("Array kosong, tambahkan data terlebih dahulu.")
			} else {
				hapusData(&data, &nData)
			}
		case 4: // Ubah data jika array tidak kosong
			if nData == 0 {
				fmt.Println("Array kosong, tambahkan data terlebih dahulu.")
			} else {
				fmt.Print("Masukkan ID sampah: ")
				fmt.Scan(&id)
				idx = binarySearch(data, nData, id)
				if idx != -1 {
					fmt.Printf("%s %d %s\n", data[idx].Jenis, data[idx].Jumlah, data[idx].DaurUlang)
					fmt.Println("Data ditemukan.")
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			}

		case 5: // Sort array jika tidak kosong
			if nData == 0 {
				fmt.Println("Array kosong, tambahkan data terlebih dahulu.")
			} else {
				var sortType, sortMethod int
				fmt.Println("Pilih Jenis Sort Berdasarkan")
				fmt.Println("1. Jumlah\n2. Jenis")
				fmt.Scan(&sortType)
				fmt.Println("Pilih Metode Sort")
				fmt.Println("1. Selection\n2. Insertion")
				fmt.Scan(&sortMethod)
				fmt.Println("1. Ascending\n2. Descending")
				fmt.Scan(&asc)
				isAsc = asc == "1"
				if sortType == 1 && sortMethod == 1 {
					if isAsc {
						selectionSortJumlahAsc(&data, nData)
					} else {
						selectionSortJumlahDesc(&data, nData)
					}
				} else if sortType == 1 && sortMethod == 2 {
					if isAsc {
						insertionSortJumlahAsc(&data, nData)
					} else {
						insertionSortJumlahDesc(&data, nData)
					}
				} else if sortType == 2 && sortMethod == 1 {
					if isAsc {
						selectionSortJenisAsc(&data, nData)
					} else {
						selectionSortJenisDesc(&data, nData)
					}
				} else if sortType == 2 && sortMethod == 2 {
					if isAsc {
						insertionSortJenisAsc(&data, nData)
					} else {
						insertionSortJenisDesc(&data, nData)
					}
				}
				fmt.Println("Data telah diurutkan.")
			}

		case 6: // Menampilkan statistik kontribusi sampah dari user
			tampilStatistik(data, nData, u)
		case 7: // Mencetak array
			cetakData(data, nData)
		case 0: // end menghentikan loop inputan piranti di menu
			end = true
		default:
			fmt.Println("Input tidak valid.")
		}
	}
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

func hapusData(T *dataSampah, n *int) {
	var id string
	var i, idx int
	var val int
	fmt.Print("Masukkan ID sampah yang akan dihapus: ")
	fmt.Scan(&id)
	val = binarySearch(*T, *n, id)

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

func login(u *User) {
	fmt.Println("=======SELAMAT DATANG DI ECOCYCLE=======")
	fmt.Println("================ LOGIN =================")
	fmt.Print("Username : ")
	fmt.Scan(&u.Username)
	fmt.Print("Password : ")
	fmt.Scan(&u.Password)
	fmt.Print("Nama     : ")
	fmt.Scan(&u.Nama)
	fmt.Print("Posisi   : ")
	fmt.Scan(&u.Posisi)
	fmt.Println("Login berhasil.\n")
	fmt.Println("========================================")
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
