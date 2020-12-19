package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	rules                [][][]int
	target, term1, term2 string // terminal chars
	tid1, tid2           int    // rule containing the terminal chars
  dp = make(map[Pair]bool)
)


func parseList(s, sep string) []int {
  str := strings.Split(s, sep)
  arr := make([]int, len(str))
  for i, v := range str {
    x, err := strconv.Atoi(v)
    if err != nil { panic("invalid") }
    arr[i] = x
  }
  return arr
}

type Pair struct {
  s string
  n int
}


// -------------------------------
func evalSeq(s string, r []int) bool {
	// a sequence of rules
	if len(s) == 0 && len(r) == 0 { return true }
	if len(r) == 0 { return false }
	for i := 1; i <= len(s); i++ {
		if evalOne(s[:i], r[0]) && evalSeq(s[i:], r[1:]) {
			return true
		}
	}
	return false
}

func evalOne(s string, n int) bool {
	if v, ok := dp[Pair{s, n}]; ok { return v }
	ret := Pair{s, n}
	var res bool
	if (n == tid1 && s == term1) || (n == tid2 && s == term2) {
		dp[ret] = true
	} else {
		for _, seq := range rules[n] { res = res || evalSeq(s, seq) }
		dp[ret] = res
	}
	return res
}
// --------------------------------


func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil { panic("err in reading file") }
	str := strings.Split(string(bytes), "\n\n")
	ruless := str[0]
	notRuless := strings.Split(str[1], "\n")

	rules = make([][][]int, len(strings.Split(ruless, "\n"))+100)
	// some offset just to be on the safe side and for p2 example

	for _, v := range strings.Split(ruless, "\n") {
		x := strings.Split(v, " ")
		ruleNo, err := strconv.Atoi(x[0][:len(x[0])-1])
		if err != nil { panic("error in rule no parsing") }
		strr := strings.Split(strings.TrimRight(strings.Join(x[1:], " "), "\n"), " | ")
		if strings.Compare(strr[0], "\"a\"") == 0 {
			tid1, term1 = ruleNo, "a"
			continue
		}
		if strings.Compare(strr[0], "\"b\"") == 0 {
			tid2, term2 = ruleNo, "b"
			continue
		}
		rules[ruleNo] = make([][]int, 0)
		for _, vv := range strr {
			rules[ruleNo] = append(rules[ruleNo], parseList(vv, " "))
		}
	}
	// fmt.Println(rules)

	ans := 0
	for _, s := range notRuless {
		ress := evalOne(s, 0)
		if ress { ans++ }

	}
	fmt.Println(ans)
	// fmt.Println(evalOne(notRuless[4], 0))

}
