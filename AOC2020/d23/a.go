package main

import (
	"flag"
	"fmt"
	"strconv"
)

var partB = flag.Bool("partb", false, "partb solve")

func main() {
	flag.Parse()
	if !*partB {
		str := "538914762"
		n := len(str)
		arr := make([]int, n)
		arr2 := make([]int, 2*n)
		for i, _ := range str {
			v, err := strconv.Atoi(str[i : i+1])
			if err != nil {
				panic("err1")
			}
			arr[i] = v
			arr2[i], arr2[i+n] = v, v

		}
		fmt.Println(arr)
		pick := arr[0]
		for times := 0; times < 100; times++ {
			var tmp = make([]int, 3)
			var newPick, placeAfter int
			for i := 0; i < 2*n; i++ {
				if arr2[i] == pick && arr2[i+4] < 2*n {
					copy(tmp, arr2[i+1:i+4])
					newPick = arr2[i+4]
					break
				}
			}
			placeAfter = pick - 1
			for i := 0; i < 20; i++ {
				if placeAfter == 0 {
					placeAfter = 9
				}
				if placeAfter == tmp[0] || placeAfter == tmp[1] || placeAfter == tmp[2] {
					placeAfter--
				} else {
					break
				}
			}
			// ---
			arrtmp := make([]int, n)
			vis := make(map[int]bool)
			for i := 1; i < 10; i++ {
				vis[i] = false
			}
			vis[placeAfter] = true
			arrtmp[0] = placeAfter
			for i := 1; i < 4; i++ {
				arrtmp[i] = tmp[i-1]
				vis[arrtmp[i]] = true
			}
			// fmt.Println("next aft", newPick, placeAfter)
			// at := 4
			for i := 0; i < 2*n; i++ {
				if arr2[i] == placeAfter && arr2[i+10] < 2*n {
					at := i + 1
					for j := 4; j < n; j++ {
						for vis[arr2[at]] {
							at++
						}
						arrtmp[j] = arr2[at]
						vis[arrtmp[j]] = true
					}
					break
				}
			}
			arr = arrtmp
			for i := 0; i < n; i++ {
				arr2[i], arr2[i+n] = arr[i], arr[i]
			}
			pick = newPick
			fmt.Println("arrtmp next", arr, pick)
		}
		for i := 0; i < 2*n; i++ {
			if arr2[i] == 1 && i+9 < 2*n {
				fmt.Println(arr2[i+1 : i+9])
			}
		}
	} else {

		// partb
		//TODO
		fmt.Println("TODO")
	}

}

type Node struct {
	next *Node
	val  int
}
