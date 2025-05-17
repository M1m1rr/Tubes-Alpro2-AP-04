package main

import "fmt"

const NMAX = 99

type Sampah struct {
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

// Fungsi utama program
func main() {
	fmt.Scan()
}
