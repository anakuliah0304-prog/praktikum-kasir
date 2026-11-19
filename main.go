package main

import (
	"fmt"
)

func main() {
	const (
		syaratDiskon = 100000.0 // Minimal belanja Rp 100.000 untuk dapat diskon
		persenDiskon = 0.10     // Diskon 10%
	)

	var namaBarang string
	var hargaSatuan float64
	var jumlahBarang int
	var uangBayar float64

	fmt.Println("============================================")
	fmt.Println("   PROGRAM KASIR & PERHITUNGAN DISKON GO    ")
	fmt.Println("============================================")

	fmt.Print("Nama Barang                     : ")
	fmt.Scanln(&namaBarang)

	fmt.Print("Harga Satuan (Rp)               : ")
	fmt.Scanln(&hargaSatuan)

	fmt.Print("Jumlah Barang                   : ")
	fmt.Scanln(&jumlahBarang)

	fmt.Print("Nominal Uang Pembeli (Rp)       : ")
	fmt.Scanln(&uangBayar)

	// Perhitungan subtotal dan diskon otomatis
	subtotal := hargaSatuan * float64(jumlahBarang)
	var diskon float64 = 0.0

	if subtotal >= syaratDiskon {
		diskon = subtotal * persenDiskon
	}

	totalBayar := subtotal - diskon
	kembalian := uangBayar - totalBayar

	// Tampilkan struk rincian diskon
	fmt.Println("\n---------------- STRUK BELANJA -------------")
	fmt.Printf("Nama Barang        : %s\n", namaBarang)
	fmt.Printf("Harga Satuan       : Rp %.2f\n", hargaSatuan)
	fmt.Printf("Jumlah Beli        : %d pcs\n", jumlahBarang)
	fmt.Println("--------------------------------------------")
	fmt.Printf("Subtotal           : Rp %.2f\n", subtotal)
	if diskon > 0 {
		fmt.Printf("Diskon (10%%)       : Rp %.2f [DAPAT DISKON 10%%]\n", diskon)
	} else {
		fmt.Println("Diskon (0%)        : Rp 0.00 [BELUM MENCAPAI MIN. RP 100.000]")
	}
	fmt.Println("--------------------------------------------")
	fmt.Printf("Total Bayar        : Rp %.2f\n", totalBayar)
	fmt.Printf("Uang Pembayaran    : Rp %.2f\n", uangBayar)
	fmt.Printf("Uang Kembalian     : Rp %.2f\n", kembalian)
	fmt.Println("============================================")
}