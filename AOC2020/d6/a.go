package main

import "fmt"
import "strings"
import "io/ioutil"

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	grps := strings.Split(string(bytes), "\n\n")
	final_cnt := 0
	for _, grp := range grps {
		there := make([]int, 26)
		split := strings.Split(grp, "\n")
		for _, line := range split {
			for _, b := range []byte(line) {
				there[b-byte('a')]++
				if there[b-byte('a')] == len(split) {
					final_cnt++
				}
			}
		}
	}
	fmt.Println(final_cnt)

}
