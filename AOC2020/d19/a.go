package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var rules [][][]int
var target string
var sp1, sp2 int
var sc1, sc2 string

func canGenerate(s string, ruleList []int) bool {
	// fmt.Println("s rlis ", s, ruleList)
	// if len(s) == 1 && len(ruleList) == 1 {
	//  fmt.Println("wokay", s, ruleList)
	//  return (ruleList[0] == sp1 && sc1 == s) || (ruleList[0] == sp2 && sc2 == s)
	// }
	if len(s) == 0 && len(ruleList) == 0 {
		return true
	}
	if len(ruleList) == 0 {
		return false
	}

	for i := 1; i <= len(s); i++ {
		//non empty
		// fmt.Println("chkR on ", s[:i], ruleList[0])
		// fmt.Println("canG on ", s[i:], ruleList[1:], "\n---")
		if checkRule(s[:i], ruleList[0]) && canGenerate(s[i:], ruleList[1:]) {
			return true
		}
	}

	return false
}

type Pair struct {
	s string
	n int
}

var dp = make(map[Pair]bool)

func checkRule(s string, n int) bool {
	if v, ok := dp[Pair{s, n}]; ok {
		return v
	}
	pair := Pair{s, n}
	var res bool
	if n == sp1 && s == sc1 {
		dp[pair] = true
		return true
	}
	if n == sp2 && s == sc2 {
		dp[pair] = true
		return true
	}
	for _, rule := range rules[n] {
		x := canGenerate(s, rule)
		res = res || x
	}
	dp[pair] = res
	return res
}

func parseList(s, sep string) []int {
	str := strings.Split(s, sep)
	arr := make([]int, len(str))
	for i, v := range str {
		x, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("v", v)
			panic("invalid")
		}
		arr[i] = x
	}
	return arr
}

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Print(err)
		return
	}
	str := strings.Split(string(bytes), "\n\n")
	ruless := str[0]
	notRuless := strings.Split(str[1], "\n")

	rules = make([][][]int, len(strings.Split(ruless, "\n"))+100)
	for _, v := range strings.Split(ruless, "\n") {
		x := strings.Split(v, " ")

		ruleNo, err := strconv.Atoi(x[0][:len(x[0])-1])
		if err != nil {
			panic("error in rule no parsing")
		}
		strr := strings.Split(strings.TrimRight(strings.Join(x[1:], " "), "\n"), " | ")
		// fmt.Println(strr)
		if strings.Compare(strr[0], "\"a\"") == 0 {
			sp1 = ruleNo
			sc1 = "a"
			// fmt.Println("sp1 set sc1", sp1, sc1)
			continue
		}
		if strings.Compare(strr[0], "\"b\"") == 0 {
			sp2 = ruleNo
			sc2 = "b"
			continue
		}
		rules[ruleNo] = make([][]int, 0)
		for _, vv := range strr {
			// fmt.Println(vv)
			rules[ruleNo] = append(rules[ruleNo], parseList(vv, " "))
		}
	}
	fmt.Println(rules)
	_ = notRuless

	ans := 0
	for _, s := range notRuless {
		ress := checkRule(s, 0)
		// fmt.Println(ress)
		if ress {
			ans++
		}

	}
	fmt.Println(ans)
	// fmt.Println(checkRule(notRuless[4], 0))

	// for k := range dp {
	//  fmt.Print(k)
	// }

}
