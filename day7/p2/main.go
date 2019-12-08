package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// type phaseSet struct {
// 	seq    []int
// 	result int
// }

type Amplifier struct {
	I     int
	Phase int
	Code  []int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	inputSlice := strings.Split(string(input), ",")

	var nums []int
	for _, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	//var phaseSets []phaseSet
	var out [][]int
	phaseSeq := []int{5, 6, 7, 8, 9}
	heapPermutation(phaseSeq, len(phaseSeq), &out)

	//out = [][]int{[]int{9, 8, 7, 6, 5}}
	//numsCopy := make([]int, len(nums))
	var result, maxPhase, last int

	for i := 0; i < len(out); i++ {
		result = 0
		var amps []Amplifier
		for j := 0; j < 5; j++ {
			var amp Amplifier
			amp.Phase = out[i][j]
			amp.Code = make([]int, len(nums))
			copy(amp.Code, nums)
			amps = append(amps, amp)
		}
	loop1:
		for {
			for j := 0; j < len(amps); j++ {
				result = amps[j].intCodeComp(result)
				fmt.Println(result, amps[j].I)
			}

			if result == 0 {
				break loop1
			}
			last = result
		}

		if last > maxPhase {
			maxPhase = last
		}
	}

	fmt.Println(maxPhase)

}

func (amp *Amplifier) intCodeComp(signal int) int {
	var output, i int
	var wait bool
	nums := amp.Code
loop:
	for i = amp.I; i < len(nums); {
		opCode := intToSlice(nums[i])
		//fmt.Println(opCode)
		//print(nums)
		if nums[i] == 99 {
			//fmt.Println("terminated")
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
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[i+1] + nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[i+1] + nums[i+2]
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
				}
			case 1:
				switch opCode[len(opCode)-4] {
				case 0:
					nums[nums[i+3]] = nums[i+1] * nums[nums[i+2]]
					i += 4
				case 1:
					nums[nums[i+3]] = nums[i+1] * nums[i+2]
					i += 4
				}
			}
		case 3:
			if i == 0 {
				nums[nums[i+1]] = amp.Phase
			} else if !wait {
				nums[nums[i+1]] = signal
				wait = true
			} else {
				break loop
			}
			i += 2
		case 4:
			switch opCode[len(opCode)-3] {
			case 0:
				output = nums[nums[i+1]]
				i += 2
			case 1:
				output = nums[i+1]
				i += 2
			}
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
				}
			}
		}
	}
	amp.I = i
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

func print(nums []int) {
	for i, v := range nums {
		fmt.Printf(" %d-%d ", i, v)
	}
	fmt.Println()
}

func heapPermutation(a []int, size int, out *[][]int) {
	if size == 1 {
		t := make([]int, len(a))
		copy(t, a)
		*out = append(*out, t)
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, out)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}
