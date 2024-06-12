package main

import "fmt"

type pemain struct {
	nama        string
	gol, assist int
}

func main() {
	var pemain1, pemain2, pemain3 pemain
	var ga1, ga2, ga3 int

	fmt.Scan(&pemain1.nama, &pemain1.gol, &pemain1.assist)
	fmt.Scan(&pemain2.nama, &pemain2.gol, &pemain2.assist)
	fmt.Scan(&pemain3.nama, &pemain3.gol, &pemain3.assist)

	ga1 = pemain1.gol + pemain1.assist
	ga2 = pemain2.gol + pemain2.assist
	ga3 = pemain3.gol + pemain3.assist

	if ga1 > ga2 {
		fmt.Println(pemain1.nama)
		if ga2 > ga3 {
			fmt.Println(pemain2.nama)
		} else {
			fmt.Println(pemain3.nama)
		}
	} else if ga2 > ga3 {
		fmt.Println(pemain2.nama)
		if ga3 > ga1 {
			fmt.Println(pemain3.nama)
		} else {
			fmt.Println(pemain1.nama)
		}
	} else {
		fmt.Println(pemain3.nama)
		if ga1 > ga2 {
			fmt.Println(pemain1.nama)
		} else {
			fmt.Println(pemain2.nama)
		}
	}
}
