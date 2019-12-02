package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var fuelSum int
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel := mass/3 - 2

		for fuel > 0 {
			mass = fuel
			fuelSum += fuel
		}
		fuelSum += fuel
	}

	fmt.Println(fuelSum)
}

func calcFuel(mass int) {
	fuel := mass/3 - 2
	if fuel > 0 {
		+calcFuel(fuel)
	}
	return
}
