package main

import "fmt"

type tDrakor struct {
	judul      string
	rating     float64
	jmlEpisode int
	durasi     int
}

const NMAX = 5

type tabDrakor [NMAX]tDrakor

func main() {
	var drakor tabDrakor

	var nDrakor int

	fmt.Scan(&nDrakor)

	bacaDataDrakor(&drakor, &nDrakor)

	fmt.Println("Data sebelum diurutkan:")
	cetakDataDrakor(drakor, nDrakor)

	urutMenurun(&drakor, nDrakor)
	fmt.Println("Data setelah diurutkan menurun berdasarkan rating:")
	cetakDataDrakor(drakor, nDrakor)

	urutMenaik(&drakor, nDrakor)
	fmt.Println("Data setelah diurutkan menaik berdasarkan durasi:")
	cetakDataDrakor(drakor, nDrakor)
}

func bacaDataDrakor(A *tabDrakor, n *int) {
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].judul)
		fmt.Scan(&A[i].rating)
		fmt.Scan(&A[i].jmlEpisode)
		fmt.Scan(&A[i].durasi)
	}
}

func cetakDataDrakor(A tabDrakor, n int) {
	fmt.Printf("%20s ", "Judul")
	fmt.Printf("%6s ", "Rating")
	fmt.Printf("%6s ", "Jum Ep")
	fmt.Printf("%6s\n", "Durasi")

	for i := 0; i < n; i++ {
		fmt.Printf("%20s ", A[i].judul)
		fmt.Printf("%6.1f ", A[i].rating)
		fmt.Printf("%6d ", A[i].jmlEpisode)
		fmt.Printf("%6d\n", A[i].durasi)
	}
	fmt.Println()
}

func urutMenaik(A *tabDrakor, n int) {
	// insertion sort durasi
	var pass, i int
	var temp tDrakor

	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.durasi < A[i-1].durasi {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func urutMenurun(A *tabDrakor, n int) {
	// selection sort rating
	var pass, idx, i int
	var temp tDrakor

	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].rating < A[i].rating {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
