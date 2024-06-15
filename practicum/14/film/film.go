package main

import "fmt"

type film struct {
	title  string
	year   int
	review float64
	rating float64
}

const NMAX = 10

type tabFilm [NMAX]film

func main() {
	// Deklarasi variabel bertipe array of tipe bentukan struktur
	var data tabFilm

	// Deklarasi variabel bertipe integer untuk menampung banyaknya data film
	var nData int

	// Pemanggilan prosedur baca data film
	bacaDataFilm(&data, &nData)

	// Pemanggilan prosedur cetak data film
	fmt.Println("Data sebelum diurutkan:")
	cetakDataFilm(data, nData)

	// Pemanggilan prosedur urut menurun film berdasarkan review.
	// Jika terdapat film dengan review yang sama, urutkan menurun berdasarkan rating
	urutMenurun(&data, nData)

	// Pemanggilan prosedur cetak data film
	fmt.Println("Data setelah diurutkan menurun berdasarkan review dan rating:")
	cetakDataFilm(data, nData)

	// Pemanggilan prosedur urut menaik film berdasarkan year.
	// Jika terdapat film dengan year yang sama, urutkan menaik berdasarkan review
	urutMenaik(&data, nData)

	// Pemanggilan prosedur cetak data film
	fmt.Println("Data setelah diurutkan menaik berdasarkan year dan review:")
	cetakDataFilm(data, nData)
}

func bacaDataFilm(A *tabFilm, n *int) {
	/*
		IS: Array A sebanyak n terdefinisi sembarang
		Proses: 1) Membaca n. Nilai n tidak boleh melebihi NMAX.
				2) Membaca n data array A dengan atribut-atributnya.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var filmInput film

	// 1.
	fmt.Scan(n)

	// 2.
	for i := 0; i < *n; i++ {
		fmt.Scan(&filmInput.title)
		fmt.Scan(&filmInput.year)
		fmt.Scan(&filmInput.review)
		fmt.Scan(&filmInput.rating)

		if i < NMAX {
			A[i] = filmInput
		}
	}

	if *n > NMAX {
		*n = NMAX
	}
}

func cetakDataFilm(A tabFilm, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen
		FS: Tercetak di layar data array A dengan format
			" Title     Year       Review    Rating
			 <title>   <year>     <review>   <rating>
			  ...	  ...		...			 ...."
			  Gunakan format string dengan kolom 41, 6, 6, dan 6
			  untuk masing-masing title, year, review, dan rating.
	*/
	fmt.Printf("%41s ", "Title")
	fmt.Printf("%6s ", "Year")
	fmt.Printf("%6s ", "Review")
	fmt.Printf("%6s\n", "Rating")

	for i := 0; i < n; i++ {
		fmt.Printf("%41s ", A[i].title)
		fmt.Printf("%6d ", A[i].year)
		fmt.Printf("%6.2f ", A[i].review)
		fmt.Printf("%6.2f\n", A[i].rating)
	}
	fmt.Println()
}

func urutMenaik(A *tabFilm, n int) {
	/*
		IS: Terdefinisi Array A sebanyak n elemen
		FS: Array A sebanyak n elemen terurut menaik berdasarkan
			tahun dengan menggunakan algoritma insertion sort.
			Jika terdapat year yang sama, urutkan menaik berdasarkan
			review-nya.
	*/
	var pass, i int
	var temp film

	// pengurutan untuk year
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.year < A[i-1].year {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}

	// pengurutan untuk review pada year berurutan yang sama
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.year == A[i-1].year && temp.review < A[i-1].review {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func urutMenurun(A *tabFilm, n int) {
	/*
		IS: Terdefinisi Array A sebanyak n elemen
		FS: Array A sebanyak n elemen terurut menurun berdasarkan
			review-nya dengan menggunakan algoritma selection sort.
			Jika terdapat review yang sama, urutkan berdasarkan
			rating-nya.
	*/
	var pass, idx, i int
	var temp film

	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].review < A[i].review {
				idx = i
			} else if A[idx].review == A[i].review {
				// jika secara terurut review memiliki nilai yang sama
				// maka diurutkan berdasarkan rating
				if A[idx].rating < A[i].rating {
					idx = i
				}
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
