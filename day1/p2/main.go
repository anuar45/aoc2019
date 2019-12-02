package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var fuelSum int
	f, _ := os.Open("../input.txt")
	scn := bufio.NewScanner(f)

	for scn.Scan() {
		mass, _ := strconv.Atoi(scn.Text())
		fuelSum += calcFuel(mass)
	}

	fmt.Println(fuelSum)
}

func calcFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel > 0 {
		return fuel + calcFuel(fuel)
	}
	return 0
}
