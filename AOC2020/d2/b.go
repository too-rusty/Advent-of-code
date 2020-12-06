package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("input", "input", "input relative to me")
var partB = flag.Bool("partb", false, "solve b")

type Data struct {
	lo, hi int
	c      byte
	str    []byte
}

func parseData(s string) Data {
	str := strings.Split(s, " ")
	val := strings.Split(str[0], "-")
	a, _ := strconv.Atoi(val[0])
	b, _ := strconv.Atoi(val[1])

	data := Data{
		lo: a, hi: b,
		c:   byte(str[1][0]),
		str: []byte(str[2]),
	}
	return data
}

func (d Data) Valid() bool {
	cnt := 0
	for _, v := range d.str {
		if v == d.c {
			cnt++
		}
	}
	return cnt >= d.lo && cnt <= d.hi
}

func (d Data) Valid2() bool {
	cond1 := d.lo <= len(d.str) && d.str[d.lo-1] == d.c
	cond2 := d.hi <= len(d.str) && d.str[d.hi-1] == d.c
	var c1, c2 int8
	if cond1 {
		c1 = 1
	}
	if cond2 {
		c2 = 1
	}
	if z := c1 ^ c2; z == 1 {
		return true
	}
	return false
}

func byte2Char(b byte) string {
	return fmt.Sprintf("%c", b)
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	strarr := strings.Split(string(bytes), "\n")
	arr := make([]Data, len(strarr))
	cnt := 0
	for i, str := range strarr {
		arr[i] = parseData(str)
		if arr[i].Valid() {
			cnt++
		}
	}
	fmt.Println(cnt)

	if *partB {
		cnt = 0
		for _, val := range arr {
			if val.Valid2() {
				cnt++
			}
		}
		fmt.Println(cnt)
	}

}
