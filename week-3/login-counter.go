package main

import "fmt"

func main() {
	var username, password string
	sumWrongLogin := 0

	fmt.Scan(&username, &password)
	login(username, password, &sumWrongLogin)
	fmt.Println(sumWrongLogin, "Login berhasil")
}

func login(username, password string, sumWrongLogin *int) {
	for username != "admin" || password != "admin" {
		*sumWrongLogin++
		fmt.Scan(&username, &password)
	}
}
