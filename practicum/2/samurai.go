package main

import "fmt"

func main() {
	var totalSamurai, pwr, spd, enemyPwr, enemySpd int
	fmt.Scan(&totalSamurai)
	fmt.Scan(&pwr, &spd)
	counter := 0

	for i := 0; i < totalSamurai; i++ {
		fmt.Scan(&enemyPwr, &enemySpd)
		if enemyPwr >= pwr && enemySpd >= spd {
			continue
		}

		if enemyPwr >= pwr {
			if enemySpd < spd {
				spd -= 6
				counter++
				continue
			}
		} else {
			if enemySpd > spd {
				pwr -= 6
				counter++
				continue
			}
		}

		if enemyPwr < pwr && enemySpd < spd {
			if pwr > spd {
				pwr -= 6
				counter++
			} else {
				spd -= 6
				counter++
			}
		}
	}

	fmt.Println(counter, pwr, spd)
}
