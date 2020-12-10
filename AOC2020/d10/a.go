package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func max(ints ...int) int {
	mx := 0
	for _, v := range ints {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func timeit() func() {
	st := time.Now()
	return func() {
		en := time.Now()
		elapsed := en.Sub(st).Seconds()
		fmt.Println("time taken(sec)", elapsed)
	}
}

func ways(arr []int) int {
	// var wg sync.WaitGroup
	stop := timeit()
	defer stop()

	yes := make(map[int]bool)
	for _, v := range arr { yes[v] = true }
	var lim int = max(arr...) + 1
	dp := make([]int, lim)
	dp[0] = 1
	for i := 1; i < lim; i++ {
		var wg sync.WaitGroup
		for j := 1; j < 4; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				if yes[i-j] { dp[i] += dp[i-j] }
			}(j)
			wg.Wait()
		}
	}
	return dp[lim-1]
}

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic("error reading file")
	}
	str := strings.Split(string(bytes), "\n")
	arr := make([]int, len(str), len(str)+2)
	for i, v := range str {
		arr[i], err = strconv.Atoi(v)
		if err != nil {
			panic("error parsing")
		}
	}
	arr = append(arr, 0)
	arr = append(arr, max(arr...)+3)
	sort.Sort(SortBy(arr))
	d1, d3 := 0, 0
	for i, v := range arr {
		if i > 0 {
			if v-arr[i-1] == 1 {
				d1++
			} else if v-arr[i-1] == 3 {
				d3++
			}
		}
	}
	//first part
	fmt.Println(d1, d3, d1*d3)
	//second part
	fmt.Println(ways(arr))

}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

/*
ok now we have understood Wgs
now to look at mutexes
and then channels

*/
