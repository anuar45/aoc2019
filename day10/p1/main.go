package main

import (
	"bufio"
	"fmt"
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

	pMap := Points{SizeX: int(x), SizeY: int(y), Points: points}

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
						m1 := pMap.Points[j].Y - pMap.Points[i].Y
						m2 := pMap.Points[j].X - pMap.Points[i].X
						b1 := pMap.Points[i].Y - m1*pMap.Points[i].X/m2
						//fmt.Printf("%d %d  %d %d  %d %d\n", pMap.Points[i].X, pMap.Points[i].Y, pMap.Points[j].X, pMap.Points[j].Y, m1, m2)
						for k := i + 1; k < j; k++ {
							if pMap.Points[k].Val == 1 {
								fmt.Println(m1 * pMap.Points[k].X % m2)
								b2 := pMap.Points[k].Y - m1*pMap.Points[k].X/m2
								//fmt.Println(pMap.Points[k].X, pMap.Points[k].Y, b1, b2)
								if b1 == b2 {
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

	for _, v := range pMap.Points {
		if v.Val == 1 {
			fmt.Println(v.X, v.Y, len(v.See))
		}
	}

	var maxPoint Point
	for _, v := range pMap.Points {
		if len(maxPoint.See) < len(v.See) {
			maxPoint = v
		}
	}

	fmt.Println(maxPoint.X, maxPoint.Y, len(maxPoint.See))

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
