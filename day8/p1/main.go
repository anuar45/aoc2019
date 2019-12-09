package main

import (
	"io/ioutil"
)

type Layer [6][24]string

type Layers []Layer

func main() {
	input, _ := ioutil.ReadFile("../input.txt")

	layers := []Layer{}

	for i, v := range string(input) {
		var layer Layer

		for j := 0; j < len(layer); j++ {

		}

	}

}
