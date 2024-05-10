package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0

	bacaData(&timnas, &nPemain)
	cetakPemain(timnas, nPemain)
	cetakNamaPemainTertinggi(timnas, nPemain)
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	var isLoop bool = true
	var firstInput string
	var i int = 0

	for isLoop {
		if *n < NMAX {
			fmt.Scan(&firstInput)

			if firstInput == "none" {
				isLoop = false
			} else {
				(*A)[i].nama = firstInput
				fmt.Scan(&(*A)[i].nomorPunggung)
				fmt.Scan(&(*A)[i].posisi)
				fmt.Scan(&(*A)[i].tinggi)

				*n++
				i++
			}
		} else {
			isLoop = false
		}
	}
}

func cetakPemain(A tabPemain, n int) {
	fmt.Println("Data Pemain:")

	for i := 0; i < n; i++ {
		fmt.Printf("%s ", A[i].nama)
		fmt.Printf("%s ", A[i].nomorPunggung)
		fmt.Printf("%s ", A[i].posisi)
		fmt.Printf("%d\n", A[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	var pemainTertinggi pemain = A[0]
	for i := 0; i < n; i++ {
		if A[i].tinggi > pemainTertinggi.tinggi {
			pemainTertinggi = A[i]
		}
	}

	fmt.Printf("Pemain dengan badan tertinggi: %s\n", pemainTertinggi.nama)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	var pemainTerendah pemain = A[0]
	for i := 0; i < n; i++ {
		if A[i].tinggi < pemainTerendah.tinggi {
			pemainTerendah = A[i]
		}
	}

	fmt.Printf("Pemain dengan badan terendah: %s\n", pemainTerendah.nama)
}
