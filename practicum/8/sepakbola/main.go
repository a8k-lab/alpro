package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol [NMAX]int
	totPertandingan,
	totGol,
	totKebobolan,
	totMenang,
	totKalah,
	totDraw,
	totPoint int
}

func main() {
	var timA, timB tim
	bacaHasil(&timA, &timB)
	fmt.Println("Statistik Tim A")
	cetakHasil(timA)
	fmt.Println("\nStatistik Tim B")
	cetakHasil(timB)
}

func bacaHasil(tA, tB *tim) {
	/*
		IS: Tim A (tA) dan tim B (tB) terdefinisi sembarang
		Proses: Membaca skor pertandingan berupa jumlah gol tim A dan tim B.
				Mengisi field-field sesuai skor kedua tim. Field-fieldnya:
				1. total pertandingan
				2. gol setiap pertandingan
				3. total gol
				4. total kebobolan
				5. total point: point 3 untuk menang, point 1 untuk draw
				6. total menang: Menang, jika gol lebih besar dari gol lawan
				7. total draw: Draw, jika gol sama dengan gol lawan
				8. total kalah: Kalah, jika gol lebih kecil dari gol lawan
				Pembacaan skor berakhir jika kedua skor bernilai negatif.
		FS: Data tim A (tA) dan data tim B (tB) berisi nilai
	*/
	var isLoop bool = true
	var i int = 0
	var golA, golB int

	for isLoop {
		fmt.Scan(&golA, &golB)

		// selama yang di-input bukan negatif
		if golA >= 0 || golB >= 0 {
			// tambahkan data hanya jika jumlah inputan masih di bawah NMAX
			if i < NMAX {
				tA.gol[i] = golA
				tB.gol[i] = golB
				tA.totPertandingan++
				tB.totPertandingan++

				if tA.gol[i] > tB.gol[i] { // jika A menang
					tA.totMenang++
					tB.totKalah++
					tA.totGol += tA.gol[i]
					tB.totGol += tB.gol[i]
					tA.totKebobolan += tB.gol[i]
					tB.totKebobolan += tA.gol[i]
					tA.totPoint += 3
				} else if tB.gol[i] > tA.gol[i] { // jika B menang
					tB.totMenang++
					tA.totKalah++
					tB.totGol += tB.gol[i]
					tA.totGol += tA.gol[i]
					tB.totKebobolan += tA.gol[i]
					tA.totKebobolan += tB.gol[i]
					tB.totPoint += 3
				} else { // jika seri
					tA.totDraw++
					tB.totDraw++
					tA.totGol += tA.gol[i]
					tB.totGol += tB.gol[i]
					tA.totKebobolan += tB.gol[i]
					tB.totKebobolan += tA.gol[i]
					tA.totPoint += 1
					tB.totPoint += 1
				}
			}
			i++
		} else {
			isLoop = false
		}
	}
}

func cetakHasil(t tim) {
	/*
		IS: Data tim (t) terdefinisi
		FS: tercetak di layar statistik pertandingan tim (t) dengan format:

		Total pertandingan: <total pertandingan>
		Gol tiap pertandingan: <gol1 gol2 ... goln>
		Total menang: <total kemenangan>
		Total draw: <total draw>
		Total kalah: <total kalah>
		Total gol: <total gol>
		Total kebobolan: <total kebobolan>
		Total point: <total point>
	*/
	fmt.Printf("Total pertandingan: %d\n", t.totPertandingan)
	fmt.Print("Gol tiap pertandingan: ")
	for i := 0; i < t.totPertandingan; i++ {
		fmt.Printf("%d ", t.gol[i])
	}
	fmt.Printf("\nTotal menang: %d\n", t.totMenang)
	fmt.Printf("Total draw: %d\n", t.totDraw)
	fmt.Printf("Total kalah: %d\n", t.totKalah)
	fmt.Printf("Total gol: %d\n", t.totGol)
	fmt.Printf("Total kebobolan: %d\n", t.totKebobolan)
	fmt.Printf("Total point: %d\n", t.totPoint)
}
