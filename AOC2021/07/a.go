package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		return
	}
	arr := make([]int, 0)
	for _, v := range strings.Split(string(bytes), ",") {
		x, _ := strconv.Atoi(v)
		arr = append(arr, x)
	}

	X := max(arr)
	ans := -1
	for i := 0; i < X+1; i++ {
		tmp := 0
		for _, v := range arr {
			n := abs(v - i)
			tmp += (n * (n + 1)) / 2
		}
		if ans == -1 || tmp < ans {
			ans = tmp
		}
	}
	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(arr []int) int {
	mx := 0
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	for _, v := range arr {
		mx = max(mx, v)
	}
	return mx
}
