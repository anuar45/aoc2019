package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// Layer is a shortcut type here
type Layer [6][25]int

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	data := []rune(string(input))

	//fmt.Println(len(data))
	var layers []Layer
	var i int
	//fmt.Printf(" %d-%c ", i, v)
	for i < len(data) {
		var layer Layer
		//fmt.Println(len(layer), len(layer[0]))
		for j := 0; j < len(layer); j++ {
			for k := 0; k < len(layer[j]); k++ {
				layer[j][k], _ = strconv.Atoi(string(data[i]))
				i++
			}
		}
		layers = append(layers, layer)
	}

	fmt.Println(layers)

	minZeros := 1<<63 - 1
	var countZeros, countOnes, countTwos int
	var result Layer
	for _, layer := range layers {
		countZeros = 99999
		countZeros = countLayerDigits(layer, 0)
		if countZeros < minZeros {
			minZeros = countZeros
			//fmt.Println(countZeros)
			countOnes = countLayerDigits(layer, 1)
			countTwos = countLayerDigits(layer, 2)
			result = layer
		}
	}

	fmt.Println(result, minZeros)
	fmt.Println(countOnes * countTwos)
}

func countLayerDigits(l Layer, digit int) int {
	var count int
	for _, row := range l {
		for _, v := range row {
			if v == digit {
				count++
			}
		}
	}
	return count

}
