package main

import "fmt"
const MAX = 10

type Interaksi struct {
	waktu   string
	pesan   string
	emosi   string
	saran   string
	urgensi int
}

var dataInteraksi [MAX]Interaksi
var jumlahInteraksi int

func main() {
	var pilihan int
	var input string
	var repeat bool
	repeat = true
	for repeat {
		fmt.Println("\n=== Halo, Aku Siap Membantu Kamu ;) ===\n")
		fmt.Println("CATATAN:\n1. Untuk mengupdate curhatanmu masukan pesan 1 kata/kalimat pendek dan untuk jarak kalimat kamu gunakan (-) ya!\n2. Untuk menu 5 dan 6 kamu bisa pilih menu 2 untuk melihat hasil update-nya ya!\n")
		fmt.Println("1. Tambah Curhatan Mu")
		fmt.Println("2. Tampilkan Semua Curhatan Mu")
		fmt.Println("3. Cari Curhatan Mu Berdasarkan Emosi (Sequential Search)")
		fmt.Println("4. Cari Berdasarkan Waktu (Binary Search)")
		fmt.Println("5. Urutkan Berdasarkan Waktu (Selection Sort)")
		fmt.Println("6. Urutkan Berdasarkan tingkat urgensi (Insertion Sort)")
		fmt.Println("7. Ubah Curhatan Mu")
		fmt.Println("8. Hapus Curhatan Mu")
		fmt.Println("0. Keluar")
		fmt.Print("\nKamu boleh pilih menu-nya (0-8): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahInteraksi(&dataInteraksi, &jumlahInteraksi)
		case 2:
			tampilkanInteraksi(dataInteraksi, jumlahInteraksi)
		case 3:
			fmt.Print("Masukkan emosi yang kamu cari: ")
			fmt.Scan(&input)
			sequentialSearchEmosi(dataInteraksi, jumlahInteraksi, input)
		case 4:
			fmt.Print("Masukkan waktu yang kamu cari (HH-BB-TTTT/JJ:MM): ")
			fmt.Scan(&input)
			binarySearchWaktu(dataInteraksi, jumlahInteraksi, input)
		case 5:
			selectionSortWaktu(&dataInteraksi, jumlahInteraksi)
		case 6:
			insertionSortUrgensi(&dataInteraksi, jumlahInteraksi)
		case 7:
			ubahInteraksi(&dataInteraksi, jumlahInteraksi)
		case 8:
			hapusInteraksi(&dataInteraksi, &jumlahInteraksi)
		case 0:
			fmt.Println("\n=== Terima kasih telah curhat dengan ku. Semoga harimu membaik ya ;) ===")
			repeat = false
		default:
			fmt.Println("== Pilihan tidak valid. Silakan coba lagi. ==")
		}
	}
}

func tambahInteraksi(data *[MAX]Interaksi, jumlah *int) {
    var waktu, pesan, emosi, saran string
    var urgensi int
	
    if *jumlah >= MAX {
        fmt.Println("\n== Data interaksi sudah penuh. ==")
    } else {
        fmt.Print("\nWaktu (HH-BB-TTTT/JJ:MM): ")
        fmt.Scan(&waktu)
        fmt.Print("Apa yang kamu ingin curhatkan?: ")
        fmt.Scan(&pesan)
        fmt.Print("Apa emosi kamu saat ini (sedih/cemas/lelah/marah): ")
        fmt.Scan(&emosi)
        fmt.Print("Pilih tingkat urgensimu saat ini (1-10): ")
        fmt.Scan(&urgensi)

        saran = rekomendasiSaran(emosi)
        data[*jumlah] = Interaksi{waktu, pesan, emosi, saran, urgensi}
        *jumlah++
        fmt.Println("\n== Interaksi kamu berhasil ditambahkan! ==")
    }
}

func rekomendasiSaran(emosi string) string {
	switch emosi {
	case "sedih":
		return "Kamu bisa mendengarkan musik atau menonton film :)."
	case "cemas":
		return "Tarik napas dalam dan fokus pada pernapasan lalu kamu bisa melakulan aktivitas untuk mengurangi kecemasan :)."
	case "lelah":
		return "Beristirahatlah sejenak/cukup minum air putih yang cukup dan prioritaskan sesuai kebutuhan diri ya :)."
	case "marah":
		return "Tenangkan diri kamu dengan cara cuci muka kamu dengan air dingin untuk merelaksasi dan evalusi terhadap diri kamu :)."
	default:
		return "Tetap semangat dan jaga kesehatanmu!"
	}
}

func tampilkanInteraksi(data [MAX]Interaksi, jumlah int) {
    var i int
    if jumlah == 0 {
        fmt.Println("\n== Belum ada curhatanmu yang tersimpan. ==")
    } else {
        fmt.Println("\n========================================================================================================================================================================================================")
        fmt.Printf("| %-2s | %-30s | %-15s | %-10s | %-7s | %-117s |\n", "No", "Waktu", "Pesan", "Emosi", "Urgensi", "Saran")
        fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

        for i = 0; i < jumlah; i++ {
            fmt.Printf("| %-2d | %-30s | %-15s | %-10s | %-7d | %-117s |\n", i+1,
                data[i].waktu, data[i].pesan, data[i].emosi, data[i].urgensi, data[i].saran)
        }
        fmt.Println("========================================================================================================================================================================================================")
    }
}

