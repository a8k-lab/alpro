package main

import "fmt"

const N int = 256

type tab [N]int

func fillArray(arr *tab, n int) {
	/*I.S. Data tersedia dalam piranti masukan
	  F.S. Array `arr` berisi sebanyak n angka yang dimasukkan user
	*/
	for i := 0; i < n; i++ {
		fmt.Scan(&(*arr)[i])
	}
}

func reverseArray(arr *[]int) {
	/*I.S. Terdefinisi array `arr` berisi sebanyak n angka
	  F.S. Isi array `arr` sudah di-reverse (dibalik isinya)
	*/
	var pointer1, pointer2, temp int

	pointer1 = 0
	pointer2 = len(*arr) - 1

	for pointer1 < pointer2 {
		temp = (*arr)[pointer1]
		(*arr)[pointer1] = (*arr)[pointer2]
		(*arr)[pointer2] = temp

		pointer1++
		pointer2--
	}
}

func isArrayPalindrome(arr []int) bool {
	/* Mengembalikan nilai boolean
	true jika array `arr` adalah palindrom
	false jika sebaliknya
	*/
	var pointer1, pointer2 int

	pointer1 = 0
	pointer2 = len(arr) - 1

	for pointer1 < pointer2 {
		if arr[pointer1] != arr[pointer2] {
			return false
		}

		pointer1++
		pointer2--
	}
	return true
}

func main() {
	var nums tab
	var nums2 = []int{1, 2, 3, 2, 1, 4}
	var nums3 = []int{1, 2, 3, 4, 5, 6, 7}

	fillArray(&nums, 3)
	fmt.Println(nums)

	fmt.Print("Apakah array nums2 adalah palindrom: ")
	fmt.Println(isArrayPalindrome(nums2))

	fmt.Print("Kondisi awal array nums3: ")
	fmt.Println(nums3)
	reverseArray(&nums3)
	fmt.Print("Kondisi array nums3 setelah di-reverse: ")
	fmt.Println(nums3)
}
