package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		syaratDiskon = 100000.0 // Minimal belanja Rp 100.000
		persenDiskon = 0.10     // Diskon 10%
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("====================================================")
	fmt.Println("        PROGRAM KASIR MINI & PERHITUNGAN DISKON     ")
	fmt.Println("          DENGAN VALIDASI KECUKUPAN PEMBAYARAN      ")
	fmt.Println("====================================================")

	// 1. Input dan Validasi Nama Barang
	fmt.Print("Masukkan Nama Barang              : ")
	namaBarang, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] Gagal membaca input nama barang.")
		return
	}
	namaBarang = strings.TrimSpace(namaBarang)
	if namaBarang == "" {
		fmt.Println("\n[VALIDASI GAGAL] Nama barang tidak boleh kosong!")
		return
	}

	// 2. Input dan Validasi Harga Satuan
	fmt.Print("Masukkan Harga Satuan (Rp)        : ")
	inputHarga, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] Gagal membaca input harga satuan.")
		return
	}
	hargaSatuan, err := strconv.ParseFloat(strings.TrimSpace(inputHarga), 64)
	if err != nil || hargaSatuan <= 0 {
		fmt.Println("\n[VALIDASI GAGAL] Harga satuan harus berupa angka dan lebih besar dari Rp 0!")
		return
	}

	// 3. Input dan Validasi Jumlah Barang
	fmt.Print("Masukkan Jumlah Barang            : ")
	inputJumlah, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] Gagal membaca input jumlah barang.")
		return
	}
	jumlahBarang, err := strconv.Atoi(strings.TrimSpace(inputJumlah))
	if err != nil || jumlahBarang <= 0 {
		fmt.Println("\n[VALIDASI GAGAL] Jumlah barang harus berupa bilangan bulat positif lebih dari 0!")
		return
	}

	// 4. Input dan Validasi Uang Pembayaran
	fmt.Print("Masukkan Nominal Uang Pembeli (Rp): ")
	inputUang, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] Gagal membaca input uang pembeli.")
		return
	}
	uangBayar, err := strconv.ParseFloat(strings.TrimSpace(inputUang), 64)
	if err != nil || uangBayar < 0 {
		fmt.Println("\n[VALIDASI GAGAL] Nominal uang pembayaran harus berupa angka valid!")
		return
	}

	// Perhitungan Subtotal dan Diskon Otomatis
	subtotal := hargaSatuan * float64(jumlahBarang)
	var diskon float64 = 0.0

	if subtotal >= syaratDiskon {
		diskon = subtotal * persenDiskon
	}

	totalBayar := subtotal - diskon

	// Tampilkan Struk Pembelian
	fmt.Println("\n------------------- STRUK PEMBELIAN -----------------")
	fmt.Printf("Nama Barang        : %s\n", namaBarang)
	fmt.Printf("Harga Satuan       : Rp %15.2f\n", hargaSatuan)
	fmt.Printf("Jumlah Beli        : %15d pcs\n", jumlahBarang)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Subtotal           : Rp %15.2f\n", subtotal)
	if diskon > 0 {
		fmt.Printf("Diskon (10%%)       : Rp %15.2f [HEMAT 10%%]\n", diskon)
	} else {
		fmt.Printf("Diskon (0%%)        : Rp %15.2f [TIDAK DAPAT DISKON]\n", 0.0)
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Total Harus Bayar  : Rp %15.2f\n", totalBayar)
	fmt.Printf("Uang Diterima      : Rp %15.2f\n", uangBayar)

	// Fitur Pengembangan: Validasi Kecukupan Uang Pembayaran
	if uangBayar < totalBayar {
		kekurangan := totalBayar - uangBayar
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("STATUS TRANSAKSI   : [PEMBAYARAN KURANG / GAGAL]\n")
		fmt.Printf("Kekurangan Uang    : Rp %15.2f\n", kekurangan)
		fmt.Println("[PERINGATAN] Uang pembayaran tidak mencukupi untuk melunasi belanjaan!")
	} else {
		kembalian := uangBayar - totalBayar
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("STATUS TRANSAKSI   : [LUNAS]\n")
		fmt.Printf("Uang Kembalian     : Rp %15.2f\n", kembalian)
		fmt.Println("  Terima kasih telah berbelanja di Kasir Kami!")
	}
	fmt.Println("=====================================================")
}