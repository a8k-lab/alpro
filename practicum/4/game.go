package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	game()
}

func game() {
	var round int = 0
	var computerNumber, yourNumber int
	var win bool = false

	fmt.Println("Start")
	for {
		round = round + 1
		fmt.Println("Round", round)
		numGenerator(&computerNumber)
		yourGuessing(&yourNumber)
		process(yourNumber, computerNumber, &win)
	}
	conclusion(round, win)
	fmt.Println("End")
}

func yourGuessing(yN *int) {
	fmt.Print("Enter your guess: ")
	fmt.Scan(yN)
}

// fungsi numGenerator tidak perlu diubah
func numGenerator(cN *int) {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	*cN = rand.Intn(5)
}

// jika memang yN dan cN dilakukan pengecekan kesamaan value di dalam prosedur `process`, maka seharusnya di-passing value-Nya, bukan pointer-Nya
// mohon koreksinya jika salah
func process(yN, cN int, w *bool) {
	if yN == cN {
		*w = true
		fmt.Printf("Your guessing %d, computer number %d, win %t\n", yN, cN, w)
	}
}

func conclusion(r int, w bool) {
	if r <= 5 && w {
		fmt.Printf("You win in %d round\n", r)
	} else {
		fmt.Printf("Computer win in %d round\n", r)
	}
}
