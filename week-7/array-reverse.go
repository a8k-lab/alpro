// package main

// import "fmt"

// const N int = 100

// type tab [N]rune

// func masukanArray(T *tab, n *int) {
// 	/*I.S. Data tersedia dalam piranti masukan
// 	  F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
// 	  Proses input selama karakter bukanlah TITIK dan n <= N
// 	  Petunjuk: Gunakan fmt.Scanf untuk melakukan input karakter
// 	*/
// 	for i := 0; i < N; i++ {
// 		fmt.Scanf("%c", &(*T)[i])
// 		if (*T)[i] == '.' {
// 			break
// 		}
// 		*n++
// 	}
// }

// func reverseArray(T *tab, n int) {
// 	/*I.S. Terdefinisi array T yang berisikan n jumlah karakter
// 	  F.S. Isi array sudah direverse (urutan isi array T terbalik)
// 	*/
// 	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
// 		(*T)[i], (*T)[j] = (*T)[j], (*T)[i]
// 	}
// }

// func cetakArray(T tab, n int) {
// 	/*{I.S. Terdefinisi sejumlah n karakter di dalam array T.
// 	  F.S. Menampilkan isi setiap karakter pada layar
// 	  Petunjuk: Untuk melakukan print pada karakter, gunakan fmt.Printf}
// 	*/
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("%c", T[i])
// 	}
// 	fmt.Println()
// }

// func main() {
// 	var T tab
// 	var n int
// 	masukanArray(&T, &n)
// 	reverseArray(&T, n)
// 	cetakArray(T, n)
// }
