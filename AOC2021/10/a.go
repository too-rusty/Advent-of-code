package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {

	bytes, _ := ioutil.ReadFile("input")
	arr := strings.Split(string(bytes), "\n")
	ans := uint(0)
	pp := make([]int, 0)
	for _, str := range arr {
		var ok bool
		var ss []byte
		x := getPoints([]byte(str))
		ans += x
		if x > 0 {
			continue
		} // remove corrupted first then answer is from the rest of the incomplete
		if ok, ss = isIncomplete([]byte(str)); ok {
			// fmt.Println("sss", ss)
			ss = []byte(Reverse(string(ss)))
			for i, v := range ss {
				ss[i] = opposite(v)
			}
			POINT := uint(0)
			for _, v := range ss {
				POINT *= 5
				POINT += Byte(v).score()
			}
			pp = append(pp, int(POINT))
			continue
		}

	}
	// fmt.Println("ans", ans)
	sort.Ints(pp)
	fmt.Println(pp[len(pp)/2])
}

var id int

func isIncomplete(b []byte) (bool, []byte) {
	// for
	// fmt.Println("check for ", string(b))
	defer func() { id++ }()
	stk := Stack(make([]byte, 0))
	stk_ptr := &stk
	for _, c := range b {
		if isOpen(c) {
			stk_ptr.Push(c)
		} else {
			stk_ptr.Pop()
		}
		fmt.Println(id, string(stk))
	}
	if len(stk) > 0 {
		fmt.Println("stk", string(stk))
		return true, []byte(stk)
	}
	return false, []byte{}
}

func getPoints(b []byte) uint {
	// this is not incomplete
	stk := Stack(make([]byte, 0))
	stk_ptr := &stk
	ret := uint(0)
	for _, c := range b {
		if isOpen(c) {
			stk_ptr.Push(c)
		} else {
			cc := stk_ptr.Pop()
			if opposite(cc) != c {
				ret += Byte(c).points()
				break // since we need first invalid char only
			}
		}
	}
	return ret
}

func isClose(b byte) bool {
	switch b {
	case ')' | '}' | ']' | '>':
		return true
	default:
		return false
	}
}

func isOpen(b byte) bool {
	switch b {
	case '(', '{', '[', '<':
		return true
	default:
	}
	return false
}

func opposite(b byte) byte {
	switch b {
	case '{':
		return '}'
	case '[':
		return ']'
	case '(':
		return ')'
	case '<':
		return '>'
	default:
		return '-'
	}
}

func (b Byte) points() uint {
	switch b {
	case '}':
		return 1197
	case ']':
		return 57
	case ')':
		return 3
	case '>':
		return 25137
	default:
		return 0
	}
}

func (b Byte) score() uint {
	switch b {
	case '}':
		return 3
	case ']':
		return 2
	case ')':
		return 1
	case '>':
		return 4
	default:
		return 0
	}
}

type Byte byte
type Stack []byte

func (s *Stack) Push(b byte) { *s = append(*s, b) }

func (s *Stack) Pop() byte {
	var pop = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}
func (s *Stack) Top() byte  { return (*s)[len(*s)-1] }
func (s *Stack) Size() uint { return uint(len(*s)) }

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// [({([[{[]{

// 	{([({{[(
