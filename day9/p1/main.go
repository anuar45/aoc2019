package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Computer struct to track state of IntCode Computer
type Computer struct {
	I    int
	Code map[int]int
	Base int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	inputSlice := strings.Split(string(input), ",")

	nums := make(map[int]int)
	for i, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		nums[i] = num
		//nums = append(nums, num)
	}

	cmp := Computer{0, nums, 0}

	out := cmp.intCodeComp(1)

	fmt.Println(out)
}

func (cmp *Computer) intCodeComp(signal int) []int {
	var output []int
	var i int
	nums := cmp.Code

	for i = cmp.I; i < len(nums); {
		opCode := intToSlice(nums[i])
		fmt.Println(opCode, cmp.Base, i)
		fmt.Println(nums)
		if nums[i] == 99 {
			fmt.Println("terminated")
			break
		}

		reverseInts(opCode)
		for len(opCode) < 4 {
			opCode = append(opCode, 0)
		}
		reverseInts(opCode)

		switch opCode[len(opCode)-1] {
		case 1:
			switch opCode[len(opCode)-3] {
			case 0:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[nums[i+1]] + nums[i+2]
					i += 4
				case 2:
					nums[nums[i+3]] = nums[nums[i+1]] + nums[cmp.Base+nums[i+2]]
					i += 4
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[i+1] + nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[i+1] + nums[i+2]
					i += 4
				case 2:
					nums[nums[i+3]] = nums[i+1] + nums[cmp.Base+nums[i+2]]
					i += 4
				}
			case 2:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[cmp.Base+nums[i+1]] + nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[cmp.Base+nums[i+1]] + nums[i+2]
					i += 4
				case 2:
					nums[nums[i+3]] = nums[cmp.Base+nums[i+1]] + nums[cmp.Base+nums[i+2]]
					i += 4
				}
			}
		case 2:
			switch opCode[len(opCode)-3] {
			case 0:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[nums[i+1]] * nums[i+2]
					i += 4
				case 2:
					nums[nums[i+3]] = nums[nums[i+1]] * nums[cmp.Base+nums[i+2]]
					i += 4
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[i+1] * nums[nums[i+2]]
					i += 4
				case 1:
					//fmt.Println(cmp.Base, cmp.I, i)
					nums[nums[i+3]] = nums[i+1] * nums[i+2]
					i += 4
				case 2:
					nums[nums[i+3]] = nums[i+1] * nums[cmp.Base+nums[i+2]]
					i += 4
				}
			case 2:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[cmp.Base+nums[i+1]] * nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[cmp.Base+nums[i+1]] * nums[i+2]
					i += 4
				case 2:
					nums[nums[i+3]] = nums[cmp.Base+nums[i+1]] * nums[cmp.Base+nums[i+2]]
					i += 4
				}
			}
		case 3:
			switch opCode[len(opCode)-3] {
			case 0:
				nums[nums[i+1]] = signal
			case 1:
				nums[i+1] = signal
			case 2:
				nums[cmp.Base+nums[i+1]] = signal
			}
			i += 2
		case 4:
			switch opCode[len(opCode)-3] {
			case 0:
				output = append(output, nums[nums[i+1]])
			case 1:
				output = append(output, nums[i+1])
			case 2:
				output = append(output, nums[cmp.Base+nums[i+1]])
			}
			fmt.Println(output)
			i += 2
		case 5:
			switch opCode[len(opCode)-3] {
			case 0:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[nums[i+1]] != 0 {
						i = nums[nums[i+2]]
					} else {
						i += 3
					}
				case 1:
					if nums[nums[i+1]] != 0 {
						i = nums[i+2]
					} else {
						i += 3
					}
				case 2:
					if nums[nums[i+1]] != 0 {
						i = nums[cmp.Base+nums[i+2]]
					} else {
						i += 3
					}
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[i+1] != 0 {
						i = nums[nums[i+2]]
					} else {
						i += 3
					}
				case 1:
					if nums[i+1] != 0 {
						i = nums[i+2]
					} else {
						i += 3
					}
				case 2:
					if nums[i+1] != 0 {
						i = nums[cmp.Base+nums[i+2]]
					} else {
						i += 3
					}
				}
			case 2:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[cmp.Base+nums[i+1]] != 0 {
						i = nums[nums[i+2]]
					} else {
						i += 3
					}
				case 1:
					if nums[cmp.Base+nums[i+1]] != 0 {
						i = nums[i+2]
					} else {
						i += 3
					}
				case 2:
					if nums[cmp.Base+nums[i+1]] != 0 {
						i = nums[cmp.Base+nums[i+2]]
					} else {
						i += 3
					}
				}
			}
		case 6:
			switch opCode[len(opCode)-3] {
			case 0:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[nums[i+1]] == 0 {
						i = nums[nums[i+2]]
					} else {
						i += 3
					}
				case 1:
					if nums[nums[i+1]] == 0 {
						i = nums[i+2]
					} else {
						i += 3
					}
				case 2:
					if nums[nums[i+1]] == 0 {
						i = nums[cmp.Base+nums[i+2]]
					} else {
						i += 3
					}
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[i+1] == 0 {
						i = nums[nums[i+2]]
					} else {
						i += 3
					}
				case 1:
					if nums[i+1] == 0 {
						i = nums[i+2]
					} else {
						i += 3
					}
				case 2:
					if nums[i+1] == 0 {
						i = nums[cmp.Base+nums[i+2]]
					} else {
						i += 3
					}
				}
			case 2:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[cmp.Base+nums[i+1]] == 0 {
						i = nums[nums[i+2]]
					} else {
						i += 3
					}
				case 1:
					if nums[cmp.Base+nums[i+1]] == 0 {
						i = nums[i+2]
					} else {
						i += 3
					}
				case 2:
					if nums[cmp.Base+nums[i+1]] == 0 {
						i = nums[cmp.Base+nums[i+2]]
					} else {
						i += 3
					}
				}
			}
		case 7:
			switch opCode[len(opCode)-3] {
			case 0:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[nums[i+1]] < nums[nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 1:
					if nums[nums[i+1]] < nums[i+2] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 2:
					if nums[nums[i+1]] < nums[cmp.Base+nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[i+1] < nums[nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 1:
					if nums[i+1] < nums[i+2] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 2:
					if nums[i+1] < nums[cmp.Base+nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				}
			case 2:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[cmp.Base+nums[i+1]] < nums[nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 1:
					if nums[cmp.Base+nums[i+1]] < nums[i+2] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 2:
					if nums[cmp.Base+nums[i+1]] < nums[cmp.Base+nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				}
			}
		case 8:
			switch opCode[len(opCode)-3] {
			case 0:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[nums[i+1]] == nums[nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 1:
					if nums[nums[i+1]] == nums[i+2] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 2:
					if nums[nums[i+1]] == nums[cmp.Base+nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[i+1] == nums[nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 1:
					if nums[i+1] == nums[i+2] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 2:
					if nums[i+1] == nums[cmp.Base+nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				}
			case 2:
				switch opCode[len(opCode)-4] {
				case 0:
					if nums[cmp.Base+nums[i+1]] == nums[nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 1:
					if nums[cmp.Base+nums[i+1]] == nums[i+2] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				case 2:
					if nums[cmp.Base+nums[i+1]] == nums[cmp.Base+nums[i+2]] {
						nums[nums[i+3]] = 1
						i += 4
					} else {
						nums[nums[i+3]] = 0
						i += 4
					}
				}
			}
		case 9:
			cmp.Base += nums[i+1]
			i += 2
		}
	}
	cmp.I = i
	//fmt.Println(amp)
	return output
}

func intToSlice(n int) []int {
	var intSlice []int
	k := 1
	for n != 0 {
		k *= 10
		d := n % 10
		intSlice = append(intSlice, d)
		n /= 10
	}
	reverseInts(intSlice)
	return intSlice
}

func reverseInts(nums []int) {
	for j, k := 0, len(nums)-1; j < k; {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}

func print(nums map[int]int) {
	for i := 0; i < len(nums); i++ {
		fmt.Printf(" %d-%d ", i, nums[i])
	}
	fmt.Println()
}
