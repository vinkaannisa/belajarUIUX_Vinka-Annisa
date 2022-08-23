package main

import (
	"fmt"
	"Alterra/hitung"
	"Alterra/minus"
)

func main() {
	fmt.Println("Saya vinka Annisa")
	fmt.Println("fiturA point 4")

	hasilTambah := hitung.Tambah(9, 9)
	fmt.Println(hasilTambah)

	fmt.Println("featureB point 5")
	hasilKali := hitung.Kali(9, 9)
	fmt.Println(hasilKali)

	hasilKurang := minus.Kurang(9, 9)
	fmt.Println(hasilKurang)
}