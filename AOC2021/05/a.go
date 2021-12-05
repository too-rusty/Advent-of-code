package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Line struct {
	from, to Point
}

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := make([]Line, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		z := strings.Split(line, " ")
		p1, p2 := strings.Split(z[0], ","), strings.Split(z[2], ",")
		// fmt.Println(z1, z2)
		A, _ := strconv.Atoi(p1[0])
		B, _ := strconv.Atoi(p1[1])
		C, _ := strconv.Atoi(p2[0])
		D, _ := strconv.Atoi(p2[1])
		lines = append(lines, Line{
			Point{A, B},
			Point{C, D},
		})
	}
	overLaps := make(map[Point]int)

	for _, line := range lines {
		// if !line.isHorizontal() && !line.isVertical() {
		// 	continue
		// } // partB
		for _, v := range line.getAllPoints() {
			overLaps[v]++
		}
	}
	ans := 0
	for k := range overLaps {
		// fmt.Println(k, overLaps[k])
		if overLaps[k] >= 2 {
			ans++
		}
	}
	fmt.Println(ans)

}

func (l Line) isHorizontal() bool { return l.from.Y == l.to.Y }

func (l Line) isVertical() bool { return l.from.X == l.to.X }

func (l Line) getAllPoints() (points []Point) {

	if l.from == l.to {
		points = []Point{l.from}
		return
	}

	if l.from.X == l.to.X {
		// vertical line
		if l.from.Y > l.to.Y {
			l.from, l.to = l.to, l.from
		}
		for y := l.from.Y; y <= l.to.Y; y++ {
			points = append(points, Point{l.from.X, y})
		}
		return
	}

	if l.from.X > l.to.X {
		l.from, l.to = l.to, l.from
	}

	m := l.slope()
	C := l.from.Y - m*l.from.X
	// y = mx + c
	// c = y - mx => y1 - slope*x1
	for x := l.from.X; x <= l.to.X; x++ {
		y := m*x + C
		points = append(points, Point{x, y})
	}

	return
}

func (l Line) slope() int {
	// assuming the line is not a point
	return (l.to.Y - l.from.Y) / (l.to.X - l.from.X)
}
