package main

import "fmt"

func main() {
	var lumin int
	fmt.Scan(&lumin)

	stellaris := lumin / 20
	nebulaCrown := stellaris / 2
	galactum := nebulaCrown / 3
	quantumShard := galactum / 10

	lumin = lumin % 20
	stellaris = stellaris % 2
	nebulaCrown = nebulaCrown % 3
	galactum = galactum % 10

	fmt.Println(quantumShard, galactum, nebulaCrown, stellaris, lumin)
}
