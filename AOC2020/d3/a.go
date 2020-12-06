package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var partb = flag.Bool("partb", false, "do part b")
var inputFile = flag.String("input", "input", "input path")

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := strings.Split(string(bytes), "\n")
	m := len(grid[0])
	cnt := 0
	for i, row := range grid {
		if i == 0 {
			continue
		}
		if row[(3*i)%m] == '#' {
			cnt++
		}
	}
	fmt.Println(cnt)

	if *partb {
		shifts := []int{1, 5, 7}
		cnts, cnt2 := make([]int, 3), 0
		for i, row := range grid {
			if i == 0 {
				continue
			}
			for idx, shift := range shifts {
				if row[(shift*i)%m] == '#' {
					cnts[idx]++
				}
			}
			if i%2 == 0 && row[(i/2)%m] == '#' {
				cnt2++
			}
		}
		res := cnt * cnt2
		for _, v := range cnts {
			res *= v
		}
		fmt.Println(res)
	}
}
