package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var inputFile = flag.String("input", "input", "file path")

func main() {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	arr := strings.Split(string(bytes), "\n")
	seat := make([]int, len(arr))
	ans := 0
	for i, str := range arr {
		// _=str
		l1, h1, l2, h2 := 0, 127, 0, 7
		for _, char := range str {
			switch char {
			case 'F':
				h1 -= (h1 - l1 + 1) / 2
			case 'B':
				l1 += (h1 - l1 + 1) / 2
			case 'R':
				l2 += (h2 - l2 + 1) / 2
			case 'L':
				h2 -= (h2 - l2 + 1) / 2
			default:
				panic("")
			}
		}
		seat[i] = l1*8 + l2
		ans = max(ans, l1*8+l2)
	}
	fmt.Println(ans)
	sort.Ints(seat)
	for i, v := range seat {
		if i > 0 && v-seat[i-1] == 2 {
			fmt.Println(v - 1)
		}
	}
}