func sequentialSearchEmosi(data [MAX]Interaksi, jumlah int, emosiCari string) {
	var i int
	var ditemukan bool
	
	ditemukan = false
	fmt.Println("\nHasil pencarian emosi:", emosiCari)
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-2s | %-30s | %-15s | %-10s | %-7s | %-117s |\n", "No", "Waktu", "Pesan", "Emosi", "Urgensi", "Saran")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	for i = 0; i < jumlah; i++ {
		if data[i].emosi == emosiCari {
			fmt.Printf("| %-2d | %-30s | %-15s | %-10s | %-7d | %-117s |\n", i+1, data[i].waktu, data[i].pesan, data[i].emosi, data[i].urgensi, data[i].saran)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("| Tidak ditemukan curhatan dengan emosi tersebut.                                                                                                                                                      |")
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
}
func binarySearchWaktu(data [MAX]Interaksi, jumlah int, waktuCari string) {
	var i, low, high, mid int
	var kiri, kanan int
	var found, repeat bool
	
	low = 0
	high = jumlah - 1
	found = false
	repeat = true

	for repeat {
		if low > high {
			repeat = false
		} else {
			mid = (low + high) / 2
			if data[mid].waktu == waktuCari {
				found = true
				repeat = false
			} else if data[mid].waktu < waktuCari {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	if found {
		kiri = mid
		for kiri > 0 && data[kiri-1].waktu == waktuCari {
			kiri--
		}
		kanan = mid
		for kanan < jumlah-1 && data[kanan+1].waktu == waktuCari {
			kanan++
		}

		fmt.Println("\nDitemukan curhatanmu dengan waktu:", waktuCari)
		fmt.Println("========================================================================================================================================================================================================")
		fmt.Printf("| %-2s | %-30s | %-15s | %-10s | %-7s | %-117s |\n", "No", "Waktu", "Pesan", "Emosi", "Urgensi", "Saran")
		fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

		for i = kiri; i <= kanan; i++ {
			fmt.Printf("| %-2d | %-30s | %-15s | %-10s | %-7d | %-117s |\n", i+1, data[i].waktu, data[i].pesan, data[i].emosi, data[i].urgensi, data[i].saran)
		}

		fmt.Println("========================================================================================================================================================================================================")
	} else {
		fmt.Println("\n== Data tidak ditemukan. ==")
	}
}


func selectionSortWaktu(data *[MAX]Interaksi, jumlah int) {
	var i, j, minIdx int
	var temp Interaksi
	for i = 0; i < jumlah-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlah; j++ {
			if data[j].waktu < data[minIdx].waktu {
				minIdx = j
			}
		}
		if minIdx != i {
			temp = data[i]
			data[i] = data[minIdx]
			data[minIdx] = temp
		}
	}
	fmt.Println("\n== Data berhasil diurutkan berdasarkan waktu (selection sort). ==")
}

func insertionSortUrgensi(data *[MAX]Interaksi, jumlah int) {
	var i, j int
	var temp Interaksi
	for i = 1; i < jumlah; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].urgensi < temp.urgensi {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
	fmt.Println("\n== Data berhasil diurutkan berdasarkan tingkat urgensi (insertion sort). ==")
}


func ubahInteraksi(data *[MAX]Interaksi, jumlah int) {
    var idx int
    var waktu string
    
    fmt.Print("Masukkan nomor curhat yang ingin diubah: ")
    fmt.Scan(&idx)

    if idx >= 1 && idx <= jumlah {
        idx--
        
        fmt.Print("Waktu baru (HH-BB-TTTT/JJ:MM): ")
        fmt.Scan(&waktu)
        if waktu != "" {
            data[idx].waktu = waktu
        }
        
        fmt.Print("Pesan baru: ")
        fmt.Scan(&data[idx].pesan)
        fmt.Print("Emosi baru (sedih/cemas/lelah/marah): ")
        fmt.Scan(&data[idx].emosi)
        fmt.Print("Tingkat urgensi baru (1-10): ")
        fmt.Scan(&data[idx].urgensi)
        data[idx].saran = rekomendasiSaran(data[idx].emosi)
        fmt.Println("\n== Interaksi berhasil diubah. ==")
    } else {
        fmt.Println("\n== Nomor tidak valid. ==")
    }
}

func hapusInteraksi(data *[MAX]Interaksi, jumlah *int) {
    var idx int
    fmt.Print("Masukkan nomor curhat yang ingin dihapus: ")
    fmt.Scan(&idx)

    if idx >= 1 && idx <= *jumlah {
        idx--
        for i := idx; i < *jumlah-1; i++ {
            data[i] = data[i+1]
        }
        *jumlah--
        fmt.Println("\n== Interaksi berhasil dihapus. ==")
    } else {
        fmt.Println("\n== Nomor tidak valid. ==")
    }
}


	



