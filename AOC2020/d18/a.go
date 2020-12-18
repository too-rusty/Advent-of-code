package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var dbg = flag.Bool("dbg", false, "print dbg statements")

func getOpeningAndClosing(s string) ([]int, []int) {
	closing, opening := make([]int, len(s)), make([]int, len(s))
	stk := make([]int, 0)
	for i := 0; i < len(s); i++ {
		closing[i], opening[i] = -1, -1
	}
	for i, v := range s {
		switch v {
		case '(':
			stk = append(stk, i)
		case ')':
			closing[stk[len(stk)-1]] = i
			opening[i] = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
	}
	if len(stk) != 0 { panic("should be empty") }
	return opening, closing
}

func isBalanced(s string) bool {
	op := 0
	for _, v := range s {
		if v == '(' {
			op++
		} else if v == ')' {
			if op == 0 {
				return false
			}
			op--
		}
	}
	return op == 0
}

func solve(s string) int {
	// expr1 + expr2 -> evaluate expr1 and expr2 and add em
	// also evaluation needs to be done from right to left
	opening, _ := getOpeningAndClosing(s)
	var evaluate func(l, r int) int
	evaluate = func(l, r int) int {
		//inclusive
		if *dbg {
			fmt.Println("eval", s[l:r+1])
		}
		if l == r {
			x, err := strconv.Atoi(string(s[l]))
			if err != nil {
				panic("unidentified token")
			}
			return x
		}
		var ret int
		if s[r] == ')' {
			ret = evaluate(opening[r]+1, r-1)
			if opening[r]-2 >= l {
				if s[opening[r]-1] == '+' {
					ret += evaluate(l, opening[r]-2)
				} else {
					ret *= evaluate(l, opening[r]-2)
				}
			}
		} else {
			ret = evaluate(r, r)
			if r-1 > l {
				if s[r-1] == '+' {
					ret += evaluate(l, r-2)
				} else {
					ret *= evaluate(l, r-2)
				}
			}
		}
		return ret
	}

	ans := evaluate(0, len(s)-1)
	if *dbg { fmt.Println("parta ans", ans) }
	return ans
}

func applyBrackets(s string) string {
	opening, closing := getOpeningAndClosing(s)
	cops, cclo := make([]int, len(s)), make([]int, len(s))
	for i, v := range s {
		if v == '+' {
			if opening[i-1] == -1 {
				cops[i-1]++
			} else {
				cops[opening[i-1]]++
			}
			if closing[i+1] == -1 {
				cclo[i+1]++
			} else {
				cclo[closing[i+1]]++
			}
		}
	}
	str := make([]rune, 0)
	for i, v := range s {
		for cops[i] > 0 {
			str = append(str, '(')
			cops[i]--
		}
		str = append(str, v)
		for cclo[i] > 0 {
			str = append(str, ')')
			cclo[i]--
		}
	}
	if *dbg {
		fmt.Println("str after brackets", string(str))
	}
	if !isBalanced(string(str)) { panic("not balanced") }
	return string(str)
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	arr := strings.Split(string(bytes), "\n")
	ans1, ans2 := 0, 0
	for _, s := range arr {
		s = strings.ReplaceAll(s, " ", "")
		ans1 += solve(s)
		ans2 += solve(applyBrackets(s))
	}
	fmt.Println("parta  partb", ans1, ans2)

}

/*
(8+(((9+7)*4*9))+4)*8*((6*6*9*8*3*7)+2)
(9+7)*4*9+4*8*(6*6*9*8*3*7)
*/
