package main

import (
	"fmt"
)

func main() {
	var namaBarang string
	var hargaSatuan float64
	var jumlahBarang int
	var uangBayar float64

	fmt.Println("============================================")
	fmt.Println("        PROGRAM KASIR SEDERHANA GO          ")
	fmt.Println("============================================")

	fmt.Print("Nama Barang                     : ")
	fmt.Scanln(&namaBarang)

	fmt.Print("Harga Satuan (Rp)               : ")
	fmt.Scanln(&hargaSatuan)

	fmt.Print("Jumlah Barang                   : ")
	fmt.Scanln(&jumlahBarang)

	fmt.Print("Nominal Uang Pembeli (Rp)       : ")
	fmt.Scanln(&uangBayar)

	// Perhitungan dasar kasir
	subtotal := hargaSatuan * float64(jumlahBarang)
	totalBayar := subtotal
	kembalian := uangBayar - totalBayar

	// Tampilkan struk sederhana
	fmt.Println("\n---------------- STRUK BELANJA -------------")
	fmt.Printf("Nama Barang        : %s\n", namaBarang)
	fmt.Printf("Harga Satuan       : Rp %.2f\n", hargaSatuan)
	fmt.Printf("Jumlah Beli        : %d pcs\n", jumlahBarang)
	fmt.Println("--------------------------------------------")
	fmt.Printf("Total Belanja      : Rp %.2f\n", totalBayar)
	fmt.Printf("Uang Pembayaran    : Rp %.2f\n", uangBayar)
	fmt.Printf("Uang Kembalian     : Rp %.2f\n", kembalian)
	fmt.Println("============================================")
}