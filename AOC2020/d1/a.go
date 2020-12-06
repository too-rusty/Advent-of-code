package main

import (
	"flag"
	"fmt"
	. "fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input", "inp fpath relative to me")
var partB = flag.Bool("partB", false, "use partb logic")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		Println("error in reading file or doesnt exist")
		return
	}
	content := string(bytes)
	split := strings.Split(content, "\n")
	arr := make([]int, len(split))
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("error in parsing")
			break
		}
		arr[i] = n
	}

	for i, n := range arr {
		for j, m := range arr {
			if i >= j {
				continue
			}
			if n+m == 2020 {
				fmt.Println(n * m)
			}
		}
	}

	if *partB {
		for i, n := range arr {
			for j, m := range arr {
				if i >= j {
					continue
				}
				for k, o := range arr {
					if j >= k {
						continue
					}
					if n+m+o == 2020 {
						fmt.Println(n * m * o)
					}
				}
			}
		}
	}

}
