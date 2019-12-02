package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var masses []int
	var fuelSum int
	f, _ := os.Open("../input.txt")

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		mass, _ := strconv.Atoi(scn.Text())
		masses = append(masses, mass)
	}

	for _, mass := range masses {
		fuel := mass/3 - 2
		fuelSum += fuel
	}

	//fmt.Println(fuelSum)
	fmt.Println(fuelSum)
}
