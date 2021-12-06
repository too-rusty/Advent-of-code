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
		fmt.Println(err)
		return
	}
	// if at some day we have n Xs (X != 0) then next day we have n (X-1)s
	// if at some day we have n 0s then next day we have n 8s and n 6s
	nums := strings.Split(string(bytes), ",")
	cnt := make(map[int]int)

	for _, v := range nums {
		x, _ := strconv.Atoi(v)
		cnt[x]++
	}

	const FINAL_DAY = 256
	ates := cnt[0]
	for day := 1; day <= FINAL_DAY; day++ {
		tmp := cnt[0]
		for i := 0; i < 8; i++ {
			cnt[i] = cnt[i+1]
		}
		cnt[8] = ates
		cnt[6] += tmp
		ates = cnt[0]
	}
	var ans int
	for i := 0; i < 9; i++ {
		ans += cnt[i]
	}
	fmt.Println(ans)

}
