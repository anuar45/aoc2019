package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	data := []rune(string(input))

	const rowsLen = 6
	const digitsLen = 25
	//fmt.Println(len(data))
	var layers [][rowsLen][digitsLen]int
	var i int
	//fmt.Printf(" %d-%c ", i, v)
	for i < len(data) {
		var layer [rowsLen][digitsLen]int
		//fmt.Println(len(layer), len(layer[0]))
		for r := 0; r < rowsLen; r++ {
			for d := 0; d < digitsLen; d++ {
				//fmt.Println(r, d, i)
				layer[r][d], _ = strconv.Atoi(string(data[i]))
				i++
			}
		}
		layers = append(layers, layer)
	}

	fmt.Println(layers)

	//var resultLayer Layer
	var digitSet, resultDigits []int
	var resultDigit int
	for r := 0; r < rowsLen; r++ {
		for d := 0; d < digitsLen; d++ {
			resultDigit = 0
			digitSet = nil
			for l := 0; l < len(layers); l++ {
				digitSet = append(digitSet, layers[l][r][d])
			}
			reverseInts(digitSet)
			resultDigit = renderColor(digitSet)
			resultDigits = append(resultDigits, resultDigit)
		}
	}

	fmt.Println("result:", resultDigits)
	var k int
	for _, d := range resultDigits {
		//fmt.Print(d)
		if d == 1 {
			fmt.Print("*")
		}
		if d == 0 {
			fmt.Print(" ")
		}
		k++
		if k == 25 {
			k = 0
			fmt.Printf("\n")
		}
	}

	// minZeros := 1<<63 - 1
	// var countZeros, countOnes, countTwos int
	// var result Layer
	// for _, layer := range layers {
	// 	countZeros = 99999
	// 	countZeros = countLayerDigits(layer, 0)
	// 	if countZeros < minZeros {
	// 		minZeros = countZeros
	// 		//fmt.Println(countZeros)
	// 		countOnes = countLayerDigits(layer, 1)
	// 		countTwos = countLayerDigits(layer, 2)
	// 		result = layer
	// 	}
	// }

	// fmt.Println(result, minZeros)
	// fmt.Println(countOnes * countTwos)
}

func renderColor(nums []int) int {
	var result int
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			result = 0
		case 1:
			result = 1
		case 2:
			continue
		}
	}
	return result
}

func reverseInts(nums []int) {
	for j, k := 0, len(nums)-1; j < k; {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}
