package main

import "fmt"

const ChessElements int = 64

// Deklarasi tipe alias chessBoard yang merupakan array karakter dengan elemen sebanyak 64
type chessBoard [ChessElements]byte

func main() {
	// Deklarasi variabel board bertipe chessboard
	var board chessBoard

	// Deklarasi variabel status bertipe string
	var status string

	// Deklarasi variabel whitevalue dan blackValue bertipe integer
	var whiteValue, blackValue int

	// Memanggil prosedur readPosition untuk membaca data posisi buah catur
	readPosition(&board)

	// Memanggil prosedur printBoard untuk mencetak data posisi buah catur
	printBoard(board)

	// Memanggil prosedur evalPiece untuk menghitung masing-masing jumlah nilai
	// buah catur putih (whiteValue), jumlah nilai buah catur hitam (blackValue),
	// dan status keunggulan
	evalPiece(board, &whiteValue, &blackValue, &status)

	// Mencetak nilai buah catur putih (whiteValue) dan buah catur hitam (blackValue)
	fmt.Printf("\n%d %d %s\n", whiteValue, blackValue, status)
}

func readPosition(A *chessBoard) {
	/*
		IS : Array A bertipe chessBoard sebanyak 64 elemen terdefinisi sembarang
		Proses : Membaca karakter 'K', 'k', 'Q', 'q', 'R', 'r', 'B', 'b', 'N', 'n', 'P', 'p', dan '.' dan memasukkannya
				 ke dalam array yang terdiri dari 64 elemen.
		FS: Array A berisi nilai
	*/
	for i := 0; i < ChessElements; i++ {
		fmt.Scanf("%c", &A[i])
	}
}

func printBoard(A chessBoard) {
	/*
		IS: Array A bertipe chessBoard sebanyak 64 elemen terdefinisi
		FS: Tercetak elemen array A di layar dengan format 8 x 8 petak (lihat contoh)
	*/
	for i := 0; i < ChessElements; i++ {
		if i != 0 && i%8 == 0 {
			// cetak enter setiap 8 elemen
			fmt.Println()
		} else {
			fmt.Printf("%c", A[i])
		}
	}
}

func evalPiece(A chessBoard, wV, bV *int, status *string) {
	/*
		IS: Array A bertipe chessBoard sebanyak 64 elemen terdefinisi
		FS: Nilai buah catur putih (wV) dan nilai buah catur hitam (bV) berisi nilai.
			Buah catur putih dan nilainya: K = 200, Q = 9, R = 5, B = 3, N = 3, P = 1.
			Buah catur hitam dan nilainya: k = 200, q = 9, r = 5, b = 3, n = 3, p = 1.
			Status berisi nilai berdasarkan nilai buah catur kedua pemain
			("unggul putih", "unggul hitam", "imbang")
	*/
	for i := 0; i < ChessElements; i++ {
		// pengkondisian penambahan value untuk putih
		if A[i] == 'K' {
			*wV += 200
		} else if A[i] == 'Q' {
			*wV += 9
		} else if A[i] == 'R' {
			*wV += 5
		} else if A[i] == 'B' {
			*wV += 3
		} else if A[i] == 'N' {
			*wV += 3
		} else if A[i] == 'P' {
			*wV += 1
		}

		// pengkondisian penambahan value untuk hitam
		if A[i] == 'k' {
			*bV += 200
		} else if A[i] == 'q' {
			*bV += 9
		} else if A[i] == 'r' {
			*bV += 5
		} else if A[i] == 'b' {
			*bV += 3
		} else if A[i] == 'n' {
			*bV += 3
		} else if A[i] == 'p' {
			*bV += 1
		}
	}

	if *wV > *bV {
		*status = "unggul putih"
	} else if *bV > *wV {
		*status = "unggul hitam"
	} else {
		*status = "imbang"
	}
}
