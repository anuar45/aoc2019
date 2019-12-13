package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	X   int
	Y   int
	Val int
	See []Point
}

type Points struct {
	SizeX  int
	SizeY  int
	Points []Point
}

func main() {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	var points []Point
	var x, y int
	for scanner.Scan() {
		x = 0
		for _, v := range scanner.Text() {
			point := Point{X: x, Y: y}
			if v == '#' {
				point.Val = 1
			}
			points = append(points, point)
			x++
		}
		y++
	}

	pMap := Points{SizeX: x, SizeY: y, Points: points}

	//fmt.Println(pMap)
	//print(pMap)

	for i := 0; i < len(pMap.Points); i++ {
		if pMap.Points[i].Val == 1 {
		loop:
			for j := i + 1; j < len(pMap.Points); j++ {
				//fmt.Println(p1)
				if pMap.Points[j].Val == 1 && !(pMap.Points[i].X == pMap.Points[j].X && pMap.Points[i].Y == pMap.Points[j].Y) {
					switch {
					case pMap.Points[i].X == pMap.Points[j].X:
						for k := i + 1; k < j; k++ {
							if pMap.Points[k].Val == 1 && pMap.Points[i].X == pMap.Points[k].X {
								continue loop
							}
						}
						pMap.Points[i].See = append(pMap.Points[i].See, pMap.Points[j])
						pMap.Points[j].See = append(pMap.Points[j].See, pMap.Points[i])
					case pMap.Points[i].Y == pMap.Points[j].Y:
						for k := i + 1; k < j; k++ {
							if pMap.Points[k].Val == 1 && pMap.Points[i].Y == pMap.Points[k].Y {
								continue loop
							}
						}
						pMap.Points[i].See = append(pMap.Points[i].See, pMap.Points[j])
						pMap.Points[j].See = append(pMap.Points[j].See, pMap.Points[i])
					default:
						m := findSlope(pMap.Points[i].X, pMap.Points[i].Y, pMap.Points[j].X, pMap.Points[j].Y)
						b1 := float64(pMap.Points[i].Y) - m*float64(pMap.Points[i].X)
						//fmt.Printf("%d %d  %d %d  %F\n", pMap.Points[i].X, pMap.Points[i].Y, pMap.Points[j].X, pMap.Points[j].Y, m)
						for k := i + 1; k < j; k++ {
							if pMap.Points[k].Val == 1 {
								//fmt.Println(m1 * pMap.Points[k].X % m2)
								b2 := float64(pMap.Points[k].Y) - m*float64(pMap.Points[k].X)
								//fmt.Println(pMap.Points[k].X, pMap.Points[k].Y, b1, b2)
								if compareFloats(b1, b2) {
									continue loop
								}
							}
						}
						pMap.Points[i].See = append(pMap.Points[i].See, pMap.Points[j])
						pMap.Points[j].See = append(pMap.Points[j].See, pMap.Points[i])
					}
				}
			}
		}
	}
	// 22 19
	fmt.Println(pMap.Points[554].X, pMap.Points[554].Y)
	fmt.Println(pMap.SizeX, pMap.SizeY)

	var clockPoints = pMap.Points[554].See
	for i := 0; i < len(clockPoints); i++ {

	}

	for i, v := range clockPoints {
		fmt.Println(i, v.X, v.Y)
	}

	// sort.Slice(pMap.Points[554].See, func(i, j int) bool {
	// 	return (pMap.Points[554].See[i].X + pMap.Points[554].See[i].Y) < (pMap.Points[554].See[j].X + pMap.Points[554].See[j].Y)
	// },
	// )

	// for i, v := range pMap.Points[554].See {
	// 	fmt.Println(i, v.X, v.Y)
	// }

	// for i, v := range pMap.Points {
	// 	if v.Val == 1 {
	// 		fmt.Println(i, v.X, v.Y, len(v.See))
	// 	}
	// }

	// sort.Slice(pMap.Points, func(i, j int) bool {
	// 	return len(pMap.Points[i].See) < len(pMap.Points[j].See)
	// },
	// )

	// //22 19 282
	// var maxPoint Point
	// for _, v := range pMap.Points {
	// 	if len(maxPoint.See) < len(v.See) {
	// 		maxPoint = v
	// 	}
	// }

	// fmt.Println(maxPoint.X, maxPoint.Y, len(maxPoint.See))

}

// TODO: Needs rework
func print(points Points) {
	for i, point := range points.Points {
		fmt.Print(point.Val)
		if i == points.SizeX {
			fmt.Println()
		}
	}
}

func findSlope(x1, y1, x2, y2 int) float64 {
	var result float64
	m1 := y2 - y1
	m2 := x2 - x1
	result = float64(m1) / float64(m2)
	return result
}

func compareFloats(f1, f2 float64) bool {
	const tolerance = .00001
	var result bool
	diff := math.Abs(f2 - f1)
	mean := math.Abs(f2+f1) / 2.0
	if (diff / mean) < tolerance {
		result = true
	}
	return result
}

func less(center, a, b Point) bool {
	if a.X-center.X >= 0 && b.X-center.X < 0 {
		return true
	}
	if a.X-center.X < 0 && b.X-center.X >= 0 {
		return false
	}
	if a.X-center.X == 0 && b.X-center.X == 0 {
		if a.Y-center.Y >= 0 || b.Y-center.Y >= 0 {
			return a.Y > b.Y
		}
		return b.Y > a.Y
	}

	// compute the cross product of vectors (center -> a) x (center -> b)
	det := (a.X-center.X)*(b.Y-center.Y) - (b.X-center.X)*(a.Y-center.Y)
	if det < 0 {
		return true
	}
	if det > 0 {
		return false
	}

	// points a and b are on the same line from the center
	// check which point is closer to the center
	d1 := (a.X-center.X)*(a.X-center.X) + (a.Y-center.Y)*(a.Y-center.Y)
	d2 := (b.X-center.X)*(b.X-center.X) + (b.Y-center.Y)*(b.Y-center.Y)
	return d1 > d2
}
