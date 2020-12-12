package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
func turnLeft(dx, dy, degree int) (int, int) {
	times := degree / 90
	f := func(dx, dy int) (int, int) {
		if dx == -1 && dy == 0 {
			dx, dy = 0, -1
		} else if dx == 0 && dy == -1 {
			dx, dy = 1, 0
		} else if dx == 1 && dy == 0 {
			dx, dy = 0, 1
		} else {
			dx, dy = -1, 0
		}
		return dx, dy
	}
	for times > 0 {
		dx, dy = f(dx, dy)
		times--
	}
	return dx, dy
}

func turnWp(wpx, wpy, degree int) (int, int) {
	times := degree / 90
	f := func(wpx, wpy int) (int, int) {
		return wpy, -wpx
	}
	for times > 0 {
		wpx, wpy = f(wpx, wpy)
		times--
	}
	return wpx, wpy
}

func main() {
	partA()
	partB()
}

func partA() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic("err in reading file")
	}
	str := strings.Split(string(bytes), "\n")
	px, py, dx, dy := 0, 0, 1, 0
	for _, v := range str {
		num, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("couldn't parse")
		}
		switch v[0] {
		case byte('F'):
			px += dx * num
			py += dy * num
		case byte('L'):
			dx, dy = turnLeft(dx, dy, num)
		case byte('R'):
			dx, dy = turnLeft(dx, dy, 360-num)
		case byte('N'):
			py += num
		case byte('S'):
			py -= num
		case byte('E'):
			px += num
		case byte('W'):
			px -= num
		}
	}
	fmt.Println(abs(px) + abs(py))
}
func partB() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic("err in reading file")
	}
	str := strings.Split(string(bytes), "\n")
	px, py := 0, 0
	wpx, wpy := 10, 1
	for _, v := range str {
		num, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("couldn't parse")
		}
		switch v[0] {
		case byte('F'):
			px += wpx * num
			py += wpy * num
		case byte('L'):
			wpx, wpy = turnWp(wpx, wpy, 360-num)
		case byte('R'):
			wpx, wpy = turnWp(wpx, wpy, num)
		case byte('N'):
			wpy += num
		case byte('S'):
			wpy -= num
		case byte('E'):
			wpx += num
		case byte('W'):
			wpx -= num
		}
	}
	fmt.Println(abs(px) + abs(py))

}

/*
 N
W  E
 S


*/
