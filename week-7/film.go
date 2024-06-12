package main

import "fmt"

const Kapasitas int = 10

type Film struct {
	Judul  string
	Rating float64
}

type tabFilm [Kapasitas]Film

func main() {
	var listFilm tabFilm
	var input Film

	//  buat perulangan untuk memasukkan judul film dan rating,
	for i := 0; i < Kapasitas; i++ {
		fmt.Scan(&input.Judul)
		if input.Judul == "#" {
			break
		}
		fmt.Scan(&input.Rating)

		//  jika rating lebih besar dari 4.00 maka masukkan data ke array listFilm
		if input.Rating > 4.00 {
			listFilm[i] = input
		}
	}

	//  menampilkan judul film dengan rating lebih dari 4.00
	for _, film := range listFilm {
		fmt.Printf("%s ", film.Judul)
	}
	fmt.Println()
}
