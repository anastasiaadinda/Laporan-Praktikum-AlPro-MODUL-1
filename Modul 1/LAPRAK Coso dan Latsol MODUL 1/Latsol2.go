package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Scan(&nama, &nim, &kelas)
	fmt.Print("Perkenalkan nama saya adalah " + nama + ", salah satu mahasiswa Prodi S1-IF dari kelas " + kelas + " dengan NIM " + nim + ".")
}
