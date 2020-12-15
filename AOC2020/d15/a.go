package main

import (
	"fmt"
	"time"
)

type Pair struct {
	fi, se int
}

func main() {
	stop1 := timeit("first sol")
	m := make(map[int]Pair)
	arr := []int{2, 0, 6, 12, 1, 3}
	const lim int = 30000000
	arr2 := make([]int, lim)
	copy(arr2, arr)
	for i := 0; i < lim; i++ {
		if i < len(arr) {
			m[arr2[i]] = Pair{0, i + 1}
			continue
		}
		if m[arr2[i-1]].fi == 0 {
			//last one was spoken first time
			arr2[i] = 0
		} else {
			//last one was spoken second time
			arr2[i] = m[arr2[i-1]].se - m[arr2[i-1]].fi
		}
		m[arr2[i]] = Pair{m[arr2[i]].se, i + 1}
	}
	fmt.Println(arr2[lim-1])
	stop1()

	stop2 := timeit("second sol")
	OnSolution()
	stop2()

}

func OnSolution() {
	//O(n) solution
	const lim int = 30000000
	arr := []int{2, 0, 6, 12, 1, 3}
	arr2 := make([]int, lim)
	copy(arr2, arr)

	fi, se := make([]int, lim), make([]int, lim)

	for i := 0; i < lim; i++ {
		if i < len(arr) {
			se[arr2[i]] = i + 1
			continue
		}
		if fi[arr2[i-1]] == 0 {
			arr2[i] = 0
		} else {
			arr2[i] = se[arr2[i-1]] - fi[arr2[i-1]]
		}
		fi[arr2[i]], se[arr2[i]] = se[arr2[i]], i+1
	}

	fmt.Println(arr2[lim-1])
}

func timeit(msg string) func() {
	st := time.Now()
	return func() {
		en := time.Now()
		elapsed := en.Sub(st).Seconds()
		fmt.Println("time taken by : ", msg, " : ", elapsed)
	}
}
